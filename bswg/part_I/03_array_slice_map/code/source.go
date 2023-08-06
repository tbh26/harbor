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
}

func arrayDemo() {
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
	a1 := [3]string{"one", "two", "three"} // array
	// s1 := []string{"one", "two", "three"}  // slice
	s1 := a1[:] // slice
	fmt.Printf(" === a1: %T  ( %s ), s1: %T  ( %s ) === \n", a1, reflect.TypeOf(a1), s1, reflect.TypeOf(s1))
	for index, value := range s1 {
		fmt.Printf(" [%d] == %v \n", index, value)
	}
}
