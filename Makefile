proto-python:
	python -m grpc_tools.protoc \
		-I proto \
		--python_out=ai-service/proto \
		--grpc_python_out=ai-service/proto \
		proto/document.proto

	sed -i '' \
	's/^import document_pb2 as/from . import document_pb2 as/' \
	ai-service/proto/document_pb2_grpc.py