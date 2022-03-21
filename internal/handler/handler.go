package handler

import (
	"Diplom/internal/model"
	"Diplom/internal/processingData"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func statusPage(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "./data/status_page.html")

	return
}

func api(c *gin.Context) {
	t := time.Now()
	var result model.ResultT

	resultSet := processingData.GetResultData()
	billing := resultSet.Billing

	if len(resultSet.SMS[1]) > 0 && len(resultSet.MMS[1]) > 0 && len(resultSet.VoiceCall) > 0 && len(resultSet.EmailSlice) > 0 &&
		len(resultSet.Support) > 0 && len(resultSet.Incidents) > 0 && (billing.CheckoutPage == true ||
		billing.CheckoutPage == false) {

		result.Data = resultSet
		result.Status = true

	} else {
		result.Status = false
		result.Error = "Error on collect data"
		result.Data = resultSet
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(http.StatusOK, result)
	fmt.Printf("Latency handler %d ms \n", time.Since(t).Milliseconds())
	return
}
