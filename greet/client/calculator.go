package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "rpc/greet/proto"
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
