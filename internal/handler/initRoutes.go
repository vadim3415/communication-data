package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/simulator", "./simulator")

	AllGroup := router.Group("/")
	{
		AllGroup.GET("/", api)
		AllGroup.GET("/api", api)
	}
	return router
}
