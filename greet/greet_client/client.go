package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/hsmtkk/fuzzy-eureka/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	clt := greetpb.NewGreetServiceClient(conn)
	doUnary(clt)
	doServerStreaming(clt)
}

func doUnary(clt greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{FirstName: "Alice", LastName: "Bravo"},
	}
	resp, err := clt.Greet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	result := resp.GetResult()
	fmt.Println(result)
}

func doServerStreaming(clt greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Charlie",
			LastName:  "Delta",
		},
	}
	resp, err := clt.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		log.Print(msg)
	}
}
