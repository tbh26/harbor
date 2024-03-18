package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var address string = "0.0.0.0:40041"

func main() {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	c, err := grpc.Dial(address, insecureCredentials)
	if err != nil {
		log.Fatalf("failed to connect (%q); %v \n", address, err)
	}
	defer c.Close()
	//
}
