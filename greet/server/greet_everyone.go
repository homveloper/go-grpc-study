package main

import (
	"io"
	"log"
	pb "rpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	log.Printf("GreetEveryone called\n")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("failed to receive: %v", err)
			return err
		}

		res := &pb.ResponseGreet{
			Result: "Hello " + req.FirstName,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
