package main

import (
	"fmt"
	"runtime"
	"time"
)

func stingyInitial(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
	}
	fmt.Println("Stingy Done")
}

func spendyInitial(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
	}
	fmt.Println("Spendy Done")
}

func initialDemo() {
	money := 100
	go stingyInitial(&money)
	go spendyInitial(&money)
	time.Sleep(2 * time.Second)
	println("Money in bank account: ", money)
}

func stingy2(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
		runtime.Gosched()
	}
	fmt.Println("Stingy runtime.Gosched() Done")
}

func spendy2(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
		runtime.Gosched()
	}
	fmt.Println("Spendy runtime.Gosched() Done")
}

func demo2() {
	money := 100
	go stingy2(&money)
	go spendy2(&money)
	time.Sleep(2 * time.Second)
	println("Money in bank account (2): ", money)
}

func main() {
	fmt.Println("Hello stingy & spendy world 🧐")

	initialDemo()

	demo2()
}
