package main

import (
	"log"
	"os"
	"strconv"

	pb "github.com/tbh26/harbor/modern_api/grpc/go_intro/calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "0.0.0.0:40042"

func main() {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	c, err := grpc.Dial(address, insecureCredentials)
	if err != nil {
		log.Fatalf("connect failed: %v\n", err)
	}
	defer c.Close()
	client := pb.NewCalculatorServiceClient(c)
	var op1 int32 = 12
	if len(os.Args) > 1 {
		arg1 := os.Args[1]
		firstOp, err := strconv.ParseInt(arg1, 10, 32)
		if err != nil {
			log.Fatalf("argument conversie failed:  %v\n", err)
		}
		op1 = int32(firstOp)
	}
	var op2 int32 = 43
	if len(os.Args) > 2 {
		arg2 := os.Args[2]
		nextOp, err := strconv.ParseInt(arg2, 10, 32)
		if err != nil {
			log.Fatalf("argument conversie failed:  %v\n", err)
		}
		op2 = int32(nextOp)
	}
	doSum(client, op1, op2)
}
