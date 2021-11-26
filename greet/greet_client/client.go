package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

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
	//doUnary(clt)
	//doServerStreaming(clt)
	//doClientStreaming(clt)
	doBidirStreaming(clt)
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

func doClientStreaming(clt greetpb.GreetServiceClient) {
	stream, err := clt.LongGreet(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	reqs := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Echo",
				LastName:  "Foxtrot",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Golf",
				LastName:  "Hotel",
			},
		},
	}
	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp)
}

func doBidirStreaming(clt greetpb.GreetServiceClient) {
	stream, err := clt.GreetEveryone(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var waitc sync.WaitGroup

	waitc.Add(1)
	go func() {
		defer waitc.Done()
		reqs := []*greetpb.GreetEveryoneRequest{
			{
				Greeting: &greetpb.Greeting{
					FirstName: "Echo",
					LastName:  "Foxtrot",
				},
			},
			{
				Greeting: &greetpb.Greeting{
					FirstName: "Golf",
					LastName:  "Hotel",
				},
			},
		}
		for _, req := range reqs {
			if err := stream.Send(req); err != nil {
				log.Print(err)
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	waitc.Add(1)
	go func() {
		defer waitc.Done()
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Print(err)
			}
			log.Print(resp.GetResult())
		}
	}()

	waitc.Wait()
}
