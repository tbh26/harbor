package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	lastMessage  = "Stop!"
	emptyMessage = ""
	nowFormat    = "15:04:05"
	bufferSize   = 3
)

func slowReceiver(tag string, messages <-chan string, wg *sync.WaitGroup) {
	msg := emptyMessage
	for msg != lastMessage {
		msg = <-messages
		time.Sleep(1 * time.Second)
		fmt.Printf("[%s] Message %q received (at: %s).\n", tag, msg, time.Now().Format(nowFormat))
	}
	wg.Done()
}

func transmitter(tag string, messageChannel chan<- string, wg *sync.WaitGroup) {
	fmt.Printf("Sender/transmitter %q start at %s.\n", tag, time.Now().Format(nowFormat))
	messagesString := fmt.Sprintf("Send or transmit into channels with direction demo. %s", lastMessage)
	for i, message := range strings.Split(messagesString, " ") {
		fmt.Printf("Move message[%d] %q into channel (at %s).\n",
			i, message, time.Now().Format(nowFormat))
		messageChannel <- message
	}
	wg.Done()
}

func demo() {
	msgChannel := make(chan string, bufferSize)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go slowReceiver("(slow) receiver", msgChannel, &waitGroup)
	go transmitter("generate words", msgChannel, &waitGroup)
	waitGroup.Wait()
	fmt.Printf("Done waiting, at %s (on channel with buffer size %d).\n", time.Now().Format(nowFormat), bufferSize)
}

func main() {
	fmt.Println("Hello channel with direction world!   📻 ")
	fmt.Println()

	demo()
}
