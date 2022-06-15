package main

import (
	"github.com/Dann-Go/InnoTaxiDriverService/internal"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	server := new(internal.Server)
	if err := server.Run(os.Getenv("SERVPORT")); err != nil {
		log.Fatalf("error while running server %s", err.Error())
	}
}
