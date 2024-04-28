package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
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

func downloadPages(quit <-chan bool, urls <-chan string) <-chan string {
	pages := make(chan string)
	go func() {
		defer close(pages)
		moreData, url := true, ""
		for moreData {
			select {
			case url, moreData = <-urls:
				if moreData {
					resp, _ := http.Get(url)
					if resp.StatusCode != 200 {
						panic("Server’s error: " + resp.Status)
					}
					body, _ := io.ReadAll(resp.Body)
					pages <- string(body)
					_ = resp.Body.Close()
				}
			case q := <-quit:
				fmt.Printf("Quitting download pages, received: '%t' \n", q)
				return
			}
		}
	}()
	return pages
}

func thirdDemo() {
	quit := make(chan bool)
	defer close(quit)
	results := downloadPages(quit, generateUrls(quit))
	for result := range results {
		fmt.Printf(" - result length: %d \n", len(result))
	}
	fmt.Println("third demo done.")
}

func extractWords(quit <-chan bool, pages <-chan string) <-chan string {
	words := make(chan string)
	go func() {
		defer close(words)
		wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
		moreData, pg := true, ""
		for moreData {
			select {
			case pg, moreData = <-pages:
				if moreData {
					for _, word := range wordRegex.FindAllString(pg, -1) {
						words <- strings.ToLower(word)
					}
				}
			case q := <-quit:
				fmt.Printf("Quitting extract words, received: '%t' \n", q)
				return
			}
		}
	}()
	return words
}

func fourthDemo() {
	quit := make(chan bool)
	defer close(quit)
	results := extractWords(quit, downloadPages(quit, generateUrls(quit)))
	for result := range results {
		fmt.Printf(" - word: %q", result)
	}
	fmt.Println()
	fmt.Println("fourth demo done.")
}

func demo() {
	firstDemo()
	fmt.Println("\n\t\t=-= use quit channel =-=\n")
	nextDemo()
	fmt.Println("\n\t\t=-= pipeline I =-=\n")
	thirdDemo()
	fmt.Println("\n\t\t=-= pipeline II =-=\n")
	fourthDemo()
	fmt.Println()
}

func main() {
	fmt.Println("Hello channel programming world!  ⛓️  ")
	fmt.Println()

	demo()
}
