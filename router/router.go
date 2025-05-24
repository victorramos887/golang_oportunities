package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default() // defininco o router

	// definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// executando nossa api
	router.Run(":8090") // executando na porta 8090
}
