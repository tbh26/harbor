package main

import (
	"fmt"
	myProto "github.com/tbh26/harbor/modern_api/protobuf/go_intro/proto"
)

func createSimple() *myProto.Simple {
	return &myProto.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "my first simple proto message (name)",
		SampleLists: []int32{1, 2, 4, 8, 16},
	}
}

func main() {
	fmt.Println("Hello proto(c) world! ")

	fmt.Println(createSimple())
}
