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
	billing := resultSet.Billing

	if len(resultSet.SMS) > 0 && len(resultSet.MMS) > 0 && len(resultSet.VoiceCall) > 0 && len(resultSet.Email) > 0 &&
		len(resultSet.Support) > 0 && len(resultSet.Incidents) > 0 && billing.CheckoutPage == true ||
		billing.CheckoutPage == false {

		result.Data = resultSet
		result.Status = true
	} else {
		result.Status = false
		result.Error = "Error on collect data"
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(http.StatusOK, result)

	return
}
