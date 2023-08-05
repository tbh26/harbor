package main

import "fmt"

const (
	Answer = 42
	Motto  = "bswg; save the world!"
)

type Month int8

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func main() {
	hello()
	vars()
	more()
	fmt.Println()
}

func hello() {
	fmt.Println("\nBuild systems with GO, save the world. (part 1, chapter 2")
}

func vars() {
	fmt.Printf("\n\n =-= vars =-= \n")
	var n int
	n = 42
	var n2 int = 24
	n3 := 21
	s := "Build systems with GO."
	var s2, s3 string
	s2, s3 = "build systems", "with Go"
	flag1, flag2 := true, false
	//
	fmt.Printf("\n number(s) \n")
	fmt.Printf(" n == %d, n2 == %d, n3 == %d\n", n, n2, n3)
	fmt.Printf("\n string(s) \n")
	fmt.Printf(" s == \"%s\"\n s2 == \"%s\", s3 == \"%s\" \n", s, s2, s3)
	fmt.Printf("\n boolean(s) \n")
	fmt.Printf(" first flag == %t , next flag == %v \n\n", flag1, flag2)
	//
}

func more() {
	fmt.Printf("\n =-= more... =-= (constants & enum) \n")
	fmt.Printf(" answer == %d \n", Answer)
	fmt.Printf(" motto == '%s' \n", Motto)
	//
	for m := Month(January); m <= Month(December); m += 1 {
		// fmt.Printf(" month %02d == %s \n", m, monthToString(m))
		fmt.Printf(" month %s == %s \n", m.toString(), monthToString(m))
	}
}

func (m Month) toString() string {
	return fmt.Sprintf("%02d", m)
}

func monthToString(m Month) string {
	result := "..."
	switch m {
	case January:
		result = "January"
	case February:
		result = "February"
	case March:
		result = "March"
	case April:
		result = "April"
	case May:
		result = "May"
	case June:
		result = "June"
	case July:
		result = "July"
	case August:
		result = "August"
	case September:
		result = "September"
	case October:
		result = "October"
	case November:
		result = "November"
	case December:
		result = "December"
	}
	return result
}
