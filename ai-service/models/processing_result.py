from dataclasses import dataclass

from models.chunk import Chunk
from models.document import Document
from models.embedding import Embedding


@dataclass
class ProcessingResult:

    document: Document

    chunks: list[Chunk]

    embeddings: list[Embedding]