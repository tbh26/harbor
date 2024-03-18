package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/tbh26/harbor/modern_api/grpc/go_intro/greet/proto"
)

func doGreet(client pb.GreetServiceClient, name string) {
	log.Printf("invoked doGreet()... \n")
	r, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: name,
	})
	if err != nil {
		log.Fatalf("doGreet failure; %v \n", err)
	}
	result := r.Result
	fmt.Printf("result: %q  (response; %v) \n", result, r)
}
