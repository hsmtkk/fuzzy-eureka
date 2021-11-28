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

	//createBlog(clt)
	//readBlog(clt)
	//updateBlog(clt)
	deleteBlog(clt)
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

func readBlog(clt blog.BlogServiceClient) error {
	req := &blog.ReadRequest{
		BlogId: "2d15e6f1-74dd-4176-b930-dca03ab84810",
	}
	resp, err := clt.Read(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp)

	req = &blog.ReadRequest{
		BlogId: "hogehoge",
	}
	_, err = clt.Read(context.Background(), req)
	log.Print(err)

	return nil
}

func updateBlog(clt blog.BlogServiceClient) error {
	req := &blog.UpdateRequest{
		Blog: &blog.Blog{
			Id:       "2d15e6f1-74dd-4176-b930-dca03ab84810",
			AuthorId: "alice",
			Content:  "alice content",
			Title:    "alice title",
		},
	}
	resp, err := clt.Update(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp)
	return nil
}

func deleteBlog(clt blog.BlogServiceClient) error {
	req := &blog.DeleteRequest{
		BlogId: "2d15e6f1-74dd-4176-b930-dca03ab84810",
	}
	resp, err := clt.Delete(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp)
	return nil
}
