package main

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	keygen "github.com/talhahavadar/microservices/keygen-service/proto"
)

type keygenServiceServer struct {
}

func (s *keygenServiceServer) Generate(ctx context.Context, req *keygen.KeygenRequest) (*keygen.KeygenResponse, error) {
	hash := sha1.New()
	seed := strings.TrimSpace(req.GetSeed())
	hash.Write([]byte(req.GetLongurl()))
	if len(seed) != 0 {
		// That means we need to create new short url using this seed.
		hash.Write([]byte(seed))
	}
	hashedURL := hash.Sum(nil)
	log.Printf("Hashed url is: %x\n", hashedURL)
	encodedURL := base64.URLEncoding.EncodeToString(hashedURL)
	log.Printf("Encoded(hashed) url is: %s\n", encodedURL)
	//encodedURL = base64.URLEncoding.EncodeToString([]byte(req.GetLongurl()))
	//log.Printf("Encoded url is: %s\n", encodedURL)
	return &keygen.KeygenResponse{Shorturl: encodedURL[:6]}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":51051")
	if err != nil {
		log.Fatalf("error occured while listening the server on port 51051. %v", err)
	}
	grpcServer := grpc.NewServer()
	keygen.RegisterKeygenServiceServer(grpcServer, &keygenServiceServer{})
	log.Println("Server is started on :51051")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("error occured while serving. %v", err)
		return
	}
}
