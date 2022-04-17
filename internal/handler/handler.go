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

var result model.ResultT

func pullData() {
	resultSet := processingData.GetResultData()

	if len(resultSet.SMS[1]) > 0 && len(resultSet.MMS[1]) > 0 && len(resultSet.VoiceCall) > 0 &&
		len(resultSet.EmailSlice) > 0 && resultSet.Support[1] > 0 && len(resultSet.Incidents) > 0 {

		result.Data = resultSet
		result.Status = true

	} else {
		result.Status = false
		result.Error = "Error on collect data"
	}
}

var counter int = 0
var d = time.NewTicker(30 * time.Second)

func api(c *gin.Context) {
	if counter == 0 {
		pullData()
		counter++
	}

	select {
	case <-d.C:
		pullData()
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.JSON(http.StatusOK, result)
		fmt.Print("update data\n")
		return

	default:
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.JSON(http.StatusOK, result)
		fmt.Print("cash data\n")
		return
	}
}
