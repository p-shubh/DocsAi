from dotenv import load_dotenv

load_dotenv()

from google import genai

from config.settings import Settings


client = genai.Client(
    api_key=Settings.GEMINI_API_KEY,
)

response = client.models.generate_content(
    model="gemini-2.5-flash",
    contents="Reply with exactly: Gemini connection successful",
)

print(response.text)