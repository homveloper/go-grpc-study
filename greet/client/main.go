package main

import (
	"log"

	pb "rpc/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:8080"

type Client struct {
	pb.GreetServiceClient
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	DoGreetEveryOne(client)
}
