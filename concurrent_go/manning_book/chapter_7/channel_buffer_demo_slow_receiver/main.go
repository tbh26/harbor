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

func slowReceiver(tag string, messages chan string, wg *sync.WaitGroup) {
	msg := emptyMessage
	for msg != lastMessage {
		msg = <-messages
		time.Sleep(1 * time.Second)
		fmt.Printf("[%s] Message %q received (at: %s).\n", tag, msg, time.Now().Format(nowFormat))
	}
	wg.Done()
}

func demo() {
	msgChannel := make(chan string, bufferSize)
	messagesString := fmt.Sprintf("Hello golang concurrent demo with channel buffer demo. %s", lastMessage)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	tag := "(slow) receiver on buffered channel"
	go slowReceiver(tag, msgChannel, &waitGroup)
	for i, message := range strings.Split(messagesString, " ") {
		fmt.Printf("Sending message[%d] %q on channel (at %s).\n",
			i, message, time.Now().Format(nowFormat))
		msgChannel <- message
	}
	waitGroup.Wait()
	fmt.Printf("Done waiting, at %s (on channel with buffer size %d).\n", time.Now().Format(nowFormat), bufferSize)
}

func main() {
	fmt.Println("Hello channel with buffer and a slow receiver world!   📡 ")
	fmt.Println()

	demo()
}
