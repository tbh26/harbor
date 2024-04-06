package main

import (
	"fmt"
	"github.com/tbh26/harbor/concurrent_go/manning_book/chapter_6/semaphore_waitgroup"
)

func doWork(id int, wg *semaphore_waitgroup.WaitGrp) {
	fmt.Println(id, "Done working ")
	wg.Done()
}

func demo() {
	wg := semaphore_waitgroup.NewWaitGrp(4)
	for i := 1; i <= 4; i++ {
		go doWork(i, wg)
	}
	wg.Wait()
	fmt.Println("All complete")
}

func main() {
	fmt.Println(" =-=  semaphore waitgroup demo =-=")

	demo()
}
