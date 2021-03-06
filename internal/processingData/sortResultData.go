package processingData

import (
	"Diplom/internal/model"
	"sort"
	"sync"
)

func GetResultData() model.ResultSetT {
	var sortResult model.ResultSetT
	var wg sync.WaitGroup

	wg.Add(7)
	go func() {
		sortResult.SMS = SortSMS()
		wg.Done()
	}()
	go func() {
		sortResult.MMS = SortMMS()
		wg.Done()
	}()
	go func() {
		sortResult.VoiceCall = ResultVoiceCall()
		wg.Done()
	}()
	go func() {
		sortResult.EmailSlice = SortEmail()
		wg.Done()
	}()
	go func() {
		sortResult.Billing = ResultBilling()
		wg.Done()
	}()
	go func() {
		sortResult.Support = SortSupport()
		wg.Done()
	}()
	go func() {
		sortResult.Incidents = SortIncident()
		wg.Done()
	}()
	wg.Wait()
	return sortResult
}

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

func SortMMS() [][]model.MMSData {
	var mmsSlice [][]model.MMSData
	var sortCountry []model.MMSData
	var sortProvider []model.MMSData

	mmsData := GetMMS()

	for _, v := range mmsData {
		v.Country = FullCountryFunc(v.Country)
		sortProvider = append(sortProvider, v)
		sortCountry = append(sortCountry, v)
	}
	sort.Slice(sortProvider, func(i, j int) bool { return sortProvider[i].Provider < sortProvider[j].Provider })
	mmsSlice = append(mmsSlice, sortProvider)

	sort.Slice(sortCountry, func(i, j int) bool { return sortCountry[i].Country < sortCountry[j].Country })
	mmsSlice = append(mmsSlice, sortCountry)

	return mmsSlice
}

func SortEmail() [][]model.EmailData {
	var emailSlice [][]model.EmailData
	var sortFast []model.EmailData
	var sortLong []model.EmailData
	var countrySlice []model.EmailData
	emailMap := make(map[string][][]model.EmailData)

	EmailData := ResultEmail()

	s := []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
	i := 0

	for _, v := range EmailData {
		if v.Country == s[i] {
			countrySlice = append(countrySlice, v)

		} else if v.Country != s[i] {

			sortFast = append(sortFast, countrySlice...)
			sortLong = append(sortLong, countrySlice...)

			sort.Slice(sortFast, func(i, j int) bool { return sortFast[i].DeliveryTime < sortFast[j].DeliveryTime })
			emailSlice = append(emailSlice, sortFast[:3])

			sort.Slice(sortLong, func(i, j int) bool { return sortLong[i].DeliveryTime > sortLong[j].DeliveryTime })
			emailSlice = append(emailSlice, sortLong[:3])

			emailMap[s[i]] = append(emailMap[s[i]], emailSlice...)

			i += 1
			sortFast = nil
			sortLong = nil
			countrySlice = nil
			emailSlice = nil
		}
	}
	output := emailMap["RU"]
	return output
}

func SortSupport() []int {
	SupportData := GetSupport()

	timeOneTickets := 3
	var tickets int
	var load int
	var averageResponseTime int
	var supportSlice []int

	for _, v := range SupportData {
		tickets += v.ActiveTickets
	}
	averageResponseTime = tickets * timeOneTickets
	switch {
	case tickets < 9:
		load = 1
	case tickets >= 9 && tickets < 17:
		load = 2
	case tickets > 16:
		load = 3
	}
	supportSlice = append(supportSlice, load)
	supportSlice = append(supportSlice, averageResponseTime)

	return supportSlice
}

func SortIncident() []model.IncidentData {
	incidentData := GetIncident()

	sort.Slice(incidentData, func(i, j int) bool { return incidentData[i].Status < incidentData[j].Status })

	return incidentData
}
