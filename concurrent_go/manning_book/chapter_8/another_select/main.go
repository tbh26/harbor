package main

import (
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"os"
	"time"
)

func primesOnly(inputs <-chan int) <-chan int {
	results := make(chan int)
	go func() {
		for c := range inputs {
			isPrime := c != 1
			for i := 2; i <= int(math.Sqrt(float64(c))); i++ {
				if c%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				results <- c
			}
		}
	}()
	return results
}

func anotherSelectDemo(log *slog.Logger) {
	log.Debug("begin anotherSelectDemo")

	numbersChannel := make(chan int)
	primes := primesOnly(numbersChannel)
	for i := 0; i < 100; {
		select {
		case numbersChannel <- rand.Intn(1000000000) + 1:
		case p := <-primes:
			fmt.Println("Found prime:", p)
			i++
		}
	}

	log.Debug(" end  anotherSelectDemo")
}

func generateAmounts(n int) <-chan int {
	amounts := make(chan int)
	go func() {
		defer close(amounts)
		for i := 0; i < n; i++ {
			amounts <- rand.Intn(100) + 1
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return amounts
}

func nilChannelUsage(log *slog.Logger) {
	log.Debug("begin nilChannelUsage")
	sales := generateAmounts(50)
	expenses := generateAmounts(40)
	endOfDayAmount := 0
	//
	// This pattern of merging channel data into one result/stream is referred to as a fan-in pattern.
	for sales != nil || expenses != nil {
		select {
		case sale, moreData := <-sales:
			if moreData {
				fmt.Println("Sale of:", sale)
				endOfDayAmount += sale
			} else {
				sales = nil
			}
		case expense, moreData := <-expenses:
			if moreData {
				fmt.Println("Expense of:", expense)
				endOfDayAmount -= expense
			} else {
				expenses = nil
			}
		}
	}
	fmt.Println("End of day profit and loss:", endOfDayAmount)
	log.Debug(" end  nilChannelUsage")
}

func demo() {
	slogOpts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slogOpts))

	anotherSelectDemo(logger)

	fmt.Println("\n\t\t=-=-=-=\n")

	nilChannelUsage(logger)
}

func main() {
	fmt.Println("Hello another select channel world!  🪃 ")
	fmt.Println()

	demo()
}
