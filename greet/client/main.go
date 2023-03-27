package main

import (
	"log"
	"math/rand"
	"time"

	pb "rpc/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:8080"

type Client struct {
	pb.GreetServiceClient
}

func benchmark(f func()) {
	start := time.Now()
	f()
	log.Printf("took %s", time.Since(start))
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	var arr []int64

	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Int63n(100000))
	}

	benchmark(func() {
		max, err := DoMax(client, arr)
		if err != nil {
			log.Fatalf("failed to max: %v", err)
		}

		log.Printf("max: %d", max)
	})

	// benchmark(func() {
	// 	sqrt, err := DoSqrt(client, -1)
	// 	if err != nil {
	// 		log.Fatalf("failed to sqrt: %v", err)
	// 	}

	// 	log.Printf("sqrt: %d", sqrt)
	// })
}
