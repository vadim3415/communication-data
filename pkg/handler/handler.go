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

var mmsJsonSlice []model.MMSData
var resultSlice []model.MMSData

func getMMS(c *gin.Context) {

	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		c.JSON(http.StatusBadRequest, resultSlice)
		logrus.Println(err)
		return
	}

	textBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultSlice)
		logrus.Println(err)
		return
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(textBytes, &mmsJsonSlice); err != nil {
		c.JSON(http.StatusBadRequest, resultSlice)
		logrus.Println(err)
		return
	}

	for _, v := range mmsJsonSlice {
		checkCountry := processingData.CheckCountryFunc(v.Country)
		checkProvider := processingData.CheckProviderFunc(v.Provider)

		if len(checkCountry) > 0 && len(checkProvider) > 0 {
			resultSlice = append(resultSlice, v)
		}
	}

	fmt.Println(resultSlice)
	c.JSON(http.StatusOK, resultSlice)
	return
}