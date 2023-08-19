package code

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func OtherDemo() {
	fmt.Println("Hello code.OtherDemo() world! (ch6)")
	//
	onceDemo()
	mutexDemo()
	atomicsDemo()
	atomicsDemoII()
	//
	fmt.Println()
}

////

var firstOnce int

func onceSetter(i int, ch chan bool, once *sync.Once) {
	t := rand.Uint32() % 300
	time.Sleep(time.Duration(t) * time.Millisecond)
	once.Do(func() {
		firstOnce = i
	})
	ch <- true
	fmt.Printf(" - onceSetter done, value: %d, t / delay: %d \n", i, t)
}

func onceDemo() {
	fmt.Println("\n=-= other; once demo =-=")
	rand.Seed(time.Now().UnixNano())

	var once sync.Once

	ch := make(chan bool)
	for i := 0; i < 5; i++ {
		go onceSetter(i, ch, &once)
	}

	for i := 0; i < 5; i++ {
		<-ch
	}
	fmt.Println(" - the firstOnce is/was", firstOnce)
	fmt.Println()
}

////

func mutexWriter(x map[int]int, factor int, m *sync.Mutex) {
	i := 1
	for {
		time.Sleep(time.Second)
		m.Lock()
		x[i] = x[i-1] * factor
		m.Unlock()
		fmt.Printf(" - mutex writer, factor: %d, i: %d \n", factor, i)
		i++
	}
}

func mutexReader(x map[int]int, m *sync.Mutex) {
	for {
		time.Sleep(time.Millisecond * 500)
		m.Lock()
		fmt.Println(" - mutex reader", x)
		m.Unlock()
	}
}

func mutexDemo() {
	fmt.Println("\n=-= other; mutex demo =-=")

	x := make(map[int]int)
	x[0] = 1

	m := sync.Mutex{}
	go mutexWriter(x, 2, &m)
	go mutexReader(x, &m)

	time.Sleep(time.Millisecond * 300)
	go mutexWriter(x, 3, &m)

	time.Sleep(time.Second * 4)
	fmt.Printf(" =-= done =-= %v \n", x)
	fmt.Println()
}

////

func atomicIncreaser(counter *int32) {
	for {
		atomic.AddInt32(counter, 2)
		time.Sleep(time.Millisecond * 500)
	}
}

func atomicDecreaser(counter *int32) {
	for {
		atomic.AddInt32(counter, -1)
		time.Sleep(time.Millisecond * 300)
	}
}

func atomicsDemo() {
	fmt.Println("\n=-= other; atomic(s) demo =-=")

	var counter int32 = 0

	go atomicIncreaser(&counter)
	go atomicDecreaser(&counter)

	for i := 0; i < 8; i++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Printf(" -- atomics demo, value: %d  (i: %d) \n", atomic.LoadInt32(&counter), i)
	}
	fmt.Printf(" -- done; atomics demo, value: %d  \n", atomic.LoadInt32(&counter))
	fmt.Println()

}

////

type Monitor struct {
	ActiveUsers int
	Requests    int
}

func updater(monitor atomic.Value, m *sync.Mutex) {
	for {
		time.Sleep(time.Millisecond * 50)
		m.Lock()
		current := monitor.Load().(*Monitor)
		current.ActiveUsers += 100
		current.Requests += 300
		monitor.Store(current)
		m.Unlock()
	}
}

func observe(monitor atomic.Value) {
	for {
		time.Sleep(time.Millisecond * 100)
		current := monitor.Load()
		fmt.Printf(" ---- Monitor %v \n", current)
	}
}

func atomicsDemoII() {
	fmt.Println("\n=-= other; next atomic(s) demo =-=")

	var monitor atomic.Value
	monitor.Store(&Monitor{0, 0})
	m := sync.Mutex{}

	go updater(monitor, &m)
	go observe(monitor)

	time.Sleep(time.Second * 2)
	fmt.Printf(" -- done; atomics demoII, monitor: %v  \n", monitor.Load())
	fmt.Println()
}
