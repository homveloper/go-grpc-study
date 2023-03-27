package main

import (
	"log"
	pb "rpc/greet/proto"
)

func (s *Server) GreetManyTimes(req *pb.RequestGreet, stream pb.GreetService_GreetManyTimesServer) error {

	log.Printf("GreetManyTimes called : %v\n", req)

	for i := 0; i < 10; i++ {

		res := &pb.ResponseGreet{
			Result: "Hello " + req.FirstName,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}
