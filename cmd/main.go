package main

import (
	"github.com/Dann-Go/InnoTaxiDriverService/internal"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/config"
	log "github.com/sirupsen/logrus"
)

// @title           InnoTaxi Driver Microservice
// @version         1.0
// @description     This is a driver microservice for InnoTaxi App.

// @host      localhost:8001
// @BasePath  /

func main() {
	err := config.EnvsCheck()
	serverCfg := config.NewServerConfig()
	if err != nil {
		log.Fatalf("envs are not set %s", err.Error())
	}
	server := new(internal.Server)
	if err := server.Run(serverCfg.Port); err != nil {
		log.Fatalf("error while running server %s", err.Error())
	}
}
