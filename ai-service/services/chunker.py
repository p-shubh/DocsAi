from models.chunk import Chunk
from models.document import Document


class Chunker:

    def chunk(
        self,
        document: Document,
        chunk_size: int = 1000,
    ) -> list[Chunk]:

        text = document.text.strip()

        if not text:
            return []

        chunks = []

        start = 0
        chunk_id = 1

        while start < len(text):

            end = start + chunk_size

            chunk_text = text[start:end]

            chunks.append(
                Chunk(
                    id=chunk_id,
                    text=chunk_text,
                )
            )

            chunk_id += 1
            start = end

        return chunks