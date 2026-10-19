import os
from models.chunk import Chunk
from models.embedding import Embedding

from providers.base import BaseEmbeddingService
from providers.ollama.client import OllamaClient


class OllamaEmbeddingService(BaseEmbeddingService):

    def __init__(self):
        self.client = OllamaClient()
        self.model = os.getenv("OLLAMA_EMBEDDING_MODEL")

    def embed(
        self,
        chunks: list[Chunk],
    ) -> list[Embedding]:

        embeddings: list[Embedding] = []

        for chunk in chunks:

            vector = self.client.embedding(
                chunk.text,
            )

            embeddings.append(
                Embedding(
                    chunk_id=chunk.id,
                    model=self.model,
                    vector=vector,
                )
            )

        return embeddings
    
    def embed_text(
        self,
        text: str,
    ) -> list[float]:

        return self.client.embedding(text)