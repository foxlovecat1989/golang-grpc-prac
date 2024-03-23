package main

import (
	"context"
	pb "github.com/foxlovecat1989/golang-grpc-prac/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Result: in.A + in.B,
	}, nil
}
