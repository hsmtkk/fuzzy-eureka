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

	doUnary(clt)
	doServerStreaming(clt)
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