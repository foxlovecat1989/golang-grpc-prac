package main

import (
	pb "github.com/foxlovecat1989/golang-grpc-prac/greet/proto"
	"log"
	"strconv"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was called with %v", in)

	for i := 0; i < 10; i++ {
		err := stream.Send(&pb.GreetResponse{
			Result: "Hello, " + in.Name + "!" + " " + "Times: " + strconv.Itoa(i),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
