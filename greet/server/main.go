package main

import (
	"log"
	"net"

	pb "rpc/greet/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:8080"

type Server struct {
	pb.GreetServiceServer
	pb.CalculatorServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listening on %s", addr)

	server := grpc.NewServer()

	pb.RegisterGreetServiceServer(server, &Server{})
	pb.RegisterCalculatorServiceServer(server, &Server{})

	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
