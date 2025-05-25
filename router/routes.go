package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/victorramos887/golang_oportunities/docs"
	"github.com/victorramos887/golang_oportunities/handler"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)

	v1.GET("/opening", handler.ShowOpeningHandler)

	v1.POST("/opening", handler.CreateOpeningHandler)

	v1.DELETE("/opening", handler.DeleteOpeningHandler)

	v1.PUT("/opening", handler.UpdateOpeningHandler)

	v1.GET("/openings", handler.ListOpeningHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
