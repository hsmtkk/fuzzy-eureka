package main

import (
	"context"
	"fmt"
	"io"
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

func (s *server) Average(stream pb.CalcService_AverageServer) error {
	nums := []int64{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp := &pb.AverageResponse{
				Average: average(nums),
			}
			stream.SendAndClose(resp)
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to receive response; %w", err)
		}
		num := req.GetNumber()
		nums = append(nums, num)
	}
}

func average(nums []int64) float32 {
	var s int64 = 0
	for _, n := range nums {
		s += n
	}
	return float32(s) / float32(len(nums))
}

func (s *server) FindMaximum(stream pb.CalcService_FindMaximumServer) error {
	var currentMax int64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to receive request; %w", err)
		}
		num := req.GetNumber()
		if num > currentMax {
			currentMax = num
		}
		resp := &pb.FindMaximumResponse{
			CurrentMaximum: int64(currentMax),
		}
		if err := stream.Send(resp); err != nil {
			return fmt.Errorf("failed to send response; %w", err)
		}
	}
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
