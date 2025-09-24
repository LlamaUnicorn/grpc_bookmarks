package main

import (
	"context"
	"grpc_bookmarks/pkg/api/bookmark"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	conn, err := grpc.NewClient(":8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := bookmark.NewBookmarkClient(conn)

	resp, err := client.CreateBookmark(context.Background(), &bookmark.CreateBookmarkRequest{
		Title: "test bookmark",
		Url:   "https://ya.ru",
		Tag:   "search",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.GetBookmarkId())

	// Use protojson.Marshal() instead of a stlib json.Marshal()
	bytes, err := protojson.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bytes))
}
