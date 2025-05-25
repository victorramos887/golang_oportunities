package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/golang_oportunities/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
