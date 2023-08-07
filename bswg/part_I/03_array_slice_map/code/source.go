package code

import (
	"fmt"
	"reflect"
)

func Demo() {
	fmt.Println("Hello code.Demo() world!")
	//
	arrayDemo()
	sliceDemo()
	mapDemo()
	//
	fmt.Println()
}

func arrayDemo() {
	fmt.Println()
	//
	var a1 [4]int
	a1 = [4]int{0, 1, 4, 9}
	fmt.Printf(" a1 == %v (%T)\n", a1, a1)
	//
	a2 := [3]int{1, 22, 333}
	//	fmt.Printf(" a2 == %v \n", a2)
	for index, value := range a2 {
		fmt.Printf(" [%d] == %03d  (a2)\n", index, value)
	}
}

func sliceDemo() {
	fmt.Println()
	//
	a1 := [3]string{"one", "two", "three"} // array
	// s1 := []string{"one", "two", "three"}  // slice
	s1 := a1[:] // slice
	fmt.Printf(" === a1: %T  ( %s %s ), s1: %T  ( %s %s ) === \n", a1, reflect.TypeOf(a1), reflect.ValueOf(a1).Kind(), s1, reflect.TypeOf(s1), reflect.ValueOf(s1).Kind())
	for index, value := range s1 {
		fmt.Printf(" [%d] == %v \n", index, value)
	}
	fmt.Printf(" s1 == '%v', len: %d, cap: %d \n", s1, len(s1), cap(s1))
	s2 := append(s1, "four")
	fmt.Printf(" s2 == '%v', len: %d, cap: %d \n", s2, len(s2), cap(s2))
	//
	s3 := make([]int, 5, 10)
	s3[1], s3[2], s3[3] = 1, 2, 1
	fmt.Printf(" s3 == '%v', len: %d, cap: %d \n", s3, len(s3), cap(s3))
	// s3[5] = 121 <- panic: runtime error: index out of range [5] with length 5
	s4 := append(s3, 121)
	fmt.Printf(" s4 == '%v', len: %d, cap: %d \n", s4, len(s4), cap(s4))
	s4 = append(s4, 212)
	fmt.Printf(" s4 == '%v', len: %d, cap: %d \n", s4, len(s4), cap(s4))
	//
	s4Len := len(s4)
	s5 := make([]int, s4Len)
	for i, v := range s4 {
		s5[s4Len-(i+1)] = v
	}
	fmt.Printf(" s5 == '%v', len: %d, cap: %d \n", s5, len(s5), cap(s5))
}

func mapDemo() {
	fmt.Println()
	//
	var m1 map[string]int
	//
	// m1["one"] = 1 <-panic: assignment to entry in nil map
	m1 = make(map[string]int, 5)
	fmt.Printf(" m1 == '%v', len: %d \n", m1, len(m1))
	m1["one"] = 1
	fmt.Printf(" m1 == '%v', len: %d \n", m1, len(m1))
	m1["two"] = 22
	fmt.Printf(" m1 == '%v', len: %d \n", m1, len(m1))
	m1 = map[string]int{
		"three": 333,
		"four":  4,
		"five":  55,
	}
	fmt.Printf(" m1 == '%v', len: %d \n", m1, len(m1))
	//
	key := "three"
	val, found := m1[key]
	fmt.Printf(" m1[%s], val: '%v', found: %v \n", key, val, found)
	key = "two"
	val, found = m1[key]
	fmt.Printf(" m1[%s], val: '%v', found: %v \n", key, val, found)
	//
	key = "five"
	delete(m1, key)
	fmt.Printf(" m1 == '%v', len: %d \n", m1, len(m1))
	//
	m1["six"] = 6
	m1["seven"] = 7
	m1["eight"] = 8
	m1["nine"] = 9
	fmt.Printf(" m1 == '%v', len: %d \n", m1, len(m1))
	for k, v := range m1 {
		fmt.Printf(" m1[%s] == %d \n", k, v)
	}
	//
}
