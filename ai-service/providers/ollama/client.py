from typing import Iterator

import ollama

from config.settings import Settings


class OllamaClient:

    def __init__(self):
        self.client = ollama.Client(
            host=Settings.OLLAMA_HOST,
        )

    def embedding(
        self,
        text: str,
    ) -> list[float]:

        try:
            response = self.client.embed(
                model=Settings.OLLAMA_EMBEDDING_MODEL,
                input=text,
            )

            return response["embeddings"][0]

        except Exception as e:
            raise RuntimeError(
                f"Failed to generate Ollama embedding: {e}"
            ) from e

    def chat(
        self,
        prompt: str,
    ) -> str:

        try:
            response = self.client.chat(
                model=Settings.OLLAMA_CHAT_MODEL,
                messages=[
                    {
                        "role": "user",
                        "content": prompt,
                    }
                ],
                think=False,
                options={
                    "num_ctx": Settings.OLLAMA_NUM_CTX,
                },
            )

            return response["message"]["content"]

        except Exception as e:
            raise RuntimeError(
                f"Failed to generate Ollama chat response: {e}"
            ) from e

    def chat_stream(
        self,
        prompt: str,
    ) -> Iterator[str]:

        try:
            stream = self.client.chat(
                model=Settings.OLLAMA_CHAT_MODEL,
                messages=[
                    {
                        "role": "user",
                        "content": prompt,
                    }
                ],
                think=False,
                stream=True,
                options={
                    "num_ctx": Settings.OLLAMA_NUM_CTX,
                },
            )

            for chunk in stream:
                content = chunk["message"]["content"]
                if content:
                    yield content

        except Exception as e:
            raise RuntimeError(
                f"Failed to stream Ollama chat response: {e}"
            ) from e