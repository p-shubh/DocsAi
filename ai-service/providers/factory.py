import os

from providers.base import BaseEmbeddingService, BaseChatService
from config.settings import Settings


def get_embedding_service() -> BaseEmbeddingService:
    provider = Settings.EMBEDDING_PROVIDER.lower()

    providers = {
        "gemini": create_gemini,
        "openai": create_openai,
        "ollama": create_ollama,
    }

    if provider not in providers:
        raise ValueError(f"Unsupported embedding provider: {provider}")

    return providers[provider]()


def create_gemini() -> BaseEmbeddingService:
    from providers.gemini.embedding import GeminiEmbeddingService

    return GeminiEmbeddingService()


def create_openai() -> BaseEmbeddingService:
    from providers.openai.embedding import OpenAIEmbeddingService

    return OpenAIEmbeddingService()


def create_ollama() -> BaseEmbeddingService:
    from providers.ollama.embedding import OllamaEmbeddingService

    return OllamaEmbeddingService()


def get_chat_service() -> BaseChatService:
    provider = Settings.CHAT_PROVIDER.lower()

    providers = {
        "gemini": create_gemini_chat,
        "ollama": create_ollama_chat,
    }

    if provider not in providers:
        raise ValueError(f"Unsupported chat provider: {provider}")

    return providers[provider]()


def create_gemini_chat() -> BaseChatService:
    from providers.gemini.chat import GeminiChatService

    return GeminiChatService()


def create_ollama_chat() -> BaseChatService:
    from providers.ollama.chat import OllamaChatService

    return OllamaChatService()