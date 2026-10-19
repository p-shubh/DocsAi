from pathlib import Path

from loaders.pdf import PDFLoader
from loaders.docx import DOCXLoader
from loaders.xlsx import XLSXLoader
from loaders.csv import CSVLoader
from loaders.txt import TXTLoader
from loaders.image import ImageLoader
from loaders.pptx import PPTXLoader


class LoaderFactory:

    @staticmethod
    def get_loader(file_path: str):

        extension = Path(file_path).suffix.lower()

        if extension == ".pdf":
            return PDFLoader()

        if extension == ".docx":
            return DOCXLoader()

        if extension == ".xlsx":
            return XLSXLoader()

        if extension == ".csv":
            return CSVLoader()

        if extension == ".txt":
            return TXTLoader()

        if extension in [".png", ".jpg", ".jpeg", ".bmp", ".tiff"]:
            return ImageLoader()

        if extension == ".pptx":
            return PPTXLoader()

        raise ValueError(f"Unsupported document type: {extension}")