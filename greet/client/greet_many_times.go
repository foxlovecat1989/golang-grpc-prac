package main

import (
	"context"
	pb "github.com/foxlovecat1989/golang-grpc-prac/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) error {
	log.Printf("Invoking GreetManyTimes RPC")
	req := &pb.GreetRequest{Name: "Ed"}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Failed to receive response: %v", err)
		}

		log.Printf("Result: %s", res.Result)
	}

	return nil
}
