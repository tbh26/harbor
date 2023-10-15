package test

// % go test  -v
// % go fmt  ./...

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	expected := 4
	r := 2 + 2
	if r != expected {
		t.Errorf("Expected %d, got %d.\n", expected, r)
	} else {
		fmt.Printf("Expected %d, and got %d, nice!\n", expected, r)
	}
}

func Test02(t *testing.T) {
	notExpected := 42
	r := 2 + 2
	if r == notExpected {
		t.Errorf("NOT expected %d, and got %d.\n", notExpected, r)
	}
}
