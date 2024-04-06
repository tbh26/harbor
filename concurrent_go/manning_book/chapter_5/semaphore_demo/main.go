package main

import (
	"fmt"
	"github.com/tbh26/harbor/concurrent_go/manning_book/chapter_5/semaphore"
)

func doWork(semaphore *semaphore.Semaphore) {
	fmt.Println("Work started")
	fmt.Println("Work finished")
	semaphore.Release()
}

func demo() {
	semaphore := semaphore.NewSemaphore(0)
	for i := 0; i < 50000; i++ {
		go doWork(semaphore)
		fmt.Println("Waiting for child goroutine")
		semaphore.Acquire()
		fmt.Println("Child goroutine finished")
	}
}

func main() {
	fmt.Println("Hello semaphore demo world..   🧙 ")

	demo()
}
