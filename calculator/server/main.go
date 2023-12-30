package main

import (
	"log"
	"net"

	pb "github.com/nathanolah/grpc-go-practice/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// New grpc server started without services registered.
	s := grpc.NewServer()

	// Register the calculator service server
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
