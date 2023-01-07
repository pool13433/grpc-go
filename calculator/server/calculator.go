package main

import (
	"context"
	pb "github.com/pool13433/grpc-go/calculator/proto"
	"log"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Calculator was invoked with %v\n", in)
	return &pb.SumResponse{Result: in.FirstNum + in.SecondNum}, nil
}