package main

import (
	"context"
	"log"
	"time"

	pb "github.com/maooz4426/grpc-playground/greeting/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetingServiceClient(conn)

	callSayHello(client, "Taro")
}

func callSayHello(client pb.GreetingServiceClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.HelloRequest{
		Name: name,
	}

	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}
	log.Printf("Response: %s", res.GetMessage())
}
