package server

import (
	"github.com/tbh26/harbor/dock/pkg/config"
	"log"
	"net/http"
)

func Start(serverConfig config.ServerStart) {
	address := serverConfig.Address()
	logger := log.Default()
	server := http.Server{Addr: address}
	logger.Printf("server start, address; %s ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	} else {
		logger.Println("server done")
	}
}
