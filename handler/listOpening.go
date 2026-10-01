package handler

import (
	"net/http"

	"github.com/Saullo-Programador/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeninghandler(ctx *gin.Context){
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil{
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list openings", openings)
}