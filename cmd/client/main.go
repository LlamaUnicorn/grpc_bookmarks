package main

import (
	"context"
	"grpc_bookmarks/pkg/api/bookmark"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	conn, err := grpc.NewClient(":8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := bookmark.NewBookmarkClient(conn)

	ctx := context.Background()

	req := &bookmark.CreateBookmarkRequest{
		Title: "test bookmark",
		//Url:   "",
		Url: "https://ya.ru",
		Tag: "search",
	}

	// client context
	cctx := metadata.NewOutgoingContext(ctx, metadata.Pairs("client-header-key", "val"))

	var headers, trailers = metadata.MD{}, metadata.MD{}
	resp, err := client.CreateBookmark(cctx, req,
		grpc.Header(&headers),
		grpc.Trailer(&trailers),
	)
	if err != nil {
		switch status.Code(err) {
		case codes.InvalidArgument:
			log.Println("некорректный запрос")
		default:
			log.Fatal(err)
		}

		if st, ok := status.FromError(err); ok {
			log.Println("code", st.Code(), "details", st.Details(), "message", st.Message())
		} else {
			log.Println("not grpc")
		}
	}

	log.Println("headers:", headers, "trailers:", trailers)
	log.Println(resp.GetBookmarkId())

	// Use protojson.Marshal() instead of a stlib json.Marshal()
	bytes, err := protojson.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bytes))
}
