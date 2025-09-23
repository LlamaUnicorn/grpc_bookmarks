package main

import (
	"context"
	"grpc_bookmarks/pkg/api/bookmark"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	server := grpc.NewServer()

	service := &BookmarkService{storage: make(map[uint64]*Bookmark, 1)}

	bookmark.RegisterBookmarkServer(server, service)

	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(server)

	log.Println("gRPC server listening on :8085")

	if err := server.Serve(lis); err != io.EOF {
		log.Fatal(err)
	}

}

type Bookmark struct {
	ID    uint64
	Title string
	URL   string
	Tag   string
}

type BookmarkService struct {
	bookmark.UnimplementedBookmarkServer

	storage map[uint64]*Bookmark
	mx      sync.RWMutex
}

func (s *BookmarkService) CreateBookmark(ctx context.Context, req *bookmark.CreateBookmarkRequest) (*bookmark.CreateBookmarkResponse, error) {

	id := rand.Uint64()
	bookmarkLocal := &Bookmark{
		ID:    id,
		Title: req.GetTitle(),
		URL:   req.GetUrl(),
		Tag:   req.GetTag(),
	}
	s.mx.Lock()
	s.storage[id] = bookmarkLocal
	s.mx.Unlock()
	return &bookmark.CreateBookmarkResponse{
		BookmarkId: id,
	}, nil
}
func (s *BookmarkService) ListBookmarks(context.Context, *bookmark.ListBookmarksRequest) (*bookmark.ListBookmarksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBookmarks not implemented")
}
