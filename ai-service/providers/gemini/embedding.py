from google.genai import types

from providers.gemini.client import GeminiClient

from models.chunk import Chunk
from models.embedding import Embedding

from providers.base import BaseEmbeddingService


class GeminiEmbeddingService(BaseEmbeddingService):

    def __init__(self):
        self.client = GeminiClient().client

    def embed(
        self,
        chunks: list[Chunk],
    ) -> list[Embedding]:

        embeddings = []

        for chunk in chunks:

            response = self.client.models.embed_content(
                model="gemini-embedding-001",
                contents=chunk.text,
                config=types.EmbedContentConfig(
                    task_type="RETRIEVAL_DOCUMENT",
                ),
            )

            embeddings.append(
                Embedding(
                    chunk_id=chunk.id,
                    vector=response.embeddings[0].values,
                    model="gemini-embedding-001",
                )
            )

        return embeddings