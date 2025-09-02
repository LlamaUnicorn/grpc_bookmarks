LOCAL_BIN := $(CURDIR)/bin

.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install github.com/bufbuild/buf/cmd/buf@v1.57.0


.PHONY: \
	.bin-deps