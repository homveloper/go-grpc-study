package main

import (
	"context"
	"io"
	"log"
	pb "rpc/greet/proto"
)

func DoGreetManyTimes(client pb.GreetServiceClient) {

	log.Println("GreetManyTimes called")

	request := &pb.RequestGreet{
		FirstName: "John",
	}

	stream, err := client.GreetManyTimes(context.Background(), request)

	if err != nil {
		log.Fatalf("failed to greet many times: %v", err)
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to receive: %v", err)
		}

		log.Printf("GreetManyTimes response: %v", resp)
	}
}
