package main

import (
	"context"
	pb "github.com/foxlovecat1989/golang-grpc-prac/calculator/proto"
	"log"
)

func doSum(client pb.SumServiceClient) {
	// Call the remote procedure
	sum, err := client.Sum(context.Background(), &pb.SumRequest{A: 3, B: 5})
	if err != nil {
		return
	}

	// Print the result
	log.Printf("Sum of 3 and 5 is %v", sum.Result)
}
