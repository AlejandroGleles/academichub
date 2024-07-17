package main

import (
	"github.com/AlejandroGleles/academichub/config"
	"github.com/AlejandroGleles/academichub/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	//Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	//initialize router
	router.Initialize()
}
