from concurrent import futures
from models.document_info import DocumentInfo
from services.processor import DocumentProcessor
from providers.factory import get_embedding_service, get_chat_service

import grpc

from .document_pb2 import (
    ProcessDocumentResponse,
    EmbedQueryResponse,
    GenerateAnswerResponse,
    GenerateAnswerChunk,
)
from .document_pb2_grpc import (
    DocumentAIServicer,
    add_DocumentAIServicer_to_server,
)


class DocumentService(DocumentAIServicer):

    def ProcessDocument(self, request, context):

        document = DocumentInfo(
            id=request.document_id,
            bucket=request.bucket,
            object_key=request.object_key,
            mime_type=request.mime_type,
        )

        processor = DocumentProcessor()

        try:
            processor.process(document)
        except Exception:
            import traceback
            traceback.print_exc()
            raise

        return ProcessDocumentResponse(
            accepted=True,
        )


    def EmbedQuery(self, request, context):

        embedding_service = get_embedding_service()

        vector = embedding_service.embed_text(

            request.text,

        )

        return EmbedQueryResponse(
            vector=vector,
            model=embedding_service.model,
        )


    def GenerateAnswer(self, request, context):

        chat_service = get_chat_service()

        chunks = [
            {
                "document_name": chunk.document_name,
                "page": chunk.page,
                "section": chunk.section,
                "text": chunk.text,
            }
            for chunk in request.context
        ]

        try:
            answer = chat_service.generate(
                request.question,
                chunks,
            )
        except Exception as e:
            context.abort(grpc.StatusCode.INTERNAL, f"Answer generation failed: {e}")
            return

        return GenerateAnswerResponse(
            answer=answer,
            model=chat_service.model,
        )


    def GenerateAnswerStream(self, request, context):

        chat_service = get_chat_service()

        chunks = [
            {
                "document_name": chunk.document_name,
                "page": chunk.page,
                "section": chunk.section,
                "text": chunk.text,
            }
            for chunk in request.context
        ]

        try:
            for delta in chat_service.generate_stream(
                request.question,
                chunks,
            ):
                yield GenerateAnswerChunk(
                    text_delta=delta,
                    done=False,
                )
        except Exception as e:
            context.abort(grpc.StatusCode.INTERNAL, f"Answer generation failed: {e}")
            return

        yield GenerateAnswerChunk(
            done=True,
            model=chat_service.model,
        )


def create_server():
    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=10)
    )

    add_DocumentAIServicer_to_server(
        DocumentService(),
        server,
    )

    server.add_insecure_port("[::]:50051")

    return server