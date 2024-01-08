package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const (
	defaultVerbose  bool   = false
	defaultHostname string = "localhost"
	defaultPort     uint16 = 2000
	defaultDirPath  string = "."
)

type Config struct {
	Verbose  bool
	Hostname string
	Port     uint16
	DirPath  string
}

func main() {
	config, err := prepareConfig()
	if err != nil {
		fmt.Println("Prepare failure;", err)
		return
	}
	err = httpServe(config)
	if err != nil {
		fmt.Println("http serve failure: ", err)
	}
}

func prepareConfig() (config Config, e error) {
	e = nil
	var p uint = uint(defaultPort)
	config = Config{
		Verbose:  defaultVerbose,
		Hostname: defaultHostname,
		Port:     defaultPort,
		DirPath:  defaultDirPath,
	}

	flag.BoolVar(&config.Verbose, "verbose", config.Verbose, "be verbose")
	flag.BoolVar(&config.Verbose, "v", config.Verbose, "be verbose")
	flag.StringVar(&config.Hostname, "hostname", config.Hostname, "http server hostname")
	flag.StringVar(&config.Hostname, "h", config.Hostname, "http server hostname")
	flag.UintVar(&p, "port", uint(config.Port), "http server port")
	flag.UintVar(&p, "p", uint(config.Port), "http server port")
	flag.StringVar(&config.DirPath, "path", config.DirPath, "http serve content from directory path")
	flag.StringVar(&config.DirPath, "d", config.DirPath, "http serve content from directory path")
	flag.Parse()
	config.Port = uint16(p)

	if _, err := os.Stat(config.DirPath); os.IsNotExist(err) {
		e = err
	}

	return
}

func httpServe(config Config) error {
	hostAddress := fmt.Sprintf("%s:%d", config.Hostname, config.Port)
	path := http.Dir(config.DirPath)
	if config.Verbose {
		fmt.Printf("serve %q on %q \n", path, hostAddress)
	}
	return http.ListenAndServe(hostAddress, http.FileServer(path))
}
