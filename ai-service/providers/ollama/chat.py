from typing import Iterator

from config.settings import Settings
from providers.base import BaseChatService
from providers.ollama.client import OllamaClient
from services.prompt_builder import build_prompt


class OllamaChatService(BaseChatService):

    def __init__(self):
        self.client = OllamaClient()
        self.model = Settings.OLLAMA_CHAT_MODEL

    def generate(
        self,
        question: str,
        context: list[dict],
    ) -> str:

        prompt = build_prompt(question, context)

        return self.client.chat(prompt)

    def generate_stream(
        self,
        question: str,
        context: list[dict],
    ) -> Iterator[str]:

        prompt = build_prompt(question, context)

        yield from self.client.chat_stream(prompt)
