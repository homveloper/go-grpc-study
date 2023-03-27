package main

import (
	"io"
	"log"
	pb "rpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {

	log.Printf("LongGreet called\n")

	var result string

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.ResponseGreet{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("failed to receive: %v", err)
			return err
		}

		result += "Hello " + req.FirstName + "\n"
	}
}
