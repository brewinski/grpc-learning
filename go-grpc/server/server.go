package main

import (
	"log"
	"net"

	"github.com/brewinski/grcp-learning/go-grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	chatServer := chat.Server{}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &chatServer)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9000: %v", err)
	}
}
