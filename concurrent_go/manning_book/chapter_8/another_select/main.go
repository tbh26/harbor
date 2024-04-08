package main

import (
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"os"
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

func demo() {
	slogOpts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slogOpts))

	anotherSelectDemo(logger)
}

func main() {
	fmt.Println("Hello another select channel world!  🪃 ")
	fmt.Println()

	demo()
}
