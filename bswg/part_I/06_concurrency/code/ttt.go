package code

import (
	"fmt"
	"time"
)

func TttDemo() {
	fmt.Println("Hello code.TttDemo() world! (ch6)")
	//
	tttDemo()
	tttDemo2(4)
	tttDemo2(6)
	//
	fmt.Println()
}

func incrementWorker(x *int) {
	for {
		time.Sleep(time.Millisecond * 400)
		*x += 1
	}
}

func tttDemo() {
	fmt.Println("\n=-= Timers, tickers and timeouts. =-=")
	//
	timer := time.NewTimer(time.Second * 4)
	//timer := time.NewTimer(time.Second * 7)
	ticker := time.NewTicker(time.Second)
	continueLoop := true
	//
	x := 0
	go incrementWorker(&x)
	//
	for {
		select {
		case t := <-timer.C:
			fmt.Printf(" -  timer -> %3d   (%v)\n", x, t)
			continueLoop = false
		case t := <-ticker.C:
			fmt.Printf(" - ticker -> %3d   (%v)\n", x, t)
		}
		if !continueLoop {
			break
		}
		//if x >= 10 {
		//	fmt.Println(" -- break -- ", x)
		//	break
		//}
	}
	fmt.Println(" =-= done =-=")
	fmt.Println()
}

func reaction(t *time.Ticker) {
	for {
		select {
		case x := <-t.C:
			fmt.Println(" - quick", x)
		}
	}
}

func slowReaction(t *time.Timer) {
	select {
	case x := <-t.C:
		fmt.Println(" - slow", x)
	}
}

func tttDemo2(delay time.Duration) {
	fmt.Println("\n=-= Timers, tickers and timeouts II =-=")
	//
	quick := time.NewTicker(time.Second)
	slow := time.NewTimer(time.Second * 5)
	stopper := time.NewTimer(time.Second * delay)
	//
	go reaction(quick)
	go slowReaction(slow)
	//
	<-stopper.C
	quick.Stop()
	//
	stopped := slow.Stop()
	fmt.Printf(" Stopped before the event: %t  (delay: %d) \n", stopped, int(delay))
	//
	fmt.Println(" =-= done =-= ")
}
