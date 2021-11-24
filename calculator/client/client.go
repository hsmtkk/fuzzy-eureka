package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/hsmtkk/fuzzy-eureka/calculator/pb"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("%s first-int second-int", os.Args[0])
	}
	firstInt, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	secondInt, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	clt := pb.NewSumServiceClient(conn)

	req := &pb.SumRequest{
		FirstInt:  firstInt,
		SecondInt: secondInt,
	}

	resp, err := clt.Sum(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	ans := resp.GetResult()
	log.Printf("%d + %d = %d\n", firstInt, secondInt, ans)
}
