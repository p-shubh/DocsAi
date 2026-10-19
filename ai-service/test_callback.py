from dotenv import load_dotenv

load_dotenv()

from clients.callback import CallbackClient
from config.settings import Settings


client = CallbackClient(
    host=Settings.CALLBACK_HOST,
    port=Settings.CALLBACK_PORT,
)

print("Callback client created successfully.")