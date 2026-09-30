package router

import "github.com/gin-gonic/gin"

func Initalize() {
	r := gin.Default()

	InitalizeRoutes(r)

	r.Run("8080")
}