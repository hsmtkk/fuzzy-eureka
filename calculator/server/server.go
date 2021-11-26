package main

import (
	"context"
	"log"
	"net"

	"github.com/hsmtkk/fuzzy-eureka/calculator/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalcServiceServer
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

func (s *server) DecomposePrimeNumber(req *pb.PrimeNumberDecompositionRequest, stream pb.CalcService_DecomposePrimeNumberServer) error {
	number := req.GetNumber()
	var k int64 = 2
	for {
		if number <= 1 {
			break
		}
		if number%k == 0 {
			resp := &pb.PrimeNumberDecompositionResponse{
				Prime: k,
			}
			if err := stream.Send(resp); err != nil {
				return err
			}
			number /= k
		} else {
			k += 1
		}
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
