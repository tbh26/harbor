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

func receiveTwice(tag string, messages chan string) {
	// receive two messages and call it a day

	msg := <-messages
	fmt.Printf("Message [%q] received; %q \n", tag, msg)

	msg = <-messages
	fmt.Printf("Message [%q] received; %q \n", tag, msg)

	fmt.Printf("End receiver %q \n", tag)
}

func firstDemo() {
	msgChannel := make(chan string)
	go receiver("first channel", msgChannel)
	messages := []string{"Hello...", "channel", "world!", "STOP"}
	for _, message := range messages {
		fmt.Printf("Sending message %q into channel.\n", message)
		msgChannel <- message
	}
	time.Sleep(321 * time.Microsecond)
	fmt.Println("\n - exit first demo... \n")
}

func nextDemo() {
	msgChannel := make(chan string)
	go receiveTwice("broken channel receiver", msgChannel)
	messages := []string{"bla", "bla bla", "whatever...", "STOP bla"}
	for _, message := range messages {
		fmt.Printf("Sending message %q into channel.\n", message)
		msgChannel <- message
	}
	time.Sleep(321 * time.Microsecond)
	fmt.Println("\n - exit next/last demo.\n")
}

func demo() {
	firstDemo()

	fmt.Println("\n\t\t=-=-=-=\n")

	nextDemo()
}

func main() {
	fmt.Println("Hello channel(s) world!   📡 ")

	demo()
}
