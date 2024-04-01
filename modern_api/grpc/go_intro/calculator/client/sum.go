package main

import (
	"context"
	"log"

	pb "github.com/tbh26/harbor/modern_api/grpc/go_intro/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient, first int32, second int32) {
	log.Println("invoked doSum")
	r, err := c.Sum(context.Background(), &pb.SumRequest{FirstOperand: first, SecondOperand: second})
	if err != nil {
		log.Fatalf("Sum failed: %v\n", err)
	}
	log.Printf("Sum: %d\n", r.Result)
}
