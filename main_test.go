package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	buf := new(bytes.Buffer)

	Greet(buf)

	expected := "Hello, \"harbor server\" greeting.\nHello, Greet from 'client'.\n"
	received := buf.String()
	if received != expected {
		t.Fatalf("== Expected ==\n%q \n== received ==\n%q", expected, received)
	}

}
