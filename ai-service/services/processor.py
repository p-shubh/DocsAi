from loaders.factory import LoaderFactory
from models.document_info import DocumentInfo
from models.processing_result import ProcessingResult
from providers.factory import get_embedding_service
from services.chunker import Chunker
from storage.minio import MinioStorage
from clients.callback import CallbackClient
from config.settings import Settings


class DocumentProcessor:

    def __init__(self):

        self.storage = MinioStorage()

        self.chunker = Chunker()

        self.embedding_service = get_embedding_service()

        self.callback = CallbackClient(
            host=Settings.CALLBACK_HOST,
            port=Settings.CALLBACK_PORT,
        )

    def process(
        self,
        document: DocumentInfo,
    ) -> ProcessingResult:

        # Download the document from MinIO
        local_path = self.storage.download(
            bucket=document.bucket,
            object_name=document.object_key,
        )

        print(f"Downloaded: {local_path}")

        # Select the appropriate loader
        loader = LoaderFactory.get_loader(local_path)

        # Extract document contents
        loaded_document = loader.load(local_path)

        # Generate chunks
        chunks = self.chunker.chunk(loaded_document)

        print("=" * 80)
        print("Chunks")
        print("=" * 80)
        print(f"Generated {len(chunks)} chunks")

        for chunk in chunks:
            print("-" * 80)
            print(f"Chunk #{chunk.id}")
            print(f"Length: {len(chunk.text)}")
            print(chunk.text[:200])

        # Generate embeddings
        embeddings = self.embedding_service.embed(chunks)

        print("=" * 80)
        print("Embeddings")
        print("=" * 80)
        print(f"Generated {len(embeddings)} embeddings")

        if embeddings:
            print(f"Vector dimensions: {len(embeddings[0].vector)}")
            print("First 10 values:")
            print(embeddings[0].vector[:10])

        # Build processing result
        processing_result = ProcessingResult(
            document=loaded_document,
            chunks=chunks,
            embeddings=embeddings,
        )

        # Send result back to the backend
        response = self.callback.save_processing_result(
            document_id=document.id,
            result=processing_result,
        )

        print("=" * 80)
        print("Callback")
        print("=" * 80)
        print(f"Success: {response.success}")

        return processing_result