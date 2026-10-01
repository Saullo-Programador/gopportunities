package main

import (
	"github.com/Saullo-Programador/gopportunities.git/config"
	"github.com/Saullo-Programador/gopportunities.git/router"
)

var(
	logger *config.Logger
)

func main(){

	logger = config.GetLogger("main")
	//Initalize Config
	err := config.Init()
	if err != nil {
		logger.Error("config initialization error: %v",err)
		return
	}

	//Initalize Router
	router.Initalize()
}