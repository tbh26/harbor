package main

import (
	"fmt"
	"sync"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func printSome(s string, wg *sync.WaitGroup) {
	defer wg.Done()
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

func nextDemo() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, word := range words {
		message := fmt.Sprintf(" - %q   (%d) ", word, i)
		//go printSomething(message)
		go printSome(message, &wg)
	}

	//time.Sleep(10 * time.Millisecond)
	wg.Wait()

	printSomething(" = nextDemo(), done.")
}

func challenge(words []string) {
	var wg sync.WaitGroup
	for i, word := range words {
		wg.Add(1)
		message := fmt.Sprintf("Hello, %s. (%d) ", word, i)
		go printSome(message, &wg)
		wg.Wait()
	}
	printSomething(" = challenge(), done.")
}

func challengeWrapper() {
	words := []string{"universe", "cosmos", "galaxy", "world"}
	challenge(words)
}

func main() {
	fmt.Println()
	fmt.Println("Hello concurrent go intro world!")
	fmt.Println()

	firstDemo()
	fmt.Println()

	nextDemo()
	fmt.Println()

	challengeWrapper()
	fmt.Println()
}
