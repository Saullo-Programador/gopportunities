package handler

import (
	"net/http"

	"github.com/Saullo-Programador/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)


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
