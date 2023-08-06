package code

import "fmt"

func Demo() {
	fmt.Println("Hello code.Demo() world!")
	//
	arrayDemo()
}

func arrayDemo() {
	//
	var a1 [4]int
	a1 = [4]int{0, 1, 4, 9}
	fmt.Printf(" a1 == %v \n", a1)
	//
	a2 := [3]int{1, 22, 333}
	//	fmt.Printf(" a2 == %v \n", a2)
	for index, value := range a2 {
		fmt.Printf(" [%d] == %03d  (a2)\n", index, value)
	}
}
