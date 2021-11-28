package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/google/uuid"
	"github.com/hsmtkk/fuzzy-eureka/blog/blog"
	scribble "github.com/nanobox-io/golang-scribble"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const COLLECTION = "blog"

type server struct {
	driver *scribble.Driver
	blog.UnimplementedBlogServiceServer
}

func NewServer(driver *scribble.Driver) *server {
	return &server{driver: driver}
}

func (s *server) Create(ctx context.Context, req *blog.CreateRequest) (*blog.CreateResponse, error) {
	id := uuid.New().String()
	reqBlog := req.GetBlog()
	authorID := reqBlog.GetAuthorId()
	content := reqBlog.GetContent()
	title := reqBlog.GetTitle()

	blogItem := blogItem{
		ID:       id,
		AuthorID: authorID,
		Content:  content,
		Title:    title,
	}

	if err := s.driver.Write(COLLECTION, id, &blogItem); err != nil {
		return nil, fmt.Errorf("failed to write scribble; %s; %w", id, err)
	}

	resp := &blog.CreateResponse{
		Blog: &blog.Blog{
			Id:       id,
			AuthorId: authorID,
			Content:  content,
			Title:    title,
		},
	}
	return resp, nil
}

func (s *server) Read(ctx context.Context, req *blog.ReadRequest) (*blog.ReadResponse, error) {
	blogID := req.GetBlogId()
	item, err := s.getBlogByID(blogID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%w", err)
	}
	resp := &blog.ReadResponse{
		Blog: &blog.Blog{
			Id:       item.ID,
			AuthorId: item.AuthorID,
			Content:  item.Content,
			Title:    item.Title,
		},
	}
	return resp, nil
}

func (s *server) Update(ctx context.Context, req *blog.UpdateRequest) (*blog.UpdateResponse, error) {
	reqBlog := req.GetBlog()
	blogID := reqBlog.GetId()
	item, err := s.getBlogByID(blogID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%w", err)
	}
	item.AuthorID = reqBlog.GetAuthorId()
	item.Content = reqBlog.GetContent()
	item.Title = reqBlog.GetTitle()
	s.driver.Write(COLLECTION, blogID, &item)
	resp := &blog.UpdateResponse{
		Blog: &blog.Blog{
			Id:       item.ID,
			AuthorId: item.AuthorID,
			Content:  item.Content,
			Title:    item.Title,
		},
	}
	return resp, nil
}

func (s *server) getBlogByID(blogID string) (blogItem, error) {
	var item blogItem
	if err := s.driver.Read(COLLECTION, blogID, &item); err != nil {
		return blogItem{}, fmt.Errorf("blog %s is not found; %w", blogID, err)
	}
	return item, nil
}

type blogItem struct {
	ID       string
	AuthorID string
	Content  string
	Title    string
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	driver, err := scribble.New("./db", nil)
	if err != nil {
		log.Fatalf("failed to create scribble database; %w", err)
	}

	s := grpc.NewServer()
	blog.RegisterBlogServiceServer(s, NewServer(driver))

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// wait until ctrl-c
	<-ch

	s.Stop()
	lis.Close()
}
