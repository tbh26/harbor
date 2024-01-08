package main

// go mod init github.com/tbh26/harbor/head_wrap/hello
// go build   # creates a binary (named as: basename `pwd` )
// go install   # installs a binary (./hello) into ${GOPATH}/bin/

import (
	"fmt"
	"strings"
)

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
	fmt.Println()
	referenceFun()
	fmt.Println()
	slices()
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

func divNames(fullName string) (first, next string) {
	parts := strings.Split(fullName, " ")
	first, next = parts[0], parts[1]
	return
}

func referenceFun() {
	fullName := "Billy Bob"
	first, next := divNames(fullName)
	fmt.Printf("First name: %s, Next name: %s\n", first, next)

	greet := func(greeter string) {
		fmt.Printf("Hello %s\n", greeter)
	}
	greet(next)

	emphasis := func(subject *string) {
		*subject = fmt.Sprintf("%s!", *subject)
	}
	emphasis(&first)
	greet(first)
}

func slices() {
	// array(s) (fixed length)
	var numbers [3]int
	fmt.Printf("numbers: %v \n", numbers)
	numbers = [3]int{2, 42, 12}
	fmt.Printf("numbers: %v \n", numbers)

	// slices
	numberSlice := []int{12, 42, 21}
	fmt.Printf("number slice: %v \n", numberSlice)
	ns2 := append(numberSlice, 2)
	ns2 = append(ns2, 42, 84, 36)
	fmt.Printf("number slice: %v, ns2: %v \n", numberSlice, ns2)
	for i, val := range ns2 {
		fmt.Printf("[%d]: %d", i, val)
		if i == len(ns2)-1 {
			fmt.Println()
		} else {
			fmt.Print(", ")
		}
	}
}
