package code

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func WaitGroupDemo() {
	fmt.Println("Hello code.WaitGroupDemo() world! (ch6)")
	//
	wgDemo()
	//
	fmt.Println()
}

func wordGenerator(out chan<- string, wg *sync.WaitGroup, sentence string) {
	//
	defer wg.Done()
	//
	for _, w := range strings.Split(sentence, " ") {
		time.Sleep(time.Millisecond * 321)
		fmt.Printf(" -=- word: %s \n", w)
		out <- w
	}
	time.Sleep(time.Millisecond * 123)
	close(out)
	fmt.Println(" =-= word generator, done. =-=")
}

func wordConsumer(id int, sleep time.Duration, input <-chan string, wg *sync.WaitGroup) {
	//
	defer wg.Done()
	//
	for w := range input {
		time.Sleep(time.Millisecond * sleep)
		fmt.Printf(" - word consumer %d, delay: %d, word: %s \n", id, int(sleep), w)
	}
	time.Sleep(time.Millisecond * 123)
	fmt.Printf(" =-= word consumer, done (%d) =-= \n", id)
}

func wgDemo() {
	fmt.Println("\n=-= demo =-=")
	//
	sentence := "Humpty Dumpty sat on a wall Humpty Dumpty had a great fall"
	channel := make(chan string, 22)
	var wg sync.WaitGroup
	wg.Add(3)
	//
	go wordGenerator(channel, &wg, sentence)
	go wordConsumer(1, 700, channel, &wg)
	go wordConsumer(7, 200, channel, &wg)
	//
	fmt.Println(" === waiting === ")
	wg.Wait()
	//
	fmt.Println(" === done === ")
}
