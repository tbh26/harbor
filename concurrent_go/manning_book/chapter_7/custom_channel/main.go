package main

import (
	"container/list"
	"fmt"
	"github.com/tbh26/harbor/concurrent_go/manning_book/chapter_7/semaphore"
	"sync"
	"time"
)

type Channel[M any] struct {
	capacitySema *semaphore.Semaphore
	sizeSema     *semaphore.Semaphore
	mutex        sync.Mutex
	buffer       *list.List
}

func NewChannel[M any](capacity int) *Channel[M] {
	return &Channel[M]{
		capacitySema: semaphore.NewSemaphore(capacity),
		sizeSema:     semaphore.NewSemaphore(0),
		buffer:       list.New(),
	}
}

func (c *Channel[M]) Send(message M) {
	c.capacitySema.Acquire()
	c.mutex.Lock()
	c.buffer.PushBack(message)
	c.mutex.Unlock()
	c.sizeSema.Release()
}

func (c *Channel[M]) Receive() M {
	c.capacitySema.Release()
	c.sizeSema.Acquire()
	c.mutex.Lock()
	v := c.buffer.Remove(c.buffer.Front()).(M)
	c.mutex.Unlock()
	return v
}

func receiver(messages *Channel[int], wGroup *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = messages.Receive()
		fmt.Println("Received:", msg)
	}
	wGroup.Done()
}

func demo() {
	channel := NewChannel[int](10)
	wGroup := sync.WaitGroup{}
	wGroup.Add(1)
	go receiver(channel, &wGroup)
	for i := 1; i <= 6; i++ {
		fmt.Println("Sending: ", i)
		channel.Send(i)
	}
	channel.Send(-1)
	wGroup.Wait()
}

func main() {
	fmt.Println("Hello custom channel world!  🇯🇪 ")
	fmt.Println()

	demo()
}
