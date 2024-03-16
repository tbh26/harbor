package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln(" - writeToFile, can't serialize to bytes", err)
		return err
	}

	if err = os.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln(" - writeToFile, can't write to file", err)
		return err
	}

	fmt.Printf("writeToFile; data has been written. (%s) \n", filename)
	return nil
}

func readFromFile(filename string, pb proto.Message) error {
	in, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(" - readFromFile, can't read from file", err)
		return err
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln(" - readFromFile, can't deserialize from file", err)
		return err
	}

	fmt.Printf("readFromFile; data is read. (%s) \n", filename)
	return nil
}
