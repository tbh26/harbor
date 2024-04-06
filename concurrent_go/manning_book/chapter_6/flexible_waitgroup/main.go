package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func fileSearch(dir string, filename string, wg *sync.WaitGroup) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			fmt.Println(fpath)
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(fpath, filename, wg)
		}
	}
	wg.Done()
}

func fileSearchDemo() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	//go fileSearch(os.Args[1], os.Args[2], &wg)
	go fileSearch("../", "main.go", &wg)
	wg.Wait()
}

////

type WaitGrp struct {
	groupSize int
	cond      *sync.Cond
}

func NewWaitGrp() *WaitGrp {
	return &WaitGrp{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (wg *WaitGrp) Add(delta int) {
	wg.cond.L.Lock()
	wg.groupSize += delta
	wg.cond.L.Unlock()
}

func (wg *WaitGrp) Wait() {
	wg.cond.L.Lock()
	for wg.groupSize > 0 {
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}

func (wg *WaitGrp) Done() {
	wg.cond.L.Lock()
	wg.groupSize--
	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}
	wg.cond.L.Unlock()
}

////

func doWork(id int, tag string, wg *WaitGrp) {
	fmt.Printf("wg %d, done working (tag: %q).\n", id, tag)
	wg.Done()
}

func improvedWaitGroupDemo() {
	wg := NewWaitGrp()
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go doWork(i, "first", wg)
		go doWork(i, "next", wg)
	}
	wg.Wait()
	fmt.Println("All complete")
}

func demo() {
	fileSearchDemo()

	fmt.Println("\n =-=-=-= \n")

	improvedWaitGroupDemo()
}

func main() {
	fmt.Println(" =-=  more flexible waitgroup demo(s) =-=")

	demo()
}
