package main

import (
	"io"
	"log"

	pb "github.com/nathanolah/grpc-go-practice/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		// if req.Number > maximum {
		// 	maximum = req.Number
		// }

		if number := req.Number; number > maximum {
			maximum = number

			// send to client
			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("Error sending data to client %v\n", err)
			}
		}
	}
}
