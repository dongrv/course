package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/arith"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient(":8086", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewArithmeticClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Add(ctx, &pb.AddRequest{A: 100, B: 200})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Addition Result: %d", resp.GetResult())

	log.Println("Starting streaming sum...")

	stream, err := client.StreamSum(ctx)
	if err != nil {
		log.Fatalf("could not stream sum: %v", err)
	}

	reqs := []*pb.SumRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
		{Number: 5},
	}

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			log.Fatalf("could not send request: %v", err)
		}
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("could not receive response: %v", err)
		}
		log.Printf("Streamed Sum: %d", resp.GetResult())
	}
}
