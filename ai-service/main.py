from dotenv import load_dotenv

load_dotenv()

from proto.server import create_server


def main():
    server = create_server()

    server.start()

    print("gRPC server listening on :50051")

    server.wait_for_termination()


if __name__ == "__main__":
    main()