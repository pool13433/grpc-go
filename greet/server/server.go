package main

import pb "github.com/pool13433/grpc-go/greet/proto"

type Server struct {
	pb.GreetServiceServer
}