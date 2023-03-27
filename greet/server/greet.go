package main

import (
	"context"
	"log"
	pb "rpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.RequestGreet) (*pb.ResponseGreet, error) {

	log.Printf("Greet called : %v\n", in)

	return &pb.ResponseGreet{
		Result: "Hello " + in.FirstName,
	}, nil
}
