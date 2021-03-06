package processingData

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

func CheckCountryFunc(country string) string {
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

func CheckProviderFunc(provider string) string {
	output := ""
	ProviderMap := map[string]string{
		"Topolo":           "Topolo",
		"Rond":             "Rond",
		"Kildy":            "Kildy",
		"TransparentCalls": "TransparentCalls",
		"E-Voice":          "E-Voice",
		"JustPhone":        "JustPhone",
		"Gmail":            "Gmail",
		"Yahoo":            "Yahoo",
		"Hotmail":          "Hotmail",
		"MSN":              "MSN",
		"Orange":           "Orange",
		"Comcast":          "Comcast",
		"AOL":              "AOL",
		"Live":             "Live",
		"RediffMail":       "RediffMail",
		"GMX":              "GMX",
		"Protonmail":       "Protonmail",
		"Yandex":           "Yandex",
		"Mail.ru":          "Mail.ru",
	}
	for _, v := range ProviderMap {
		if provider == v {
			output = v
		}
	}
	return output
}

func FullCountryFunc(country string) string {
	output := ""
	CountryMap := map[string]string{
		"RU": "Russian Federation",
		"US": "United States of America",
		"GB": "United Kingdom of Great Britain and Northern Ireland",
		"FR": "France",
		"BL": "Saint Barthélemy",
		"AT": "Austria",
		"BG": "Bulgaria",
		"DK": "Denmark",
		"CA": "Canada",
		"ES": "Spain",
		"CH": "Switzerland",
		"TR": "Turkey",
		"PE": "Peru",
		"NZ": "New Zealand",
		"MC": "Monaco",
	}
	for i, v := range CountryMap {
		if country == i {
			output = v
		}
	}
	return output
}

func convertingFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		logrus.Println(err)
	}
	return float32(f)
}

func convertingInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		logrus.Println(err)
	}
	return i
}

func convertingBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		logrus.Println(err)
		return false
	}
	return b
}
