package processingData

import (
	"Diplom/internal/model"
	"sort"
)

func SortSMS() [][]model.SMSData {
	smsData := ResultSMS()

	var fullCountrySlice []model.SMSData

	for _, v := range smsData {
		v.Country = FullCountryFunc(v.Country)
		fullCountrySlice = append(fullCountrySlice, v)
	}

	sort.Slice(smsData, func(i, j int) bool { return smsData[i].Provider < smsData[j].Provider })
	sort.Slice(fullCountrySlice, func(i, j int) bool { return fullCountrySlice[i].Country < fullCountrySlice[j].Country })

	var smsSlice model.ResultSetT

	smsSlice.SMS = append(smsSlice.SMS, smsData)
	smsSlice.SMS = append(smsSlice.SMS, fullCountrySlice)

	return smsSlice.SMS
}
