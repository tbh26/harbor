package main

import (
	"fmt"

	"example.org/bswg/p1/ch2/demo"
)

var x = setX()

func setX() int {
	fmt.Println("set X")
	return 2
}

func init() {
	fmt.Println("init function")
	x += 40
}

func main() {
	fmt.Println("main")
	fmt.Printf(" x == %d \n", x)
	demo.Greet("world!")
	fmt.Println("done")
}
