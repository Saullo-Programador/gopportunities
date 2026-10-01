package router

import (
	"github.com/Saullo-Programador/gopportunities.git/docs"
	"github.com/Saullo-Programador/gopportunities.git/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitalizeRoutes(router *gin.Engine) {
	//Initialize Handler

	handler.InitializeHandler()
	basePath := "/api/v1"
	
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeninghandler)
		v1.POST("/opening", handler.CreateOpeninghandler)
		v1.DELETE("/opening", handler.DeleteOpeninghandler)
		v1.PUT("/opening", handler.UpdateOpeninghandler)
		v1.GET("/openings", handler.ListOpeninghandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
