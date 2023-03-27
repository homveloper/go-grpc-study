package main

import (
	"context"
	"io"
	"log"
	pb "rpc/greet/proto"
	"time"
)

func DoGreetEveryOne(client pb.GreetServiceClient) {

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
	}

	stream, err := client.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("failed to greet everyone: %v", err)
	}

	wait := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("Sending request: %v \n", req)
			err := stream.Send(req)

			if err != nil {
				log.Fatalf("failed to send request: %v", err)
			}

			time.Sleep(100 * time.Millisecond)
		}

		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("failed to close stream: %v", err)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("failed to receive: %v", err)
			}

			log.Printf("GreetEveryone response: %v", resp)
		}
		close(wait)
	}()

	<-wait

}
