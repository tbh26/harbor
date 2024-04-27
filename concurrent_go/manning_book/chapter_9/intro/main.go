package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randSrc = rand.NewSource(time.Now().UnixNano())
var random = rand.New(randSrc)

func generateRandomWord(strLen int) string {
	chars := "abcdefghijklmnopqrstuvwxyz"
	word := make([]byte, strLen)
	for i := range word {
		word[i] = chars[random.Intn(len(chars))]
	}
	return string(word)
}

func wordsReceiver(words <-chan string, quit chan bool, max int) {
	go func() {
		var word string
		fmt.Printf("wordsReceiver, words: '%v', quit: '%v', max: %d \n", words, quit, max)
		for i := 0; i < max; i++ {
			word = <-words
			fmt.Printf(" - received: %q  (%d)\n", word, len(word))
			time.Sleep(123 * time.Millisecond)
		}
		close(quit)
	}()
}

func firstDemo() {
	words := make(chan string)
	quit := make(chan bool)
	wordsReceiver(words, quit, 5)
	var next string
	nextLen := 2
	for {
		next = generateRandomWord(nextLen)
		nextLen += 1
		select {
		case words <- next:
		case q := <-quit:
			fmt.Printf("Quitting words gen. (%v) \n", q)
			return
		}
	}
}

func generateUrls(quit <-chan bool) <-chan string {
	urls := make(chan string)
	go func() {
		defer close(urls)
		for i := 1000; i <= 1030; i++ {
			url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
			time.Sleep(12 * time.Millisecond)
			select {
			case urls <- url:
			case q := <-quit:
				fmt.Printf("Quitting urls gen, received: '%t' \n", q)
				return
			}
		}
	}()
	return urls
}

func nextDemo() {
	quit := make(chan bool)
	defer close(quit)
	results := generateUrls(quit)
	for result := range results {
		fmt.Printf("Result: %q \n", result)
		if result == "https://www.rfc-editor.org/rfc/rfc1025.txt" {
			quit <- true
		}
	}
	fmt.Println("next demo done.")
}

func demo() {
	firstDemo()

	fmt.Println("\n\t\t=-=-=-=\n")

	nextDemo()
	fmt.Println()
}

func main() {
	fmt.Println("Hello channel programming world!  ⛓️  ")
	fmt.Println()

	demo()
}
