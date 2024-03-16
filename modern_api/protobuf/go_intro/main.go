package main

import (
	"fmt"
	myProto "github.com/tbh26/harbor/modern_api/protobuf/go_intro/proto"
	"google.golang.org/protobuf/proto"
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

func printOneOf(item interface{}) {
	switch t := item.(type) {
	case *myProto.Result_Id:
		fmt.Printf(" - This item has an Id: %d\n", item.(*myProto.Result_Id).Id)
	case *myProto.Result_Message:
		fmt.Printf(" - This item has an Message: %s\n", item.(*myProto.Result_Message).Message)
	default:
		fmt.Printf(" - item has unexpected type: %T  {{%v}} \n", t, item)
	}
}

func createMap() *myProto.MapExample {
	message := &myProto.MapExample{
		Ids: map[string]*myProto.IdWrapper{
			"my_id":  {Id: 42},
			"my_id2": {Id: 84},
			"my_id3": {Id: 333},
		},
	}
	return message
}

func useFile(p proto.Message, path string) {

	_ = writeToFile(path, p)
	sm2 := &myProto.Simple{}
	_ = readFromFile(path, sm2)
	fmt.Println("content read;", sm2)
}

func main() {
	fmt.Println("Hello proto(c) world! ")
	fmt.Println()

	fmt.Println("createSimple(); ", createSimple())
	fmt.Println("createComplex(); ", createComplex())
	fmt.Println("createEyeEnum(); ", createEyeEnum())

	fmt.Println("printOneOf():  ...  (next line) ")
	printOneOf(&myProto.Result_Id{Id: 42})
	fmt.Println("printOneOf():  ...  (next line) ")
	printOneOf(&myProto.Result_Message{Message: "Hello one_of world!"})
	fmt.Println("printOneOf():  ...  (next line) ")
	printOneOf("Hello world?  (some filler text) ")

	fmt.Println("createMap; ", createMap())
	fmt.Println()

	filePath := "simple.bin"
	useFile(createSimple(), filePath)
}
