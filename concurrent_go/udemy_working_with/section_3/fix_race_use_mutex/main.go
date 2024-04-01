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

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func firstDemo() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	go updateMessage("Hello, universe!", &mutex)
	go updateMessage("Hello, cosmos!", &mutex)

	time.Sleep(100 * time.Millisecond)
	fmt.Println(msg)
	time.Sleep(100 * time.Millisecond)
	me := filepath.Base(os.Args[0])
	fmt.Println(fmt.Sprintf("%s; firstDemo done!", me))
}

func nextDemo() {
	msg = "Hello, world!"

	var nextMutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &nextMutex)
	go updateMessage("Hello, cosmos!", &nextMutex)
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
