package main

import (
	"fmt"

	"github.com/victorramos887/golang_oportunities/config"
	"github.com/victorramos887/golang_oportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		//panic(err)
		logger.Errorf("config initialization error: %v", err)
		fmt.Println("error:", err)
		return
	}

	router.Initialize() // chamando a funcao Initialize do router
}
