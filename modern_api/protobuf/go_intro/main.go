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

func createComplex() *myProto.Complex {
	return &myProto.Complex{
		OneDummy: &myProto.Dum{Id: 42, Name: "First name. (one)"},
		MultipleDummies: []*myProto.Dum{
			{Id: 62, Name: "Next name. (62)"},
			{Id: 73, Name: "Third name. (73)"},
		},
	}
}

func createEyeEnum() *myProto.Enumeration {
	return &myProto.Enumeration{
		//EyeColor: 1,
		EyeColor: myProto.EyeColor_EYE_COLOR_GREEN,
	}
}

func main() {
	fmt.Println("Hello proto(c) world! ")
	fmt.Println()

	fmt.Println("createSimple(); ", createSimple())
	fmt.Println("createComplex(); ", createComplex())
	fmt.Println("createEyeEnum(); ", createEyeEnum())
}
