package main

import (
	"context"
	"log"
	"os"

	pb "github.com/dgoeke/golang-rpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:8001"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("The server greeted you with: %s\n", response.Message)
}
