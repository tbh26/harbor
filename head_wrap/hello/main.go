package main

// go mod init github.com/tbh26/harbor/head_wrap/hello
// go build   # creates a binary (named as: basename `pwd` )
// go install   # installs a binary (./hello) into ${GOPATH}/bin/

import "fmt"

var (
	Author  = "tbh"
	comment = "bla bla"
)

const (
	Version = "0.1.2"
	greet   = "howdy 👋"
)

func main() {
	fmt.Println()
	fmt.Println("Hello head wrapping go...")
	fmt.Println()

	demoVars()
	fmt.Println()
	controlStructures()
}

func demoVars() {
	var message string = "Hello demo vars"
	fmt.Println(message)

	var x, y, z int = 2, 40, 42
	fmt.Printf("x: %d, y: %d, z: %d \n", x, y, z)

	var s1 = "var without type, using type inference"
	s2 := "short variable declaration syntax demo"
	fmt.Printf("s1: %s, s2: %q \n", s1, s2)

	e, d, f := 123, 456.789, true
	fmt.Printf("e: %v, d: %.2f, f: %t \n", e, d, f)

	explain := "using globals, uppercase global var will be exported"
	comment = "bla"
	fmt.Printf("Author: %#v, comment: %q  (%s) \n", Author, comment, explain)
	fmt.Printf("Version: %q, greet: %q \n", Version, greet)
}

func controlStructures() {
	n, m := 3, 10
	if n == m {
		fmt.Println(" n == m ")
	} else {
		fmt.Println(" n != m ")
	}

	for x := n; x <= m; x += n {
		fmt.Printf(" -  x: %d \n", x)
	}

	for {
		n += n
		fmt.Printf(" =  n: %d \n", n)
		if n > m {
			break
		}
	}
}
