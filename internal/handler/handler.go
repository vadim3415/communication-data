package handler

import (
	"net/http"

	"Diplom/internal/model"
	"Diplom/internal/processingData"

	"github.com/gin-gonic/gin"
)

func statusPage(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "./data/status_page.html")

	return
}

func api(c *gin.Context) {

	var result model.ResultT

	resultSet := processingData.GetResultData()
	//billing := resultSet.Billing

	result.Data = resultSet
	result.Status = true

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(http.StatusOK, result)

	return
}
