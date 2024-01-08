package main

// go mod init github.com/tbh26/harbor/head_wrap/hello
// go build   # creates a binary (named as: basename `pwd` )
// go install   # installs a binary (./hello) into ${GOPATH}/bin/

import (
	"encoding/json"
	"fmt"
	"os"
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
	fmt.Println()
	maps()
	fmt.Println()
	jsonConfig()
	fmt.Println()
	useInterfaces()
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
	fmt.Printf("number slice: %v, ns2: %#v \n", numberSlice, ns2)
	for i, val := range ns2 {
		fmt.Printf("[%d]: %d", i, val)
		if i == len(ns2)-1 {
			fmt.Println()
		} else {
			fmt.Print(", ")
		}
	}
}

func maps() {
	// maps
	scores := map[string]int{
		"Alice":   86,
		"Bob":     62,
		"Charlie": 90,
		"Debby":   77,
	}
	scores["Ellie"] = 72
	//fmt.Printf("scores: %v\n", scores)
	fmt.Printf("scores: %#v\n", scores)
	delete(scores, "Peter") // no-op
	delete(scores, "Charlie")
	//
	counter := 0
	entriesCount := len(scores)
	for k, v := range scores {
		fmt.Printf("%s: %d", k, v)
		counter += 1
		if counter != entriesCount {
			fmt.Print(", ")
		} else {
			fmt.Println(".")
		}
	}
}

func loadJson(path string) (result map[string]interface{}, e error) {
	e = nil
	//data = map[string]interface{}
	data, err := os.ReadFile(path)
	if err != nil {
		//panic(err)
		e = err
		return
	}
	e = json.Unmarshal(data, &result)
	return
}

func jsonConfig() {
	dataPath := "data/example.json"
	data, err := loadJson(dataPath)
	if err != nil {
		fmt.Println("data loading failed:", err)
		return
	}
	fmt.Printf("data: %#v \n", data)

	scores := data["scores"].(map[string]interface{})
	aliceScore := scores["Alice"].(float64)
	fmt.Printf("scores: %#v, alice(score): %d \n", scores, int(aliceScore))
}

type Cat struct {
	name string
}

func (c Cat) Pet() {
	fmt.Printf("%q: Prrr..\n", c.name)
}

func (c Cat) Name() string {
	return c.name
}

type Dog struct {
	name string
}

func (d Dog) Pet() {
	fmt.Printf("%q barks: woef woef!\n", d.name)
}

func (d Dog) Name() string {
	return d.name
}

type Animal interface {
	Pet()
	Name() string
}

func compilment(a Animal) {
	fmt.Printf("Good job,2020 %s! \n", a.Name())
	a.Pet()
}

func useInterfaces() {
	c := Cat{"Alice"}
	compilment(c)

	d := Dog{"Bobby"}
	compilment(d)
}
