package main

import (
	"context"
	pb "github.com/foxlovecat1989/golang-grpc-prac/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was called with %v", in)

	return &pb.GreetResponse{
		Result: "Hello, " + in.Name + "!",
	}, nil
}
