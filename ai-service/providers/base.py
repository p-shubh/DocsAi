from abc import ABC, abstractmethod
from typing import Iterator

from models.chunk import Chunk
from models.embedding import Embedding


class BaseEmbeddingService(ABC):

    @abstractmethod
    def embed(
        self,
        chunks: list[Chunk],
    ) -> list[Embedding]:
        pass

    @abstractmethod
    def embed_text(
        self,
        text: str,
    ) -> list[float]:
        pass


class BaseChatService(ABC):

    @abstractmethod
    def generate(
        self,
        question: str,
        context: list[dict],
    ) -> str:
        pass

    @abstractmethod
    def generate_stream(
        self,
        question: str,
        context: list[dict],
    ) -> Iterator[str]:
        pass