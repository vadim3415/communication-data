package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	ok := "ok"
	c.JSON(http.StatusOK, ok)
}
