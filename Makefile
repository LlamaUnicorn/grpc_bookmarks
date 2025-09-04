LOCAL_BIN := $(CURDIR)/bin

BUF_BIN := $(LOCAL_BIN)/buf

.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install github.com/bufbuild/buf/cmd/buf@v1.57.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.buf-generate:
	$(info run buf generate...)
	PATH="$(LOCAL_BIN):$(PATH)" $(BUF_BIN) generate

generate: .buf-generate

.PHONY: \
	.bin-deps