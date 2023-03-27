package main

import (
	"context"
	"log"
	pb "rpc/greet/proto"
)

func DoGreet(client pb.GreetServiceClient) {

	request := &pb.RequestGreet{
		FirstName: "John",
	}

	if resp, err := client.Greet(context.Background(), request); err == nil {
		log.Printf("Greet response: %v", resp)
	} else {
		log.Fatalf("failed to greet: %v", err)
	}

}
