package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_income(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	calcIncome()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	expectedResult := "Final bank balance: $37440.00"
	if !strings.Contains(output, expectedResult) {
		t.Errorf("Expected result (%q) NOT in output. \n", expectedResult)
		t.Errorf(" -- output:\n\n%s\n", output)
	}
}
