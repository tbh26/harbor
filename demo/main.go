package main

import (
	"fmt"
	hello2 "github.com/tbh/harbor/demo/client/hello"
	"github.com/tbh/harbor/demo/server/hello"
	"io"
	"os"
)

func Greet(w io.Writer) {
	fmt.Fprintln(w, hello.Greet())
	fmt.Fprintln(w, hello2.Greet())
}

func main() {
	Greet(os.Stdout)
}
