package code

import (
	"fmt"
	"time"
)

func IntroDemo() {
	fmt.Println("Hello code.IntroDemo() world! (ch6)")
	//
	intro()
	delayIntro()
	//
	fmt.Println("...")
	time.Sleep(time.Millisecond * 1234)
	fmt.Println()
}

func Intro() {
	for {
		time.Sleep(time.Millisecond * 200)
		fmt.Println(" -- intro demo! (go-routine)")
	}
}

func intro() {

	fmt.Println("\n=-= intro demo (go-routine is sortof lightweight thread) =-=")

	go Intro()

	word := "intro-demo"
	for i, ch := range word {
		fmt.Printf(" - intro [%d], char: %c  (main / %s)\n", i, ch, word)
		time.Sleep(time.Millisecond * 50)
	}

	fmt.Println("=-= done (also main)=-=")
}

func DelayDemo(delay int) {
	time.Sleep(time.Millisecond * time.Duration(20*delay))
	fmt.Printf(" -- delay demo; %d   (go-routine) \n", delay)
}

func delayIntro() {

	fmt.Println("\n=-= intro demo II, delay =-=")

	word := "delay-demo"
	for i, ch := range word {
		fmt.Printf(" - delay [%d], char: %c  (main / %s)\n", i, ch, word)
		go DelayDemo(i)
		time.Sleep(time.Millisecond * 50)
	}

	fmt.Println("=-= done (delay)=-=")
}
