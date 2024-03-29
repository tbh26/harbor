package main

import (
	"fmt"
	"time"
)

func doWork(id int) {
	layout := time.StampMilli
	fmt.Printf(" - Work %d started  at %s\n", id, time.Now().Format(layout))
	time.Sleep(10 * time.Millisecond)
	fmt.Printf(" - Work %d finished at %s\n", id, time.Now().Format(layout))
}

func firstDemo(upper int) {
	for i := 0; i < upper; i++ {
		doWork(i)
	}
}

func main() {
	fmt.Println()
	fmt.Println("Hello concurrent go intro world!  (II)")
	fmt.Println()

	firstDemo(5)
	fmt.Println()
}
