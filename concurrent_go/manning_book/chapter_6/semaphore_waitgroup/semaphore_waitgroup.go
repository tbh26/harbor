package semaphore_waitgroup

import (
	"github.com/tbh26/harbor/concurrent_go/manning_book/chapter_5/semaphore"
)

type WaitGrp struct {
	sema *semaphore.Semaphore
}

func NewWaitGrp(size int) *WaitGrp {
	return &WaitGrp{sema: semaphore.NewSemaphore(1 - size)}
}

func (wg *WaitGrp) Wait() {
	wg.sema.Acquire()
}

func (wg *WaitGrp) Done() {
	wg.sema.Release()
}
