from dataclasses import dataclass


@dataclass
class DocumentInfo:

    id: str

    bucket: str

    object_key: str

    mime_type: str