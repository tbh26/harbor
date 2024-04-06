package main

import (
	"fmt"
	"github.com/tbh26/harbor/concurrent_go/manning_book/chapter_6/barrier"
	"time"
)

func workAndWait(name string, timeToWork int, barrier *barrier.Barrier) {
	start := time.Now()
	for {
		fmt.Printf("%q start working, at %v \n", name, time.Since(start))
		time.Sleep(time.Duration(timeToWork) * time.Second)
		fmt.Printf("%q worker waiting, at %v \n", name, time.Since(start))
		barrier.Wait()
	}
}

func barrierDemo() {
	barrier := barrier.NewBarrier(3)
	go workAndWait("Red", 2, barrier)
	go workAndWait("Green", 5, barrier)
	go workAndWait("Blue", 7, barrier)
	time.Sleep(15 * time.Second)
}

func main() {
	fmt.Println(" =-=  barrier demo =-=")

	barrierDemo()
}
