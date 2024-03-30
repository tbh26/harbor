package main

import (
	"fmt"
	"time"
)

func countdown(n *int) {
	for *n > 0 {
		var m int = *n * 42
		time.Sleep(time.Duration(m) * time.Millisecond)
		*n -= 1
	}
}

func countDown() {
	count := 5
	go countdown(&count)
	for count > 0 {
		time.Sleep(40 * time.Millisecond)
		fmt.Println(count)
	}
}

func main() {
	//fmt.Println("Hello countdown world.")
	countDown()
}
