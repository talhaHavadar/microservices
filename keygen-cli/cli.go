package main

import (
	"context"
	"log"

	keygen "github.com/talhahavadar/microservices/keygen-service/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:51051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connetc grpc: %v\n", err)
	}
	defer conn.Close()

	client := keygen.NewKeygenServiceClient(conn)
	req := keygen.KeygenRequest{
		Longurl: "https://www.youtube.com/watch?v=qQEFpbojWN0",
	}
	res, err := client.Generate(context.Background(), &req)
	if err != nil {
		log.Fatalf("could not generate: %v", err)
		return
	}
	log.Printf("Response: %v\n", res)

	req.Seed = res.GetShorturl()
	res, err = client.Generate(context.Background(), &req)
	if err != nil {
		log.Fatalf("could not generate: %v", err)
		return
	}
	log.Printf("Response: %v\n", res)

	req.Seed = res.GetShorturl()
	res, err = client.Generate(context.Background(), &req)
	if err != nil {
		log.Fatalf("could not generate: %v", err)
		return
	}
	log.Printf("Response: %v\n", res)

}
