package main

// go mod init github.com/tbh26/harbor/head_wrap/hello
// go build   # creates a binary (named as: basename `pwd` )
// go install   # installs a binary (./hello) into ${GOPATH}/bin/

import "fmt"

func main() {
	fmt.Println("Hello head wrapping go...")
}
