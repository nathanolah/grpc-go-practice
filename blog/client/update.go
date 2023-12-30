package main

import (
	"context"
	"log"

	pb "github.com/nathanolah/grpc-go-practice/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Nathan",
		Title:    "A new title",
		Content:  "Content of the first blog, with some updates.",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error occurred while updating: %v\n", err)
	}

	log.Println("Blog was updated.")
}
