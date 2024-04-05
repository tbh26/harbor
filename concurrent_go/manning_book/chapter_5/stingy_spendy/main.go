package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money += 10
		mutex.Unlock()
	}
	fmt.Println("Stingy with mutes done.")
}

func spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 500000; i++ {
		mutex.Lock()
		*money -= 20
		mutex.Unlock()
	}
	fmt.Println("Spendy with mutex done.")
}

func conditionalSpendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 500000; i++ {
		mutex.Lock()
		*money -= 20
		if *money < -1234 {
			fmt.Printf("(account) money is too low (negative), exit.   (%d, i=%d)\n", *money, i)
			os.Exit(1)
		}
		mutex.Unlock()
	}
	fmt.Println("Spendy with mutex done.")
}

func demo() {
	money := 100
	mutex := sync.Mutex{}
	go stingy(&money, &mutex)
	go spendy(&money, &mutex)
	time.Sleep(1 * time.Second)

	mutex.Lock()
	fmt.Println("Money in bank account: ", money)
	mutex.Unlock()

	go stingy(&money, &mutex)
	go conditionalSpendy(&money, &mutex)
	time.Sleep(1 * time.Second)

	mutex.Lock()
	fmt.Println("Money in bank account: ", money)
	mutex.Unlock()
}

func main() {
	fmt.Println("Hello stingy & spendy world 🧐 (chapter 5)")

	demo()
}
