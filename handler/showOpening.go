package handler

import (
	"net/http"

	"github.com/Saullo-Programador/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show Opening
// @Description Show a job opening by ID
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeninghandler(ctx *gin.Context){
	id := ctx.Query("id")

	if id == ""{
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return 
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err!= nil{
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}
	sendSuccess(ctx, "show opening", opening)
}
