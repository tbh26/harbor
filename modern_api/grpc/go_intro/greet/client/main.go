package main

import (
	pb "github.com/tbh26/harbor/modern_api/grpc/go_intro/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var address string = "0.0.0.0:40041"

func main() {
	insecureCredentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	c, err := grpc.Dial(address, insecureCredentials)
	if err != nil {
		log.Fatalf("failed to connect (%q); %v \n", address, err)
	}
	defer func(c *grpc.ClientConn) {
		err := c.Close()
		if err != nil {
			log.Println("close failed", err)
		}
	}(c)
	client := pb.NewGreetServiceClient(c)
	name := "Bob!"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	doGreet(client, name)
}
