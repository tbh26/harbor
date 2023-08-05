package main

import (
	"errors"
	"fmt"
)

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
	fun()
	valVsRef()
	useError()
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

func fun() {
	fmt.Printf("\n =-= some fun(c) or closure =-= \n")
	a := 20
	b := 1
	c := 2
	//
	r := doit(a*c, c, add)
	fmt.Printf(" first result: %d \n", r)
	//
	r = doit(a+b, c, mul)
	fmt.Printf("  next result: %d \n", r)
	//
	fmt.Printf("  last result: %d \n", sum(a, b, b, a))
	//
	fmt.Println(" =-=-=-=-=-= ")
	//
	fn := accu(1)
	fmt.Printf(" fn()  --> %d \n", fn())
	fmt.Printf(" fn()  --> %d \n", fn())
	fmt.Printf(" fn()  --> %d \n", fn())
	fmt.Printf(" fn()  --> %d \n", fn())
	fn2 := accu(2)
	fmt.Printf(" fn2() --> %d \n", fn2())
	fmt.Printf(" fn2() --> %d \n", fn2())
	fmt.Printf(" fn2() --> %d \n", fn2())
	fmt.Printf(" fn2() --> %d \n", fn2())
}

func add(a int, b int) int {
	return a + b
}

func mul(a int, b int) (result int) {
	result = a * b
	return
}

func doit(n int, m int, f func(int, int) int) int {
	return f(n, m)
}

func sum(numbers ...int) (r int) {
	r = 0
	for _, n := range numbers {
		r += n
	}
	return r
}

func accu(n int) (f func() int) {
	a := 0
	f = func() int {
		r := a
		a += n
		return r
	}
	return f
}

func valVsRef() {
	fmt.Printf("\n =-= regular arg, pointer, nil and zero value =-= \n")
	//
	n := 42
	noDouble(n)
	fmt.Printf(" n == %d \n", n)
	n = 21
	double(&n)
	fmt.Printf(" n == %d \n", n)
	//
	var a int
	var b *int
	var c bool
	var d func()
	var e string
	fmt.Printf(" a == %v, b == %v, c == %v, d == %v, c == '%v' \n", a, b, c, d, e)
}

func noDouble(n int) {
	n *= 2
	// fmt.Printf(" n == %d \n", n)
}

func double(n *int) {
	*n *= 2
}

func useError() {
	fmt.Printf("\n =-= error(s) =-= \n")
	//
	var m *int
	n, e := getValue(m)
	fmt.Printf(" getValue == %v, error == '%v' \n", n, e)
	n = 42
	m = &n
	fmt.Printf(" m == %v, *m == %v \n", m, *m)
	n, e = getValue(m)
	fmt.Printf(" getValue == %v, error == '%v' \n", n, e)
	//
	e = errors.New(fmt.Sprintf("Oops!"))
	fmt.Printf(" error == '%v' \n", e)
}

func getValue(ref *int) (int, error) {
	result := 0
	var err error
	err = nil
	if ref == nil {
		err = errors.New(fmt.Sprintf("int reference nil, No value"))
	} else {
		result = *ref
	}
	return result, err
}
