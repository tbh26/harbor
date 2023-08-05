package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	hello()
	inspectArgs()
	argsAdd()
}

func hello() {
	fmt.Println("Build systems with GO, save the world.")
}

func inspectArgs() {
	fmt.Println()
	args := os.Args
	for index, arg := range args {
		fmt.Printf(" [%d], arg: '%s' \n", index, arg)
	}
	fmt.Printf("\n%v\n", args)
}

func argsAdd() {
	args := os.Args[1:]
	total := 0
	for _, arg := range args {
		n, e := strconv.Atoi(arg)
		if e == nil {
			total += n
		}
	}
	fmt.Printf("\n total (sum) == %d \n", total)
}
