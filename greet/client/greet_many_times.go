package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nathanolah/grpc-go-practice/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTImes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Nathan",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
