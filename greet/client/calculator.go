package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "rpc/greet/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DoAdd(client pb.CalculatorServiceClient, lhs, rhs int64) (int64, error) {

	request := &pb.RequestOperand{
		Lhs: lhs,
		Rhs: rhs,
	}

	resp, err := client.Add(context.Background(), request)

	if err != nil {
		log.Fatalf("failed to add: %v", err)
		return 0, err
	}

	log.Printf("Add response: %v", resp)

	return resp.Result, nil
}

func DoPrime(client pb.CalculatorServiceClient, number int64) ([]int64, error) {

	request := &pb.RequestPrime{
		Number: number,
	}

	stream, err := client.Prime(context.Background(), request)

	if err != nil {
		log.Fatalf("failed to prime: %v", err)
		return nil, err
	}

	var result []int64

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to receive: %v", err)
			return nil, err
		}

		log.Printf("Prime response: %v", resp)
		result = append(result, resp.Result)
	}

	return result, nil
}

func DoAverage(client pb.CalculatorServiceClient, numbers []int64) (float64, error) {

	stream, err := client.Average(context.Background())

	if err != nil {
		log.Fatalf("failed to average: %v", err)
		return 0, err
	}

	for _, number := range numbers {
		request := &pb.RequestNumber{
			Number: number,
		}

		if err := stream.Send(request); err != nil {
			log.Fatalf("failed to send: %v", err)
			return 0, err
		}

		time.Sleep(100 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("failed to receive: %v", err)
		return 0, err
	}

	log.Printf("Average response: %v", resp)

	return resp.Result, nil
}

func DoMax(client pb.CalculatorServiceClient, numbers []int64) (int64, error) {

	stream, err := client.Max(context.Background())

	if err != nil {
		log.Fatalf("failed to max: %v", err)
		return 0, err
	}

	waitc := make(chan struct{})

	go func() {
		for _, number := range numbers {
			request := &pb.RequestNumber{
				Number: number,
			}

			if err := stream.Send(request); err != nil {
				log.Fatalf("failed to send: %v", err)
				break
			}
		}

		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("failed to close stream: %v", err)
		}
	}()

	var max int64 = 0

	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("failed to receive: %v", err)
				break
			}

			// log.Printf("Max response: %v", resp)

			max = resp.Number
		}

		close(waitc)
	}()

	<-waitc

	return max, nil
}

func DoSqrt(client pb.CalculatorServiceClient, number int64) (int64, error) {

	request := &pb.RequestNumber{
		Number: number,
	}

	resp, err := client.Sqrt(context.Background(), request)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error From Server Message : %s, Code : %s\n", e.Message(), e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("Error From Server Details : %s\n", e.Details())
			}

		} else {
			log.Printf("Error None gRPC: %v\n", err)
		}

		return 0, err
	}

	log.Printf("Sqrt response: %v", resp)

	return resp.Number, nil
}
