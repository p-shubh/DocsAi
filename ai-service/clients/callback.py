import grpc

from models.processing_result import ProcessingResult

from proto import document_pb2
from proto import document_pb2_grpc


class CallbackClient:

    def __init__(
        self,
        host: str,
        port: int,
    ):

        self.channel = grpc.insecure_channel(
            f"{host}:{port}"
        )

        self.client = document_pb2_grpc.DocumentCallbackStub(
            self.channel
        )

    def save_processing_result(
        self,
        document_id: str,
        result: ProcessingResult,
    ):

        request = document_pb2.SaveProcessingResultRequest(
            document_id=document_id,
        )

        # Processing Result
        request.result.page_count = result.document.page_count

        if isinstance(result.document.metadata, dict) and result.document.metadata:
            request.result.metadata.update(
                {
                    str(key): str(value)
                    for key, value in result.document.metadata.items()
                }
            )

        # Chunks
        for chunk in result.chunks:

            pb_chunk = request.result.chunks.add()
            pb_chunk.id = chunk.id
            pb_chunk.text = chunk.text
            pb_chunk.page = chunk.page
            pb_chunk.section = chunk.section
            if isinstance(chunk.metadata, dict) and chunk.metadata:
                pb_chunk.metadata.update(
                    {
                        str(key): str(value)
                        for key, value in chunk.metadata.items()
                    }
                )

        # Embeddings
        for embedding in result.embeddings:

            pb_embedding = request.result.embeddings.add()
            pb_embedding.chunk_id = embedding.chunk_id
            pb_embedding.vector.extend(embedding.vector)
            pb_embedding.model = embedding.model

        response = self.client.SaveProcessingResult(
            request
        )

        return response