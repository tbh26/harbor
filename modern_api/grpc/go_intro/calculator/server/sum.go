package main

import (
	"context"
	"log"

	pb "github.com/tbh26/harbor/modern_api/grpc/go_intro/calculator/proto"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("invoked Sum with %v\n", in)
	return &pb.SumResponse{Result: in.FirstOperand + in.SecondOperand}, nil
}
