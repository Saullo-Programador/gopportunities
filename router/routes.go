package router

import (
	"github.com/Saullo-Programador/gopportunities.git/handler"
	"github.com/gin-gonic/gin"
)

func InitalizeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")

	v1.GET("/opening", handler.ShowOpeninghandler)
	v1.POST("/opening",handler.CreateOpeninghandler)
	v1.DELETE("/opening", handler.DeleteOpeninghandler)
	v1.PUT("/opening", handler.UpdateOpeninghandler)
	v1.GET("/opening", handler.ListOpeninghandler)
}