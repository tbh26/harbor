package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	lastMessage = "Stop!"
	nowFormat   = "15:04:05"
	bufferSize  = 3
)

func slowReceiver(tag string, messages <-chan string, wg *sync.WaitGroup) {
	var msg string
	var ok bool
	for msg != lastMessage {
		msg, ok = <-messages
		time.Sleep(1 * time.Second)
		fmt.Printf("[%s] Message %q received (at: %s, more: %v).\n", tag, msg, time.Now().Format(nowFormat), ok)
		if !ok {
			fmt.Println("(slow receiver) NO lastMessage, but channel seems closed, break-up.")
			break
		}
	}
	// note: a receiving channel can not be closed
	// close(messages) // close can only be called on a bidirectional or send-only channel
	wg.Done()
}

func transmitter(tag string, msgSrc string, messageChannel chan<- string, wg *sync.WaitGroup) {
	fmt.Printf("Sender/transmitter %q start at %s.\n", tag, time.Now().Format(nowFormat))
	for i, message := range strings.Split(msgSrc, " ") {
		fmt.Printf("Move message[%d] %q into channel (at %s).\n",
			i, message, time.Now().Format(nowFormat))
		messageChannel <- message
	}
	close(messageChannel)
	wg.Done()
}

func firstDemo() {
	msgChannel := make(chan string, bufferSize)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	messagesSource := fmt.Sprintf("Send or transmit into channels with direction demo. %s", lastMessage)
	go slowReceiver("(slow) receiver", msgChannel, &waitGroup)
	go transmitter("generate words", messagesSource, msgChannel, &waitGroup)
	waitGroup.Wait()
	fmt.Printf("Done waiting, at %s (on channel with buffer size %d).\n", time.Now().Format(nowFormat), bufferSize)
}

func nextDemo() {
	msgChannel := make(chan string, bufferSize)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	messagesSource := fmt.Sprintf("NO last_message here just nonsense %s", nowFormat)
	go slowReceiver("(slow) receiver", msgChannel, &waitGroup)
	go transmitter("generate words", messagesSource, msgChannel, &waitGroup)
	waitGroup.Wait()
	fmt.Printf("Done waiting, at %s (on channel with buffer size %d).\n", time.Now().Format(nowFormat), bufferSize)
}

func demo() {
	firstDemo()

	fmt.Println("\n\t=-=-=\n")

	nextDemo()
}

func main() {
	fmt.Println("Hello channel with direction world!   📻 ")
	fmt.Println()

	demo()
}
