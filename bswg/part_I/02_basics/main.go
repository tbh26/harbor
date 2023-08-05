package main

import "fmt"

func main() {
	hello()
	vars()
}

func hello() {
	fmt.Println("Build systems with GO, save the world. (part 1, chapter 2")
}

func vars() {
	var n int
	n = 42
	var n2 int = 24
	n3 := 21
	s := "Build systems with GO."
	var s2, s3 string
	s2, s3 = "build systems", "with Go"
	flag1, flag2 := true, false
	//
	fmt.Printf("\n\n number(s) \n")
	fmt.Printf(" n == %d, n2 == %d, n3 == %d\n", n, n2, n3)
	fmt.Printf("\n string(s) \n")
	fmt.Printf(" s == \"%s\"\n s2 == \"%s\", s3 == \"%s\" \n", s, s2, s3)
	fmt.Printf("\n boolean(s) \n")
	fmt.Printf(" first flag == %t , next flag == %v \n\n", flag1, flag2)
	//
}
