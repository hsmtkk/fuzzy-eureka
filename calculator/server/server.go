package main

import (
	"context"
	"log"
	"net"

	"github.com/hsmtkk/fuzzy-eureka/calculator/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSumServiceServer
}

func (s *server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	firstInt := req.GetFirstInt()
	secondInt := req.GetSecondInt()
	ans := firstInt + secondInt
	resp := &pb.SumResponse{
		Result: ans,
	}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
