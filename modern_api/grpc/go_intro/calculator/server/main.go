package main

import (
	"log"
	"net"

	pb "github.com/tbh26/harbor/modern_api/grpc/go_intro/calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var address string = "0.0.0.0:40042"

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("listen failed: %v\n", err)
	}

	log.Printf("listening at %s\n", address)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("serve failed: %v\n", err)
	}
}
