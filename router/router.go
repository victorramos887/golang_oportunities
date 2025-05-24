package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Initialize Router
	router := gin.Default()

	// initialize Routes
	initializeRoutes(router)
	// Run the server
	router.Run(":8090") // executando na porta 8090
}
