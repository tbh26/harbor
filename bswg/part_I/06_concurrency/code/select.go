package code

import (
	"fmt"
	"strings"
	"time"
)

func SelectDemo() {
	fmt.Println("Hello code.SelectDemo() world! (ch6)")
	//
	selectIntro()
	selectWithFlags()
	selectWithDefault(0)
	selectWithDefault(1)
	//
	fmt.Println()
}

func wordSender(out chan<- string, sentence string) {
	for _, w := range strings.Split(sentence, " ") {
		time.Sleep(time.Millisecond * 150)
		out <- w
	}
	time.Sleep(time.Millisecond * 20)
	close(out)
	fmt.Println(" =-= word sender, done. =-=")
}

func numberSender(out chan<- int, numbers []int) {
	for _, n := range numbers {
		time.Sleep(time.Millisecond * 150)
		out <- n
	}
	time.Sleep(time.Millisecond * 20)
	close(out)
	fmt.Println(" =-= number sender, done. =-=")
}

func selectIntro() {
	fmt.Println("\n=-= channel, select intro =-=")
	//
	sentence := "Roses are red violets are blue"
	numbers := []int{1, 2, 3, 42, 55, 6}

	wordChannel := make(chan string)
	numberChannel := make(chan int)

	go wordSender(wordChannel, sentence)
	go numberSender(numberChannel, numbers)

	for i := 1; i < 15; i += 1 {
		select {
		case w, hasValue := <-wordChannel:
			fmt.Printf(" -- received[%d] word: %s  (%t)\n", i, w, hasValue)
		case n, hasValue := <-numberChannel:
			fmt.Printf(" -- received[%d] number: %d  (%t)\n", i, n, hasValue)
		}
	}
	time.Sleep(time.Second)
	fmt.Println(" - done (channel closed)")
}

func selectWithFlags() {
	fmt.Println("\n=-= channel, select with flags =-=")
	//
	sentence := "The quick brown fox jumps over the lazy dog"
	numbers := []int{1, 22, 3333, 42}

	wordChannel := make(chan string)
	numberChannel := make(chan int)

	go wordSender(wordChannel, sentence)
	go numberSender(numberChannel, numbers)

	wordsAvail, numbersAvail := true, true

	for wordsAvail || numbersAvail {
		select {
		case w, ok := <-wordChannel:
			if ok {
				fmt.Printf(" -- received word: %s  (%t)\n", w, ok)
			} else {
				wordsAvail = false
			}
		case n, ok := <-numberChannel:
			if ok {
				fmt.Printf(" -- received number: %d  (%t)\n", n, ok)
			} else {
				numbersAvail = false
			}
		}
	}
	time.Sleep(time.Second)
	fmt.Println(" - done (channel closed)")
}

func selectWithDefault(bs int) {
	fmt.Println("\n=-= channel, select with default =-= ", bs)
	//
	//ch := make(chan int)
	//ch := make(chan int, 1)
	ch := make(chan int, bs)

	select {
	case i, ok := <-ch:
		fmt.Printf("Received %d, %t \n", i, ok)
	default:
		fmt.Println("Nothing received")
	}

	select {
	case ch <- 42:
		fmt.Println("Send 42")
	default:
		fmt.Println("Nothing sent")
	}

	select {
	case i, ok := <-ch:
		fmt.Printf("Received %d, %t \n", i, ok)
	default:
		fmt.Println("Nothing received")
	}
	//
	fmt.Println()
}
