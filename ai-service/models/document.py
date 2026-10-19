from dataclasses import dataclass, field


@dataclass
class Document:

    text: str

    page_count: int = 0

    metadata: dict = field(default_factory=dict)