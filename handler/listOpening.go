package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeninghandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg" : "Gets",
	})
}