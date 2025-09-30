# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

gRPC-based bookmark management service built with Go and Protocol Buffers. Uses buf for protobuf management and includes protovalidate for request validation.
It's a follow along tutorial. Original final code of the course located at ./example_plain. My current implementation is in progress. Instead of example Note service I have my Bookmark.

## Architecture

### Directory Structure
- `protobuf/api/bookmark/` - Proto definitions (services and messages)
- `pkg/api/bookmark/` - Generated Go code from protos
- `cmd/bookmark/` - gRPC server implementation
- `cmd/client/` - Example gRPC client
- `vendor.protobuf/` - Vendored proto dependencies (google/api, buf/validate, etc.)

### Core Components

**Server** (`cmd/bookmark/main.go`):
- In-memory bookmark storage with `sync.RWMutex`
- gRPC reflection enabled on port 8085
- Request validation using protovalidate
- Service methods: `CreateBookmark` (implemented), `ListBookmarks` (unimplemented)

**Proto Package** (`github.llamaunicorn.grpc_bookmarks.protobuf.api.bookmark`):
- Uses buf validation constraints (min_len, max_len)
- JSON naming conventions configured via json_name options
- Generates to `pkg/api/bookmark` via `go_package` option

**Client**: Uses `protojson.Marshal()` for JSON serialization (not stdlib json)

## Common Commands

### Build & Generate
```bash
# Install buf and protoc plugins
make .bin-deps

# Generate Go code from protos
make generate

# Vendor proto dependencies (googleapis, protovalidate, etc.)
make vendor
```

### Development
```bash
# Run server
go run cmd/bookmark/main.go

# Run client
go run cmd/client/main.go
```

## Protobuf Workflow

1. Edit proto files in `protobuf/api/bookmark/`
2. Run `make generate` to regenerate Go code
3. Generated files appear in `pkg/api/bookmark/`

### Buf Configuration
- `buf.yaml`: Defines modules (protobuf/, vendor.protobuf/) and linting
- `buf.gen.yaml`: Configures generation with protoc-gen-go and protoc-gen-go-grpc
- `vendor.proto.mk`: Makefile for vendoring external proto dependencies

### Adding Validation
Import `buf/validate/validate.proto` and use field options:
```proto
string title = 1 [(buf.validate.field).string = {min_len: 3, max_len: 256}];
```

## Important Notes
- Always run `make generate` after modifying proto files
- Use `protojson.Marshal()` for proto-to-JSON conversion, not standard library
- Validation happens server-side via protovalidate.Validator.Validate()
- Server uses gRPC reflection for tools like grpcurl