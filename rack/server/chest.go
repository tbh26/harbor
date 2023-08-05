package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	port  int
	debug bool
}

type Server struct {
	c   Config
	app *fiber.App
}

func Prepare() Server {
	config := Config{port: 5000, debug: true}
	app := fiber.New()
	server := Server{
		c:   config,
		app: app,
	}
	return server
}

func (s Server) Startup() error {
	address := fmt.Sprintf(":%d", s.c.port)
	e := s.app.Listen(address)
	return e
}
