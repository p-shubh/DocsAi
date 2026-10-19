import os


class Settings:

    MINIO_ENDPOINT = os.getenv(
        "MINIO_ENDPOINT",
        "localhost:9000",
    )

    MINIO_ACCESS_KEY = os.getenv(
        "MINIO_ACCESS_KEY",
        "docsai",
    )

    MINIO_SECRET_KEY = os.getenv(
        "MINIO_SECRET_KEY",
        "docsai123",
    )

    MINIO_BUCKET = os.getenv(
        "MINIO_BUCKET",
        "documents",
    )

    MINIO_SECURE = (
        os.getenv(
            "MINIO_USE_SSL",
            "false",
        ).lower() == "true"
    )

    TEMP_DIRECTORY = os.getenv(
        "TEMP_DIRECTORY",
        "tmp",
    )

    GEMINI_API_KEY = os.getenv(
    "GEMINI_API_KEY",
    "",
    )

    CALLBACK_HOST = os.getenv(
    "CALLBACK_HOST",
    "localhost",
    )

    CALLBACK_PORT = int(
        os.getenv(
            "CALLBACK_PORT",
            "50052",
        )
    )

    OLLAMA_HOST = os.getenv(
    "OLLAMA_HOST",
    "http://localhost:11434",
    )

    OLLAMA_EMBEDDING_MODEL = os.getenv(
        "OLLAMA_EMBEDDING_MODEL",
        "qwen3-embedding:0.6b",
    )

    CHAT_PROVIDER = os.getenv(
    "CHAT_PROVIDER",
    "gemini",
    )

    EMBEDDING_PROVIDER = os.getenv(
        "EMBEDDING_PROVIDER",
        "gemini",
    )

    OLLAMA_CHAT_MODEL = os.getenv(
        "OLLAMA_CHAT_MODEL",
        "gemma4:e4b",
    )

    OLLAMA_NUM_CTX = int(
        os.getenv(
            "OLLAMA_NUM_CTX",
            "8192",
        )
    )