package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/tbh26/harbor/modern_api/grpc/go_intro/greet/proto"
)

var address string = "0.0.0.0:40041"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen failed (%q); %v \n", address, err)
	}
	log.Printf("listing on; %q \n", address)
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(l); err != nil {
		log.Fatalf("failed to serve; %v \n", err)
	}
}
