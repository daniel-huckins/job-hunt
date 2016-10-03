package main

import (
	"os"
)

func main() {
	server := newServer()

	pwd, err := os.Getwd()
	if err != nil {
		log.WithError(err).Fatal("error getting wd")
	}

	log.Info(pwd)

	server.Run(":8080")
}
