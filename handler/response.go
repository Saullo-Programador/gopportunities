package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"message":   message,
		"errorCode": statusCode,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"operation": fmt.Sprintf("Operation from handler: %s successfull", op),
		"data":    data,
	})
}