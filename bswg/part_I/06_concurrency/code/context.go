package code

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

func ContextDemo() {
	fmt.Println("Hello code.ContextDemo() world! (ch6)")
	//
	contextIntro()
	contextWithTimeout()
	contextWithDeadline(3)
	contextWithDeadline(1)
	contextWithValue()
	parentContexts(300)
	parentContexts(700)
	//
	fmt.Println()
}

func incrementer(id int, c *int32, ctx context.Context) {
	t := time.NewTicker(time.Millisecond * 300)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf(" - incrementer done, id: %d, c: %d \n", id, *c)
			return
		case <-t.C:
			atomic.AddInt32(c, 1)
		}
	}
}

func contextIntro() {
	fmt.Println("\n=-= context intro (with cancel) =-=")

	ctx, cancel := context.WithCancel(context.Background())

	var c int32 = 0
	for i := 0; i < 5; i++ {
		go incrementer(i, &c, ctx)
	}

	time.Sleep(time.Second * 1)
	fmt.Println(" - value after second : ", c)
	time.Sleep(time.Millisecond * 500)
	fmt.Println(" - value  in between  : ", c)
	time.Sleep(time.Millisecond * 500)
	fmt.Println(" - value before cancel: ", c)
	cancel() // usually defer-ed
	time.Sleep(time.Second)
	fmt.Println(" =- done =-= value:", c)

}

func delayedIncrementer(i int, info chan<- int) {
	t := time.Duration(i*100) * time.Millisecond
	time.Sleep(t)
	result := i + 1
	info <- result
}

func contextWithTimeout() {
	fmt.Println("\n=-= context with timeout =-=")

	//d := time.Millisecond * 300
	d := time.Millisecond * 277
	ch := make(chan int)
	i := 0
	for {
		ctx, cancel := context.WithTimeout(context.Background(), d)
		go delayedIncrementer(i, ch)
		select {
		case x := <-ch:
			fmt.Printf(" - received %d (i: %d) \n", x, i)
		case <-ctx.Done():
			fmt.Println(" - done!")
		}
		if ctx.Err() != nil {
			fmt.Println(ctx.Err())
			//return
			break // out of for
		}
		cancel() // in a loop, should NOT be defer-ed
		i += 1
	}
	fmt.Println(" =-= done =-=")
	fmt.Println()
}

func accum(c *uint32, ctx context.Context) {
	//t := time.NewTicker(time.Millisecond * 250)
	t := time.NewTicker(time.Millisecond * 400)
	for {
		select {
		case <-t.C:
			atomic.AddUint32(c, 1)
		case <-ctx.Done():
			fmt.Println(" - accum;  ctx done  (stop ticker)")
			t.Stop()
			return
			// break
		}
		if ctx.Err() != nil {
			fmt.Println(ctx.Err())
			//return
			break // out of for
		}
	}
	//t.Stop()
	fmt.Println(" - accum done!   (ticker stopped)")
}

func contextWithDeadline(delay int) {
	fmt.Println("\n=-= context with deadline =-= ", delay)

	d := time.Now().Add(time.Second * time.Duration(delay))
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	var counter uint32 = 0

	for i := 0; i < 5; i++ {
		go accum(&counter, ctx)
	}

	<-ctx.Done()
	// sleep some more...
	time.Sleep(time.Second)
	fmt.Println(" - context with deadline done, counter is:", counter)
	time.Sleep(time.Second)
	fmt.Println()

}

func contextWithValue() {
	fmt.Println("\n=-= context with value =-= ")

	fun := func(ctx context.Context, a int, b int) (int, error) {

		switch ctx.Value("action") {
		case "+":
			return a + b, nil
		case "-":
			return a - b, nil
		default:
			return 0, errors.New("unknown action")
		}
	}

	ctx1 := context.WithValue(context.Background(), "action", "+")
	r, err := fun(ctx1, 40, 2)
	fmt.Printf(" - context fun, result: %d, err: %v \n", r, err)
	ctx2 := context.WithValue(context.Background(), "action", "-")
	r, err = fun(ctx2, 50, 8)
	fmt.Printf(" - context fun, result: %d, err: %v \n", r, err)
	ctx3 := context.WithValue(context.Background(), "action", "~")
	r, err = fun(ctx3, 123, 321)
	fmt.Printf(" - context fun, result: %d, err: %v \n", r, err)

	fmt.Println(" =-= done =-=")
	fmt.Println()
}

func calc(ctx context.Context, d int) {
	switch ctx.Value("action") {
	case "quick":
		fmt.Println(" - quick answer", d)
	case "slow":
		time.Sleep(time.Millisecond * 500)
		fmt.Println(" - slow answer", d)
	case <-ctx.Done():
		fmt.Println(" - calc done")
	default:
		panic("unknown action")
	}
}

func parentContexts(delay int) {
	fmt.Println("\n=-= parent contexts =-= delay:", delay)

	t := time.Millisecond * time.Duration(delay)
	ctx, cancel := context.WithTimeout(context.Background(), t)
	qCtx := context.WithValue(ctx, "action", "quick")
	defer cancel()

	go calc(qCtx, delay)
	<-qCtx.Done()

	ctx2, cancel2 := context.WithTimeout(context.Background(), t)
	sCtx := context.WithValue(ctx2, "action", "slow")
	defer cancel2()

	go calc(sCtx, delay)
	<-sCtx.Done()

	fmt.Println(" =-= parent contexts done =-=")

}
