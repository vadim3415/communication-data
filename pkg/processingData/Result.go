package processingData

import (
	"Diplom/pkg/model"
	"fmt"
	"sort"
)

func GetResultData() {
	smsData := ResultSMS()

	var fullCountrySlice []model.SMSData

	for _, v := range smsData {
		v.Country = FullCountryFunc(v.Country)
		fullCountrySlice = append(fullCountrySlice, v)
	}

	sort.Slice(smsData, func(i, j int) bool { return smsData[i].Provider < smsData[j].Provider })
	sort.Slice(fullCountrySlice, func(i, j int) bool { return fullCountrySlice[i].Country < fullCountrySlice[j].Country })

	var a model.ResultSetT

	a.SMS = append(a.SMS, smsData)
	a.SMS = append(a.SMS, fullCountrySlice)
	fmt.Println("aaa", a)

}
