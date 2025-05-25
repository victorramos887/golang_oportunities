package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/golang_oportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	//logger.Infof("request received: %+v", request)

	if err := request.Validate(); err != nil {
		//logger.Errorf("validation error: %v", err)
		// ctx.AbortWithStatusJSON(400, gin.H{
		// 	"error": err.Error(),
		// })
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("erro creating opening: %+v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "erro creting opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
