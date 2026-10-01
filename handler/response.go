package handler

import (
	"fmt"
	"net/http"

	"github.com/Saullo-Programador/gopportunities.git/schemas"
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

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message	 string  `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message	 string  `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}

type ListOpeningResponse struct {
	Message	 string  `json:"message"`
	Data []schemas.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message	 string  `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}