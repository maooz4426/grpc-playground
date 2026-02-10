package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/maooz4426/grpc-playground/greeting/internal/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetingServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponce, error) {
	log.Printf("Received: %v", req.GetName())

	message := fmt.Sprintf("Hello, %s! Welcome to gRPC!", req.GetName())

	return &pb.HelloResponce{
		Message: message,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreetingServiceServer(s, &server{})

	log.Println("Server is running on port 50051...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
