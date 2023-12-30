package main

import (
	"context"
	"log"

	pb "github.com/nathanolah/grpc-go-practice/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Printf("doAvg was invoked")

	// reqs := []*pb.AvgRequest{
	// 	{Number: 2},
	// 	{Number: 8},
	// 	{Number: 32},
	// }

	// stream, err := c.Avg(context.Background())

	// if err != nil {
	// 	log.Fatalf("Error while opening the stream: %v\n", err)
	// }

	// for _, req := range reqs {
	// 	log.Printf("Sending req: %v\n", req)
	// 	stream.Send(req)
	// 	time.Sleep(1 * time.Second)
	// }

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)

		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Result)
}
