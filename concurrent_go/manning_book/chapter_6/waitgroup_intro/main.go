package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork(id int, wg *sync.WaitGroup) {
	i := rand.Intn(5)
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Printf("worker %d done. (after %d sec.) \n", id, i)
	wg.Done()
}

func demo() {
	wg := sync.WaitGroup{}
	wgSize := 6
	wg.Add(wgSize)
	for i := 1; i <= wgSize; i++ {
		go doWork(i, &wg)
	}
	wg.Wait()
	fmt.Println("done!")
}

func main() {
	fmt.Println("Hello waitgroup world!    🫷🫸 ")

	demo()
}
