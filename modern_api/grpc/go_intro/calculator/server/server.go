package main

import pb "github.com/tbh26/harbor/modern_api/grpc/go_intro/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
