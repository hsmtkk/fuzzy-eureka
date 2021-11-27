package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/hsmtkk/fuzzy-eureka/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Hello %s number %d", firstName, i)
		resp := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(resp)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	results := []string{"Hello"}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := strings.Join(results, " ")
			resp := &greetpb.LongGreetResponse{
				Result: result,
			}
			if err := stream.SendAndClose(resp); err != nil {
				return fmt.Errorf("failed to send response; %w", err)
			}
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to receive request; %w", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		lastName := req.GetGreeting().GetLastName()
		results = append(results, fmt.Sprintf("%s %s!", firstName, lastName))
	}
}

func (s *server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to receive request; %w", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result := fmt.Sprintf("Hello %s !", firstName)
		if err := stream.Send(&greetpb.GreetEveryoneResponse{Result: result}); err != nil {
			return fmt.Errorf("failed to send response; %w", err)
		}
	}
}

func main() {
	// fmt.Println("Hello World")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	certFile := "./selfsignedcert/cert.pem"
	keyFile := "./selfsignedcert/key.pem"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
