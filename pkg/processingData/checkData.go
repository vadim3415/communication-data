package processingData

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
