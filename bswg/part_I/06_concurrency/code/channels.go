package code

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func ChannelsDemo() {
	fmt.Println("Hello code.ChannelsDemo() world! (ch6)")
	//
	firstChannel()
	nextChannel()
	bufferedChannel()
	channelClose()
	channelRange()
	channelDirections()
	//
	fmt.Println()
}

func firstGenerator(ch chan int) {
	sum := 0
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 400)
		sum = sum + i
		fmt.Printf(" -- sum: %2d, i: %d  (firstGenerator) \n", sum, i)
	}
	ch <- sum
}

func firstChannel() {
	fmt.Println("\n=-= first channel =-=")
	//
	ch := make(chan int)
	go firstGenerator(ch)
	fmt.Println(" - waiting for firstGenerator..")
	r := <-ch
	fmt.Printf(" - done, result: %d \n", r)
	//
	fmt.Println()
}

func nextGenerator(ch chan int) {
	fmt.Println(" -- nextGenerator waits for (int) input")
	n := <-ch
	fmt.Println(" -- nextGenerator received: ", n)
	sum := 0
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 200)
		sum = sum + i
		fmt.Printf(" -- sum: %2d, i: %d  (nextGenerator) \n", sum, i)
	}
	ch <- sum
}

func nextChannel() {
	fmt.Println("\n=-= next channel =-=")
	//
	ch := make(chan int)
	go nextGenerator(ch)
	fmt.Println(" - small break")
	time.Sleep(time.Second)
	ch <- 6
	fmt.Println(" - waiting for nextGenerator")
	r := <-ch
	fmt.Printf(" - done, result: %d \n", r)
	//
	fmt.Println()
}

func genAbc(ch chan string) {
	time.Sleep(time.Second)
	ch <- "Abc"
}

func genXyz(ch chan string) {
	time.Sleep(time.Millisecond * 500)
	ch <- "Xyz"
}

func bufferedChannel() {
	fmt.Println("\n=-= buffered channel =-=")
	//
	//ch := make(chan string) // <- fatal error: all goroutines are asleep - deadlock!
	ch := make(chan string, 1) //
	ch <- "from main"          // first message
	go genXyz(ch)
	go genAbc(ch)
	go genXyz(ch)
	fmt.Println(" -- waiting for channel messages")
	l := len(ch)
	fmt.Printf(" - received from channel: %s (len: %d, cap: %d) \n", <-ch, l, cap(ch))
	fmt.Printf(" - received from channel: %s (len: %d, cap: %d) \n", <-ch, len(ch), cap(ch))
	fmt.Printf(" - received from channel: %s (len: %d, cap: %d) \n", <-ch, len(ch), cap(ch))
	fmt.Printf(" - received from channel: %s (len: %d, cap: %d) \n", <-ch, len(ch), cap(ch))
	fmt.Println(" -- done --")
	fmt.Println()
}

func genWords(ch chan string, sentence string) {
	for _, w := range strings.Split(sentence, " ") {
		time.Sleep(time.Millisecond * 200)
		// fmt.Println(" word: ", w)
		ch <- w
	}
	close(ch)
	fmt.Println("  =-= genWords, done =-=")
}

func channelClose() {
	fmt.Println("\n=-= channel & close =-=")
	//
	sentence := "The quick brown fox jumps over the lazy dog"
	ch := make(chan string)
	go genWords(ch, sentence)
	for {
		w, hasValue := <-ch
		if hasValue {
			fmt.Printf(" - received from channel, word: %s \n", w)
		} else {
			fmt.Println(" - done (channel closed)")
			break
		}
	}
}

func channelRange() {
	fmt.Println("\n=-= channel & range =-=")
	//
	sentence := "Love Will Tear Us Apart"
	ch := make(chan string)
	go genWords(ch, sentence)
	for word := range ch {
		fmt.Printf(" - received from channel, word: %s \n", word)
	}
	fmt.Println(" - done (channel closed)")
}

func wordsSender(out chan<- string, sentence string) {
	for _, w := range strings.Split(sentence, " ") {
		time.Sleep(time.Millisecond * 200)
		out <- w
	}
	close(out)
	log.Printf(" =-= words sender, done. =-=")
}

func receiver(input <-chan string) {
	for i := range input {
		log.Printf(" -- receiver, got: %s \n", i)
	}
	log.Printf(" =-= receiver, done. =-=")
}

func channelDirections() {
	fmt.Println("\n=-= channel & direction =-=")
	//
	sentence := "Roses are red violets are blue"
	ch := make(chan string)
	go wordsSender(ch, sentence)
	go receiver(ch)
	time.Sleep(time.Second * 2)
	fmt.Println(" - done (channel closed)")
}
