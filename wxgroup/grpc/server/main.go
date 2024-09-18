package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc/arith"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedArithmeticServer
}

func (*server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.GetA() + req.GetB()
	return &pb.AddResponse{Result: result}, nil
}

func (*server) StreamSum(stream pb.Arithmetic_StreamSumServer) error {
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		sum += req.GetNumber()
		if err := stream.Send(&pb.SumResponse{Result: sum}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8086")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterArithmeticServer(s, &server{})
	reflection.Register(s) // Enable reflection for development
	log.Println("Starting gRPC server on :8086...")
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
