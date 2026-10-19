from typing import Iterator

from providers.gemini.client import GeminiClient
from providers.base import BaseChatService
from services.prompt_builder import build_prompt


class GeminiChatService(BaseChatService):

    def __init__(self):
        self.client = GeminiClient().client
        self.model = "gemini-2.5-flash"

    def generate(
        self,
        question: str,
        context: list[dict],
    ) -> str:

        prompt = build_prompt(question, context)

        response = self.client.models.generate_content(
            model=self.model,
            contents=prompt,
        )

        return response.text

    def generate_stream(
        self,
        question: str,
        context: list[dict],
    ) -> Iterator[str]:

        prompt = build_prompt(question, context)

        stream = self.client.models.generate_content_stream(
            model=self.model,
            contents=prompt,
        )

        for chunk in stream:
            if chunk.text:
                yield chunk.text
