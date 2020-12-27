package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/simple/proto"
)

func main() {
	conn, err := grpc.Dial("192.168.86.122:443", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Bad Dial: %v", err)
	}

	defer conn.Close()
	client := pb.NewGreetingServiceClient(conn)
	res, err := client.GetGreeting(context.Background(), &pb.GetGreetingRequest{})
	log.Printf("RES %v, %v", res, err)
}
