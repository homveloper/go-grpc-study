package main

import (
	"context"
	"log"
	pb "rpc/greet/proto"
	"time"
)

func DoLongGreet(client pb.GreetServiceClient) {

	requests := []*pb.RequestGreet{
		{
			FirstName: "John",
		},
		{
			FirstName: "Paul",
		},
		{
			FirstName: "George",
		},
		{
			FirstName: "Ringo",
		},
	}

	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("failed to call LongGreet: %v", err)
	}

	for _, req := range requests {
		log.Printf("Sending request: %v \n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("failed to receive response: %v", err)
	}

	log.Printf("LongGreet response: %s \n", resp)
}
