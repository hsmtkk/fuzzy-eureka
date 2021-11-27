package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/hsmtkk/fuzzy-eureka/calculator/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	clt := pb.NewCalcServiceClient(conn)

	//doUnary(clt)
	//doServerStreaming(clt)
	//doClientStreaming(clt)
	doBidirStreaming(clt)
}

func doUnary(clt pb.CalcServiceClient) {
	req := &pb.SumRequest{
		FirstInt:  10,
		SecondInt: 20,
	}

	resp, err := clt.Sum(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	ans := resp.GetResult()
	log.Printf("%d + %d = %d\n", 10, 20, ans)
}

func doServerStreaming(clt pb.CalcServiceClient) {
	req := &pb.PrimeNumberDecompositionRequest{
		Number: 210,
	}
	resp, err := clt.DecomposePrimeNumber(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		prime, err := resp.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(prime)
	}
}

func doClientStreaming(clt pb.CalcServiceClient) {
	stream, err := clt.Average(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	stream.Send(&pb.AverageRequest{Number: 3})
	stream.Send(&pb.AverageRequest{Number: 2})
	stream.Send(&pb.AverageRequest{Number: 1})
	stream.Send(&pb.AverageRequest{Number: 4})
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp.GetAverage())
}

func doBidirStreaming(clt pb.CalcServiceClient) {
	stream, err := clt.FindMaximum(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	reqs := []*pb.FindMaximumRequest{
		{Number: 3},
		{Number: 1},
		{Number: 5},
		{Number: 2},
		{Number: 6},
	}
	for _, req := range reqs {
		log.Printf("sending %d", req.GetNumber())
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
		resp, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("current maximum %d", resp.GetCurrentMaximum())
	}
}
