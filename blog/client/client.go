package main

import (
	"context"
	"log"

	"github.com/hsmtkk/fuzzy-eureka/blog/blog"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	clt := blog.NewBlogServiceClient(conn)

	req := &blog.CreateRequest{
		Blog: &blog.Blog{
			AuthorId: "alpha",
			Content:  "alpha content",
			Title:    "alipha title",
		},
	}
	log.Printf("creating blog: %v", req)
	resp, err := clt.Create(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("created blog: %v", resp)
}
