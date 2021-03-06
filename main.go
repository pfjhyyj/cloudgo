package main

import (
	"os"

	"./service"

	flag "github.com/spf13/pflag"
)

const (
	// PORT default port 8080
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	// Open a new server default on 8080 port
	server := service.NewServer()
	server.Run(":" + port)
}
