package main

import "testing"

// run test with (one of):
// - $ go test -race ./race_demo
// - $ make test_race

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!" // is package global
	expectMsg := "Goodbye?"

	wg.Add(2)
	go updateMessage("bla bla bla")
	go updateMessage(expectMsg)
	wg.Wait()

	if msg != expectMsg {
		t.Errorf("expected message: %q \n", expectMsg)
	}
}
