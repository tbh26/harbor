package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {
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
	fmt.Println("Completed:", url)
}

func inSequence() {
	var frequency = make([]int, len(allLetters))
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		countLetters(url, frequency)
	}
	for i, c := range allLetters {
		fmt.Printf(" %c - %d \n", c, frequency[i])
	}
}

func simultaneous() {
	var frequency = make([]int, len(allLetters))
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency)
	}
	time.Sleep(5 * time.Second)
	for i, c := range allLetters {
		fmt.Printf(" %c - %d \n", c, frequency[i])
	}
}

func main() {
	inSequence()

	fmt.Println(" =-=-=-=-=-=-=-= ")

	simultaneous()
}
