package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/web", "./web")

	AllGroup := router.Group("/")
	{
		AllGroup.GET("/", statusPage)
		AllGroup.GET("/api", api)
	}
	return router
}
