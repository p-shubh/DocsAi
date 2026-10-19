import os

from minio import Minio

from config.settings import Settings


class MinioStorage:

    def __init__(self):

        self.client = Minio(
            Settings.MINIO_ENDPOINT,
            access_key=Settings.MINIO_ACCESS_KEY,
            secret_key=Settings.MINIO_SECRET_KEY,
            secure=Settings.MINIO_SECURE,
        )

        os.makedirs(
            Settings.TEMP_DIRECTORY,
            exist_ok=True,
        )

    def download(
        self,
        bucket: str,
        object_name: str,
    ) -> str:

        filename = os.path.basename(object_name)

        destination = os.path.join(
            Settings.TEMP_DIRECTORY,
            filename,
        )

        self.client.fget_object(
            bucket,
            object_name,
            destination,
        )

        return destination