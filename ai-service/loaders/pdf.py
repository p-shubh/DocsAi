from loaders.base import BaseLoader
from models.document import Document


class PDFLoader(BaseLoader):

    def load(self, file_path: str) -> Document:

        return Document(
            text="",
            page_count=0,
            metadata={},
        )