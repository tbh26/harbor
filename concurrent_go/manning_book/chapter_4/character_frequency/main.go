package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int, m *sync.Mutex) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	m.Lock()
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		} // else skip
	}
	m.Unlock()
	fmt.Printf("Completed url: %q (at; %v)\n", url, time.Now().Format("15:04:05"))
}

func simultaneousWithMutex() {
	var mutex sync.Mutex
	var frequency = make([]int, len(allLetters))
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex)
	}
	time.Sleep(3 * time.Second)
	mutex.Lock()
	for i, c := range allLetters {
		fmt.Printf(" %c - %d \n", c, frequency[i])
	}
	mutex.Unlock()
}

func simultaneousWithTryLock() {
	mutex := sync.Mutex{}
	var frequency = make([]int, 26)
	for i := 2000; i <= 2200; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex)
	}
	for i := 0; i < 30; i++ {
		time.Sleep(100 * time.Millisecond)
		if mutex.TryLock() {
			for i, c := range allLetters {
				fmt.Printf(" %c - %d \n", c, frequency[i])
			}
			mutex.Unlock()
			fmt.Println()
		} else {
			fmt.Println("Mutex in use. (try-lock failed)")
		}
	}
}

func main() {
	fmt.Println(" =-=  count letter/char frequency =-=")

	simultaneousWithMutex()

	fmt.Println(" =-=-=-=-=-=-=-= ")

	simultaneousWithTryLock()
}
