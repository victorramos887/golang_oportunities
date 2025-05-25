package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/golang_oportunities/schemas"
)



// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	//ctx.JSON(http.StatusOK, gin.H{"message": "DELETE Opening"})
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

	if err := db.Delete(&opening).Error; err!=nil{
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
