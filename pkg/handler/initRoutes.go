package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.New()

	AllGroup := router.Group("/")
	{
		AllGroup.GET("/", getMMS)
	}
	return router
}