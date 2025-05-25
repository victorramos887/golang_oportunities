package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/golang_oportunities/schemas"
)



// @BasePath /api/v1

// @Summary List opening
// @Description List job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ListOpeningHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
