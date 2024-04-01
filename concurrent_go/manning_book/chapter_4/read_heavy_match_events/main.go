package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func matchRecorder(matchEvents *[]string, mutex *sync.Mutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		*matchEvents = append(*matchEvents,
			"Match event "+strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(15 * time.Millisecond)
		fmt.Println(" = Appended match event")
	}
}

func clientHandler(mEvents *[]string, mutex *sync.Mutex, st time.Time) {
	for i := 0; i < 50; i++ {
		mutex.Lock()
		allEvents := copyAllEvents(mEvents)
		mutex.Unlock()

		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), "events copied in", timeTaken)
	}
}

func copyAllEvents(matchEvents *[]string) []string {
	allEvents := make([]string, 0, len(*matchEvents))
	for _, e := range *matchEvents {
		allEvents = append(allEvents, e)
	}
	return allEvents
}

func initialDemo() {
	mutex := sync.Mutex{}
	var matchEvents = make([]string, 0, 10000)
	for j := 0; j < 10000; j++ {
		matchEvents = append(matchEvents, "Match event")
	}
	go matchRecorder(&matchEvents, &mutex)
	start := time.Now()
	for j := 0; j < 400; j++ {
		go clientHandler(&matchEvents, &mutex, start)
	}
	time.Sleep(6 * time.Second)
}

func matchRecorder2(matchEvents *[]string, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		*matchEvents = append(*matchEvents,
			"Match event "+strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(5 * time.Millisecond)
		fmt.Println(" =-= Appended match event (2)")
	}
}

func clientHandler2(mEvents *[]string, mutex *sync.RWMutex, st time.Time) {
	for i := 0; i < 50; i++ {
		mutex.RLock()
		allEvents := copyAllEvents(mEvents)
		mutex.RUnlock()
		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), "events copied (2) in", timeTaken)
	}
}

func nextDemo() {
	mutex := sync.RWMutex{}
	var matchEvents = make([]string, 0, 10000)
	for j := 0; j < 10000; j++ {
		matchEvents = append(matchEvents, "Match event")
	}
	go matchRecorder2(&matchEvents, &mutex)
	start := time.Now()
	for j := 0; j < 400; j++ {
		go clientHandler2(&matchEvents, &mutex, start)
	}
	time.Sleep(2 * time.Second)
}

func main() {
	fmt.Println("Hello read heavy match events world... ⚽️")

	initialDemo()

	fmt.Println(" =-=-=-=-=-=-=-= ")

	nextDemo()
}
