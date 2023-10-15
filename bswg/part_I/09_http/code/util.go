package code

import (
	"fmt"
	"net/http"
	"strings"
)

func mergeThem(in []string, joiner string) (result string) {
	result = strings.Join(in, joiner)
	return
}

func spaceMerge(in []string) (r string) {
	r = mergeThem(in, " ")
	return
}

func showHeaders(info string, header http.Header) {
	fmt.Printf(" = %s = \n", info)
	for k, v := range header {
		fmt.Printf("%s: %s \n", k, spaceMerge(v))
	}
	fmt.Println()
}
