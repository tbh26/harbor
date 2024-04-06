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
		fmt.Printf("[%s] Message %q received (at: %s, more: %v).\n", tag, msg, time.Now().Format(nowFormat), ok)
		time.Sleep(300 * time.Millisecond)
		if !ok {
			fmt.Println("(slow rangeReceiver) NO lastMessage, but channel seems closed, break-up.")
			break
		}
	}
	// note: a receiving channel can not be closed
	// close(messages) // close can only be called on a bidirectional or send-only channel
	wg.Done()
}

func rangeReceiver(tag string, messagesChannel <-chan string, wg *sync.WaitGroup) {
	for msg := range messagesChannel {
		fmt.Printf("[%s] received message %q received (at: %s).\n", tag, msg, time.Now().Format(nowFormat))
		time.Sleep(200 * time.Millisecond)
	}
	wg.Done()
}

func resultReceiver(tag string, messagesChannel chan string, wg *sync.WaitGroup) {
	messageCounter := 0
	for msg := range messagesChannel {
		messageCounter += 1
		fmt.Printf("[%s] received message %q received (at: %s, counter: %d).\n",
			tag, msg, time.Now().Format(nowFormat), messageCounter)
		if msg == lastMessage {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	messagesChannel <- fmt.Sprintf("result receiver got %d messages", messageCounter)
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

func transmitterWithResult(tag string, msgSrc string, messageChannel chan string, wg *sync.WaitGroup) {
	fmt.Printf("Transmitter %q start at %s.\n", tag, time.Now().Format(nowFormat))
	for i, message := range strings.Split(msgSrc, " ") {
		fmt.Printf("Move message[%d] %q into channel (at %s).\n", i, message, time.Now().Format(nowFormat))
		messageChannel <- message
	}
	// time.Sleep(200 * time.Millisecond)
	result := <-messageChannel
	fmt.Printf("finishing transmitter, received result: %q (at %v) \n", result, time.Now().Format(nowFormat))
	wg.Done()
}

func firstDemo() {
	msgChannel := make(chan string, bufferSize)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	messagesSource := fmt.Sprintf("Send or transmit into channels with direction demo. %s", lastMessage)
	go slowReceiver("(slow) rangeReceiver", msgChannel, &waitGroup)
	go transmitter("generate first words", messagesSource, msgChannel, &waitGroup)
	waitGroup.Wait()
	fmt.Printf("Done waiting, at %s (on channel with buffer size %d).\n", time.Now().Format(nowFormat), bufferSize)
}

func nextDemo() {
	msgChannel := make(chan string, bufferSize)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	messagesSource := fmt.Sprintf("NO last_message here just nonsense %s", nowFormat)
	go slowReceiver("(slow) rangeReceiver again", msgChannel, &waitGroup)
	go transmitter("next generator", messagesSource, msgChannel, &waitGroup)
	waitGroup.Wait()
	fmt.Printf("Done waiting, at %s (on channel with buffer size %d).\n", time.Now().Format(nowFormat), bufferSize)
}

func rangeReceiverDemo() {
	msgChannel := make(chan string, bufferSize)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	messagesSource := fmt.Sprintf("Again NO last_message just bla %s", "bla")
	go rangeReceiver("range receiver (slowish)", msgChannel, &waitGroup)
	go transmitter("range_rec generator", messagesSource, msgChannel, &waitGroup)
	waitGroup.Wait()
	fmt.Printf("Done waiting, at %s (on channel with buffer size %d).\n", time.Now().Format(nowFormat), bufferSize)
}

func receiveResultFromChannelDemo() {
	//resultChannel := make(chan string, bufferSize)
	resultChannel := make(chan string)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	messagesSource := fmt.Sprintf("transmit receive result on duplex channel %s", lastMessage)
	go resultReceiver("result receiver (also slow)", resultChannel, &waitGroup)
	go transmitterWithResult("transmitter (with result receiver)", messagesSource, resultChannel, &waitGroup)
	waitGroup.Wait()
	fmt.Printf("Done waiting, at %s (on channel with buffer size %d).\n", time.Now().Format(nowFormat), bufferSize)
}

func demo() {
	firstDemo()

	fmt.Println("\n\t=-=-=\n")

	nextDemo()

	fmt.Println("\n\t=-=-=\n")

	rangeReceiverDemo()

	fmt.Println("\n\t=-=-=\n")

	receiveResultFromChannelDemo()
}

func main() {
	fmt.Println("Hello channel with direction world!   📻 ")
	fmt.Println()

	demo()
}
