package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/golang_oportunities/schemas"
)


// @BasePath /api/v1

// @Summary Update opening
// @Description Update a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Param opening body UpdateOpeningResponse true "Opening data to Update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {

	//ctx.JSON(http.StatusOK, gin.H{"message": "PUT Opening"})
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Error("validation error: " + err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary != 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("failed to update opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "failed to update opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
