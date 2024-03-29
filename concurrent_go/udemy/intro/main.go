package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func firstDemo() {
	// if you run the program with this line uncommented, and the lines 20 commented,
	// everything works as expected
	go printSomething("This may be the first thing to be printed!")

	// but if you comment out line 15 and uncomment the one below this comment,
	// running the program will (probably) just print out the final message,
	// since the program terminates before the goroutine started by this
	// command does not have time to finish.
	go printSomething("This may be the next thing to be printed!")

	// in order to give the goroutine from line 20 time to finish, we could
	// wait for second by uncommenting the line below, but this is hardly
	// a good solution.
	time.Sleep(5 * time.Millisecond)

	printSomething("This should be the last thing to be printed!")
}

func main() {
	fmt.Println()
	fmt.Println("Hello concurrent go intro world!")
	fmt.Println()

	firstDemo()
	fmt.Println()
}
