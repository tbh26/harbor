package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSome(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	word := "hello"
	var wg sync.WaitGroup
	wg.Add(1)
	go printSome(word, &wg)
	wg.Wait()
	_ = w.Close()
	output, _ := io.ReadAll(r)
	os.Stdout = stdOut // restore standard out
	if !strings.Contains(string(output), word) {
		em := fmt.Sprintf("Excepted %q in the output.", word)
		t.Errorf(em)
	}
}

func Test_challenge(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	words := []string{"foo", "bar"}
	challenge(words)
	_ = w.Close()
	output, _ := io.ReadAll(r)
	os.Stdout = stdOut // restore standard out
	for _, word := range words {
		if !strings.Contains(string(output), word) {
			em := fmt.Sprintf("Excepted %q in the output.", word)
			t.Errorf(em)
		}
	}
}
