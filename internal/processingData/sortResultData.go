package processingData

import (
	"Diplom/internal/model"
	"sort"
)

func SortSMS() [][]model.SMSData {

	var smsSlice [][]model.SMSData
	var sortCountry []model.SMSData
	var sortProvider []model.SMSData

	smsData := ResultSMS()

	for _, v := range smsData {
		v.Country = FullCountryFunc(v.Country)
		sortCountry = append(sortCountry, v)
		sortProvider = append(sortProvider, v)
	}

	sort.Slice(sortProvider, func(i, j int) bool { return sortProvider[i].Provider < sortProvider[j].Provider })
	smsSlice = append(smsSlice, sortProvider)

	sort.Slice(sortCountry, func(i, j int) bool { return sortCountry[i].Country < sortCountry[j].Country })
	smsSlice = append(smsSlice, sortCountry)

	return smsSlice
}
