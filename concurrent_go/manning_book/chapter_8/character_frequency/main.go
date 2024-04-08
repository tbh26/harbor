package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string) <-chan []int {
	result := make(chan []int)
	go func() {
		defer close(result)
		frequency := make([]int, len(allLetters))
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			panic("Server returning error status code: " + resp.Status)
		}
		body, _ := io.ReadAll(resp.Body)
		for _, b := range body {
			c := strings.ToLower(string(b))
			cIndex := strings.Index(allLetters, c)
			if cIndex >= 0 {
				frequency[cIndex] += 1
			} // else skip
		}
		fmt.Printf("Completed url: %q (at; %v)\n", url, time.Now().Format("15:04:05"))
		result <- frequency
	}()
	return result
}

func charFrequencyDemo() {
	results := make([]<-chan []int, 0)
	totalFrequencies := make([]int, len(allLetters))
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		results = append(results, countLetters(url))
	}
	for _, c := range results {
		frequencyResult := <-c
		for i := 0; i < 26; i++ {
			totalFrequencies[i] += frequencyResult[i]
		}
	}
	for i, c := range allLetters {
		fmt.Printf(" - %c - %d \n", c, totalFrequencies[i])
	}
}

func charFrequencyDemo2() {
	results := make([]<-chan []int, 0)
	totalFrequencies := make([]int, 26)
	for i := 2000; i <= 2200; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		results = append(results, countLetters(url))
	}
	for _, c := range results {
		frequencyResult := <-c
		for i := 0; i < 26; i++ {
			totalFrequencies[i] += frequencyResult[i]
		}
	}
	for i, c := range allLetters {
		fmt.Printf(" - %c - %d \n", c, totalFrequencies[i])
	}
}

func demo() {
	charFrequencyDemo()

	fmt.Println("\n\t\t=-=-=-=\n")

	charFrequencyDemo2()
}

func main() {
	fmt.Println(" =-=  count letter/char frequency using channels =-=    🧮 ")

	demo()
}
