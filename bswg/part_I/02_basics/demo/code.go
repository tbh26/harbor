package demo

import "fmt"

func init() {
	fmt.Println("init demo")
}

func init() {
	fmt.Println("init II (demo)")
}

func Greet(s string) {
	fmt.Printf("Hello '%s'.\n", s)
}
