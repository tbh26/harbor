package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func stingy(money *int, cond *sync.Cond) {
	for i := 0; i < 1000000; i++ {
		cond.L.Lock()
		*money += 10
		cond.Signal()
		cond.L.Unlock()
	}
	fmt.Println("Stingy with conditional mutes done.")
}

func spendy(money *int, cond *sync.Cond) {
	for i := 0; i < 500000; i++ {
		cond.L.Lock()
		for *money < 20 {
			fmt.Printf("Spendy waiting for enough balance. (%d, i=%d)\n", *money, i)
			cond.Wait()
		}
		*money -= 20
		if *money < 0 {
			fmt.Println("Money is negative, exit!")
			os.Exit(1)
		}

		cond.L.Unlock()
	}
	fmt.Println("Spendy with conditional mutex done.")
}

func demo() {
	money := 100
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	go stingy(&money, cond)
	go spendy(&money, cond)
	time.Sleep(1 * time.Second)

	mutex.Lock()
	fmt.Println("Money in bank account: ", money)
	mutex.Unlock()
}

func main() {
	fmt.Println("Hello stingy & spendy world II 🧐 (chapter 5)")

	demo()
}
