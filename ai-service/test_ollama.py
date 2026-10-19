from providers.ollama.client import OllamaClient

client = OllamaClient()

vector = client.embedding("Hello DocsAI!")

print(type(vector))
print(len(vector))
print(vector[:10])