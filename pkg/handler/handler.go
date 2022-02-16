package handler

import (
	"Diplom/pkg/model"
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
		checkCountry := checkCountryFunc(v.Country)
		checkProvider := checkProviderFunc(v.Provider)

		if len(checkCountry) > 0 && len(checkProvider) > 0 {
			resultSlice = append(resultSlice, v)
		}
	}

	fmt.Println(resultSlice)
	c.JSON(http.StatusOK, resultSlice)
	return
}

func checkCountryFunc(country string) string {
	output := ""
	mmsCountryMap := map[string]string{
		"RU": "RU",
		"US": "US",
		"GB": "GB",
		"FR": "FR",
		"BL": "BL",
		"AT": "AT",
		"BG": "BG",
		"DK": "DK",
		"CA": "CA",
		"ES": "ES",
		"CH": "CH",
		"TR": "TR",
		"PE": "PE",
		"NZ": "NZ",
		"MC": "MC",
	}
	for _, v := range mmsCountryMap {
		if country == v {
			output = v
		}
	}
	return output
}

func checkProviderFunc(provider string) string {
	output := ""
	mmsProviderMap := map[string]string{
		"Topolo": "Topolo",
		"Rond":   "Rond",
		"Kildy":  "Kildy",
	}
	for _, v := range mmsProviderMap {
		if provider == v {
			output = v
		}
	}
	return output
}
