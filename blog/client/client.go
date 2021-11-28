package main

import (
	"context"
	"fmt"
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
	if err := createBlog(clt); err != nil {
		log.Fatal(err)
	}
}

func createBlog(clt blog.BlogServiceClient) error {
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
		return fmt.Errorf("failed to create blog; %v", err)
	}
	log.Printf("created blog: %v", resp)
	return nil
}
