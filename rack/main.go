package main

import (
	"fmt"
	"github.com/tbh26/harbor/rack/server"
)

func main() {
	fmt.Println("hello rack")
	s := server.Prepare()

	s.Startup()
}
