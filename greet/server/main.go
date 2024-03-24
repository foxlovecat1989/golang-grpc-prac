package main

import (
	pb "github.com/foxlovecat1989/golang-grpc-prac/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var address = "localhost:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	// Start the server
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			log.Fatalf("Failed to close listener: %v", err)
		}
	}(listen)
	log.Printf("Server listening at %v", address)

	server := grpc.NewServer()
	pb.RegisterGreetServiceServer(server, &Server{})
	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
