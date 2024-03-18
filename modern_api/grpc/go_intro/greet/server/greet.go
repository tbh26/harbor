package main

import (
	"context"
	"log"

	pb "github.com/tbh26/harbor/modern_api/grpc/go_intro/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with; %v, context: %v \n", in, ctx)
	name := in.FirstName
	return &pb.GreetResponse{Result: "Hello " + name}, nil
}
