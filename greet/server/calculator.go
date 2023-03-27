package main

import (
	"context"
	"io"
	"log"

	pb "rpc/greet/proto"
)

func add(rhs, lhs int64) int64 {
	return rhs + lhs
}

func (s *Server) Add(ctx context.Context, in *pb.RequestOperand) (*pb.ResponseResult, error) {

	log.Printf("Add called : %v\n", in)

	return &pb.ResponseResult{
		Result: add(in.Lhs, in.Rhs),
	}, nil
}

func (s *Server) Prime(req *pb.RequestPrime, stream pb.CalculatorService_PrimeServer) error {

	log.Printf("Prime called : %v\n", req)

	var k int64 = 2

	res := &pb.ResponseResult{
		Result: k,
	}

	for req.Number > 1 {
		if req.Number%k == 0 {
			res.Result = k
			if err := stream.Send(res); err != nil {
				return err
			}
			req.Number = req.Number / k
		} else {
			k = k + 1
		}
	}

	return nil
}

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {

	log.Printf("Average called\n")

	var sum int64 = 0
	var count int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {

			var result float64 = 0

			if count > 0 {
				result = float64(sum) / float64(count)
			}

			return stream.SendAndClose(&pb.ResponseDouble{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("failed to receive: %v", err)
			return err
		}

		sum += req.Number
		count += 1
	}

	return nil
}
