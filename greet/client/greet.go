package main

import (
	"context"
	pb "github.com/foxlovecat1989/golang-grpc-prac/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("Invoking Greet RPC")
	response, err := c.Greet(context.Background(), &pb.GreetRequest{Name: "Ed"})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	log.Printf("Result: %s", response.Result)
}
