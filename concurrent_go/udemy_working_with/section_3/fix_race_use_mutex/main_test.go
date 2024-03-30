package main

import (
	"sync"
	"testing"
)

// run test with (one of):
// - $ go test -v -race ./fix_race_use_mutex
// - $ make test_fix

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!" // is package global
	expectMsg := "Goodbye?"
	expectMsg2 := "bla bla"

	var mutext sync.Mutex
	wg.Add(2)
	go updateMessage(expectMsg, &mutext)
	go updateMessage(expectMsg2, &mutext)
	//go updateMessage("bla?", &mutext)
	wg.Wait()

	//fmt.Println("message:", msg)
	if msg != expectMsg && msg != expectMsg2 {
		t.Errorf("Unexpected message: %q \n", msg)
	}
}
