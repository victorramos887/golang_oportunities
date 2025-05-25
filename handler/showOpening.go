package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/golang_oportunities/schemas"
)


// @BasePath /api/v1

// @Summary Show opening
// @Description Show job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {

	//ctx.JSON(http.StatusOK, gin.H{"message": "GET Opening"})
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "get-opening", opening)
}
