package main

import pb "github.com/pool13433/grpc-go/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}