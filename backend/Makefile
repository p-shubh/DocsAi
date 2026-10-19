.PHONY: proto

PROTO_DIR=../proto
GO_OUT=proto/documentpb
PY_OUT=../ai-service/grpc

proto:
	mkdir -p $(PY_OUT)

	protoc \
		-I=$(PROTO_DIR) \
		--go_out=$(GO_OUT) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/document.proto

	cd ../ai-service && \
	python -m grpc_tools.protoc \
		-I ../proto \
		--python_out=grpc \
		--grpc_python_out=grpc \
		../proto/document.proto