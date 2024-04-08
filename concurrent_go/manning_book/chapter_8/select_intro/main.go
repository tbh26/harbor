package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

func writeEvery(msg string, d time.Duration) <-chan string {
	messages := make(chan string)
	go func() {
		for {
			time.Sleep(d)
			messages <- msg
		}
	}()
	return messages
}

func firstDemo(log *slog.Logger) {
	log.Debug("begin firstDemo")

	messagesFromA := writeEvery("Huey", 200*time.Millisecond)
	messagesFromB := writeEvery("Dewey", 500*time.Millisecond)
	messagesFromC := writeEvery("Louie", 700*time.Millisecond)

	for {
		select {
		case msgA := <-messagesFromA:
			fmt.Printf("message: %q on channel %s\n", msgA, "A")
		case msgB := <-messagesFromB:
			fmt.Printf("message: %q on channel %s\n", msgB, "B")
		case msgC := <-messagesFromC:
			fmt.Printf("message: %q on channel %s\n", msgC, "C")
		}
	}

	log.Debug(" end  firstDemo")
}

func demo() {
	slogOpts := slog.HandlerOptions{
		//Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slogOpts))

	firstDemo(logger)

	fmt.Println("\n\t\t=-=-=-=\n")

	// nextDemo()
}

func main() {
	fmt.Println("Hello select (channels) world!   🪮 ")

	demo()
}
