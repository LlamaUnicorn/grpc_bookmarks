# Install

1. go mod init github.com/username/repo
2. mkdir -p ./protobuf/api/bookmark
   1. create bookmark_service.proto and messages.proto
3. touch Makefile
4. make .bin-deps
5. ./bin/buf config init
6. touch buf.gen.yaml
   either update paths and include protobuf in imports: 
   ```
    inputs:
    - directory: .
      paths:
        - protobuf # путь до наших protobuf файлов
   ```
   or edit buf.yaml and import without protobuf:
   ```
   modules:
   - path: protobuf
   ```
7. make generate
8. mkdir -p ./cmd/bookmark
9. touch ./cmd/bookmark/main.go
10. start server, service and listener
11. implement Bookmark methods Create and List in main.go
12. create /cmd/client.go
13. add [json options] to protobuf/messages.proto 
14. regenerate new code
15. Validation: protovalidate plugin 
    1. Import validate.proto
    2. Create new Makefile vendor.proto.mk .vendor-protovalidate
    3. buf.yaml: - path: vendor.protobuf
    4. Update makefile: include vendor.proto.mk
    5. make vendor (or you can run 'vendor' in editor)
    6. Add validation to fields
    7. make generate

# Part 4: error handling

1. Add errors and error codes in cmd/bookmark/main.go: 
   1. Default: status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
   2. Below is more detailed option
2. Add errors in cmd/client/main.go: switch status.Code(err){}. Check st.Details documentation in-code.
3. Send invalid url to check the error handling (client/main.goj)

# Part 5: middleware

1. Add header := metadata.Pairs("header-key", "val") to server
2. Add ctx, header, trailer and retrieve header in client
   1. Send header in client
   2. Update server to handle client headers