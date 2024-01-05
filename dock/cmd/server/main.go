package main

import (
	"github.com/tbh26/harbor/dock/pkg/config"
	"github.com/tbh26/harbor/dock/pkg/server"
)

func main() {
	sc := config.DefaultServer
	server.Start(sc)
}
