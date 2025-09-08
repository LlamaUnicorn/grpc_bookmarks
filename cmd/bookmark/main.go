package main

import (
	"grpc_bookmarks/pkg/api/bookmark"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := grpc.NewServer()

	service := BookmarkService{}

	bookmark.RegisterBookmarkServer(server, service)

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(server)

	log.Println("gRPC server listening on :8082")

	if err := server.Serve(lis); err != io.EOF {
		log.Fatal(err)
	}

}

type BookmarkService struct {
	bookmark.UnimplementedBookmarkServer
}
