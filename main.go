package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	pb "github.com/brotherlogic/simple/proto"
	"google.golang.org/grpc"
)

type server struct{}

//GetGreeting simple
func (s *server) GetGreeting(ctx context.Context, req *pb.GetGreetingRequest) (*pb.GetGreetingResponse, error) {
	return &pb.GetGreetingResponse{Greeting: "hello"}, nil
}
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	port := os.Getenv("PORT")
	hport := os.Getenv("HPORT")
	if hport == "" {
		hport = "8082"
	}
	if port == "" {
		port = "8081"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Error in listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetingServiceServer(grpcServer, &server{})

	// Serve grpc
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// Serve http
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(fmt.Sprintf(":%v", hport), nil)
}
