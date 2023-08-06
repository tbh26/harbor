package main

import "fmt"

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
	fmt.Println("done")
}
