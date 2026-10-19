from loaders.base import BaseLoader
from models.document import Document


class ImageLoader(BaseLoader):

    def load(self, file_path: str) -> Document:

        return Document(
            text="",
            page_count=0,
            metadata={},
        )