from dataclasses import dataclass, field


@dataclass
class Chunk:

    id: int

    text: str

    source: str = ""

    page: int = 0

    section: str = ""

    metadata: dict = field(default_factory=dict)