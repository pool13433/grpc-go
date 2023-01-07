package main

import (
	"context"
	pb "github.com/pool13433/grpc-go/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	params := &pb.SumRequest{FirstNum: 100, SecondNum: 200}
	r, err := c.Sum(context.Background(), params)

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Sum (%d + %d) : %d\n", params.FirstNum, params.SecondNum, r.Result)
}