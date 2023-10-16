package config

import "fmt"

type ServerStart struct {
	Host string
	Port uint16
}

var DefaultServer ServerStart = ServerStart{Host: "localhost", Port: 8080}

func (s ServerStart) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
