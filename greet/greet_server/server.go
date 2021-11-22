package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hsmtkk/fuzzy-eureka/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("greet from %v", req)
	greet := req.GetGreeting()
	firstName := greet.GetFirstName()
	lastName := greet.GetLastName()
	msg := fmt.Sprintf("Hello %s %s", firstName, lastName)
	return &greetpb.GreetResponse{Result: msg}, nil
}

func main() {
	// fmt.Println("Hello World")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
