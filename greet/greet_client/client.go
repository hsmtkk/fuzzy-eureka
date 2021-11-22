package main

import (
	"context"
	"fmt"
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
	// log.Print(clt)

	greeting := &greetpb.Greeting{FirstName: "Alice", LastName: "Bravo"}
	req := &greetpb.GreetRequest{Greeting: greeting}
	resp, err := clt.Greet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	result := resp.GetResult()
	fmt.Println(result)
}
