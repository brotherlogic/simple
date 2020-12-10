package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/brotherlogic/simple/proto"
	"google.golang.org/grpc"
)

type server struct{}

//GetGreeting simple
func (s *server) GetGreeting(ctx context.Context, req *pb.GetGreetingRequest) (*pb.GetGreetingResponse, error) {
	return &pb.GetGreetingResponse{Greeting: "hello"}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Error in listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetingServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
