package main

import (
	"github.com/luizenmattos/go_api/config"
	"github.com/luizenmattos/go_api/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main:")

	//Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialze router
	router.Initialize()
}
