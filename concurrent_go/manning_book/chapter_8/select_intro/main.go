package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

func writeEvery(msg string, d time.Duration) <-chan string {
	messages := make(chan string)
	go func() {
		for {
			time.Sleep(d)
			messages <- msg
		}
	}()
	return messages
}

func firstDemo(log *slog.Logger) {
	log.Debug("begin firstDemo")
	counter := 1

	messagesFromA := writeEvery("Huey ", 200*time.Millisecond)
	messagesFromB := writeEvery("Dewey", 500*time.Millisecond)
	messagesFromC := writeEvery("Louie", 700*time.Millisecond)

	for {
		select {
		case msgA := <-messagesFromA:
			fmt.Printf("message: %q on channel %s\n", msgA, "A")
		case msgB := <-messagesFromB:
			fmt.Printf("message: %q on channel %s\n", msgB, "B")
		case msgC := <-messagesFromC:
			fmt.Printf("message: %q on channel %s\n", msgC, "C")
		default:
			fmt.Printf(" =-= no messages received =-=  (counter: %d) \n", counter)
			counter += 1
			if counter > 42 {
				return
			}
			time.Sleep(150 * time.Millisecond)
		}
	}

	log.Debug(" end  firstDemo")
}

////

const (
	passwordToGuess = "go far"
	alphabet        = " abcdefghijklmnopqrstuvwxyz"
)

func toBase27(n int) string {
	result := ""
	for n > 0 {
		result = string(alphabet[n%27]) + result
		n /= 27
	}
	return result
}

func guessPassword(from int, upto int, stop chan int, result chan string) {
	for guessN := from; guessN < upto; guessN += 1 {

		select {

		case <-stop:
			fmt.Printf("Stopped at %d [%d, %d]\n", guessN, from, upto-1)
			return

		default:
			if toBase27(guessN) == passwordToGuess {
				result <- toBase27(guessN)
				//fmt.Println("guessN:", guessN)
				close(stop)
				return
			}
		}
	}
	fmt.Printf("Not found between [%d, %d]\n", from, upto-1)
}

func fictitiousPasswordGuessDemo(log *slog.Logger) {
	log.Debug("begin fictitiousPasswordGuessDemo")

	finished := make(chan int)
	passwordFound := make(chan string)
	for i := 1; i <= 387_420_488; i += 10_000_000 {
		go guessPassword(i, i+10_000_000, finished, passwordFound)
	}

	fmt.Printf("fictitiousPasswordGuessDemo, password found: %q \n", <-passwordFound)
	close(passwordFound)
	time.Sleep(1 * time.Second) // why?

	magicN := 108418383
	fmt.Printf("magic N: %d, toBase27(magicN): %q \n", magicN, toBase27(magicN))

	log.Debug(" end  fictitiousPasswordGuessDemo")
}

////

func sendMsgAfter(seconds time.Duration) <-chan string {
	messages := make(chan string)
	go func() {
		time.Sleep(seconds)
		messages <- "Hello channel/select world!"
	}()
	return messages
}

func timeOutDemo(log *slog.Logger, d1 time.Duration, d2 time.Duration) {
	log.Debug("begin timeOutDemo", slog.Any("d1", d1), slog.Any("d2", d2))
	messages := sendMsgAfter(d1)
	timeoutDuration := d2
	fmt.Printf("Waiting for message for %v \n", timeoutDuration)
	select {
	case msg := <-messages:
		fmt.Printf("Message received: %q \n", msg)
	case tNow := <-time.After(timeoutDuration):
		fmt.Printf("Timed out(%v). Waited until %s \n", d2, tNow.Format("15:04:05"))
	}
	log.Debug(" end  timeOutDemo")
}

////

func demo() {
	slogOpts := slog.HandlerOptions{
		//Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slogOpts))

	firstDemo(logger)

	fmt.Println("\n\t\t=-=-=-=\n")

	fictitiousPasswordGuessDemo(logger)

	fmt.Println("\n\t\t=-=-=-=\n")

	timeOutDemo(logger, 1*time.Second, 2*time.Second)
	timeOutDemo(logger, 2*time.Second, 1*time.Second)
}

func main() {
	fmt.Println("Hello select (channels) world!   🪮 ")
	fmt.Println()

	demo()
}
