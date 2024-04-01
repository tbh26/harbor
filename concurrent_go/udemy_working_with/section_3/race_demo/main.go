package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func firstDemo() {
	msg = "Hello, world!"

	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")

	time.Sleep(100 * time.Millisecond)
	fmt.Println(msg)
	time.Sleep(100 * time.Millisecond)
	me := filepath.Base(os.Args[0])
	fmt.Println(fmt.Sprintf("%s; firstDemo done!", me))
}

func nextDemo() {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")
	wg.Wait()

	fmt.Println(msg)
	time.Sleep(100 * time.Millisecond)
	me := filepath.Base(os.Args[0])
	fmt.Println(fmt.Sprintf("%s; nextDemo done!", me))
}

func main() {
	//firstDemo()
	nextDemo()
}
