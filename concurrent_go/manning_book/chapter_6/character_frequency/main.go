package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const nowFormat = "15:04:05"
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
	fmt.Printf("Completed url: %q (at; %v)\n", url, time.Now().Format(nowFormat))
}

func simultaneousWithWaitGroup() {
	wg := sync.WaitGroup{}
	wgSize := 31 // so we do 32!  ( 00..31 )
	var frequency = make([]int, len(allLetters))
	freqMutex := sync.Mutex{}
	wg.Add(wgSize + 1)
	fmt.Printf("start simultaneous count (waitgroup version) at; %v \n", time.Now().Format(nowFormat))
	for i := 1000; i <= 1000+wgSize; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		go func() {
			countLetters(url, frequency, &freqMutex)
			fmt.Printf(" - i=%d done\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
	freqMutex.Lock()
	for i, c := range allLetters {
		fmt.Printf(" %c - %d \n", c, frequency[i])
	}
	freqMutex.Unlock()
	fmt.Println()
}

func demo() {
	simultaneousWithWaitGroup()
}

func main() {
	fmt.Println(" =-=  count letter/char frequency waitgroup version =-=")

	demo()
}
