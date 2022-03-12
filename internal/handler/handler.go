package handler

import (
	"Diplom/internal/model"
	"Diplom/internal/processingData"
	"github.com/gin-gonic/gin"
	"net/http"
)

func get(c *gin.Context) {

	var result model.ResultT

	resultSet := processingData.GetResultData()

	result.Data = resultSet
	result.Status = true

	c.JSON(http.StatusOK, result)
}
