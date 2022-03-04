package handler

import (
	"Diplom/pkg/model"
	"Diplom/pkg/processingData"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func get(c *gin.Context) {
	ok := "ok"
	c.JSON(http.StatusOK, ok)
}

func getMMS(c *gin.Context) {
	var JsonSliceMMS []model.MMSData
	var resultSliceMMS []model.MMSData
	var nilSliceMMS []model.MMSData

	statusCode := c.Writer.Status()
	if statusCode == 200 {
		resp, err := http.Get("http://127.0.0.1:8383/mms")
		if err != nil {
			c.JSON(http.StatusBadRequest, nilSliceMMS)
			logrus.Println(err)
			return
		}

		textBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, nilSliceMMS)
			logrus.Println(err)
			return
		}
		defer resp.Body.Close()

		if err := json.Unmarshal(textBytes, &JsonSliceMMS); err != nil {
			c.JSON(http.StatusBadRequest, JsonSliceMMS)
			logrus.Println(err)
			return
		}

		for _, v := range JsonSliceMMS {
			checkCountry := processingData.CheckCountryFunc(v.Country)
			checkProvider := processingData.CheckProviderFunc(v.Provider)

			if len(checkCountry) > 0 && len(checkProvider) > 0 {
				resultSliceMMS = append(resultSliceMMS, v)
			}
		}
		fmt.Println(resultSliceMMS)
		c.JSON(http.StatusOK, resultSliceMMS)
		//resultSliceMMS = nil
		//JsonSliceMMS = nil
		return
	}
	c.JSON(http.StatusInternalServerError, nilSliceMMS)
}

func getSupport(c *gin.Context) {
	var JsonSliceSupport []model.SupportData
	var resultSliceSupport []model.SupportData
	var nilSliceSupport []model.SupportData

	statusCode := c.Writer.Status()
	if statusCode == 200 {
		resp, err := http.Get("http://127.0.0.1:8383/support")
		if err != nil {
			c.JSON(http.StatusBadRequest, nilSliceSupport)
			logrus.Println(err)
			return
		}

		textBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, nilSliceSupport)
			logrus.Println(err)
			return
		}
		defer resp.Body.Close()

		if err := json.Unmarshal(textBytes, &JsonSliceSupport); err != nil {
			c.JSON(http.StatusBadRequest, JsonSliceSupport)
			logrus.Println(err)
			return
		}

		for _, v := range JsonSliceSupport {
			resultSliceSupport = append(resultSliceSupport, v)
		}

		fmt.Println(resultSliceSupport)
		c.JSON(http.StatusOK, resultSliceSupport)
		//resultSliceSupport = nil
		//JsonSliceSupport = nil
		return
	}
	c.JSON(http.StatusInternalServerError, nilSliceSupport)
}

func getIncident(c *gin.Context) {
	var JsonSliceIncident []model.IncidentData
	var resultSliceIncident []model.IncidentData
	var nilSliceIncident []model.IncidentData

	statusCode := c.Writer.Status()
	if statusCode == 200 {
		resp, err := http.Get("http://127.0.0.1:8383/accendent")
		if err != nil {
			c.JSON(http.StatusBadRequest, nilSliceIncident)
			logrus.Println(err)
			return
		}

		textBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, nilSliceIncident)
			logrus.Println(err)
			return
		}
		defer resp.Body.Close()

		if err := json.Unmarshal(textBytes, &JsonSliceIncident); err != nil {
			c.JSON(http.StatusBadRequest, JsonSliceIncident)
			logrus.Println(err)
			return
		}

		for _, v := range JsonSliceIncident {
			resultSliceIncident = append(resultSliceIncident, v)
		}

		fmt.Println(resultSliceIncident)
		c.JSON(http.StatusOK, resultSliceIncident)
		//resultSliceIncident = nil
		//JsonSliceIncident = nil
		return
	}
	c.JSON(http.StatusInternalServerError, nilSliceIncident)
}
