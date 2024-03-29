package main

import (
	"fmt"
	"runtime"
	"time"
)

func doWork(id int) {
	layout := time.StampMilli
	fmt.Printf(" - Work %d started  at %s\n", id, time.Now().Format(layout))
	time.Sleep(12 * time.Millisecond)
	fmt.Printf(" - Work %d finished at %s\n", id, time.Now().Format(layout))
}

func firstDemo(upper int) {
	for i := 0; i < upper; i++ {
		doWork(i)
	}
	fmt.Println("sequential work done")
}

func nextDemo(upper int) {
	for i := 0; i < upper; i++ {
		go doWork(i)
	}
	time.Sleep(33 * time.Millisecond)
	fmt.Println("next done (parallel)")
}

func runtimeInfo() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}

func main() {
	fmt.Println()
	fmt.Println("Hello concurrent go intro world!  (II)")
	fmt.Println()

	upper := 6
	firstDemo(upper)
	fmt.Println()

	nextDemo(upper)
	fmt.Println()

	runtimeInfo()
	fmt.Println()
}
