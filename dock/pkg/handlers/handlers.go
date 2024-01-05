package handlers

import "net/http"

func GetHandlers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	return mux
}
