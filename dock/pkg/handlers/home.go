package handlers

import (
	"fmt"
	"net/http"
)

func home(rw http.ResponseWriter, rp *http.Request) {
	fmt.Fprintf(rw, "hello '%s' handler\n", "home")
}
