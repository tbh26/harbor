package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func doWorkUnsafe(cond *sync.Cond) {
	fmt.Println("Work started (unsafe)")
	fmt.Println("Work finished (unsafe)")
	cond.Signal() // should have lock/unlock around it
}

func doWork(cond *sync.Cond) {
	fmt.Println("Work started")
	fmt.Println("Work finished")
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()
}

func waitDemo() {
	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()
	for i := 0; i < 50000; i++ {
		//go doWorkUnsafe(cond)
		go doWork(cond)
		fmt.Println("Waiting for child goroutine")
		runtime.Gosched() // increase the chance of "all goroutines are asleep - deadlock!" error, if using doWorkUnsafe
		cond.Wait()
		fmt.Println("Child goroutine finished")
	}
	cond.L.Unlock()
}

////

func playerHandler(cond *sync.Cond, playersRemaining *int, playerId int) {
	cond.L.Lock()
	fmt.Printf("player %d, connected.\n", playerId)
	*playersRemaining--
	if *playersRemaining == 0 {
		cond.Broadcast()
	}
	for *playersRemaining > 0 {
		fmt.Printf("player %d; waiting for other players.\n", playerId)
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Printf("All connected!   (player %d ready) \n", playerId)
	// Game start...
}

func gameSyncDemo() {
	cond := sync.NewCond(&sync.Mutex{})
	playersInGame := 4
	for playerId := 0; playerId < 4; playerId++ {
		go playerHandler(cond, &playersInGame, playerId)
		time.Sleep(1 * time.Second)
	}
}

////

func demo() {
	waitDemo()

	fmt.Println("\n =-=-= \n")

	gameSyncDemo()
}

func main() {
	fmt.Println("Hello wait, signal & broadcast world.")

	demo()
}
