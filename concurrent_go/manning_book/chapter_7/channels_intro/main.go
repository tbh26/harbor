package main

import (
	"fmt"
	"time"
)

func receiver(tag string, messages chan string) {
	msg := ""
	for msg != "STOP" {
		msg = <-messages
		fmt.Printf("Message [%q] received; %q \n", tag, msg)
	}
	fmt.Printf("End receiver %q \n", tag)
}

func demo() {
	msgChannel := make(chan string)
	go receiver("first channel", msgChannel)
	messages := []string{"Hello...", "channel", "world!", "STOP"}
	for _, message := range messages {
		fmt.Printf("Sending message %q into channel.\n", message)
		msgChannel <- message
	}
	time.Sleep(321 * time.Microsecond)
	fmt.Println("\n - exit demo... \n")
}

func main() {
	fmt.Println("Hello channel(s) world!   📡 ")

	demo()
}
