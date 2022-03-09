package processingData

import (
	"Diplom/internal/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func ResultSMS() []model.SMSData {
	//читаем файл
	file, err := os.Open("sms.data")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()
	//открываем файл
	readFile, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// делим содержимое на строки
	line := strings.Split(string(readFile), "\n")

	var callSMS model.SMSData
	var sliceSMS []model.SMSData
	// цикл работает пока есть строки
	for i := 0; i < len(line); i++ {
		// считаем колличество разделителей
		checkLine := strings.Count(line[i], ";")
		// проверяем целостность строки
		if checkLine == 3 {
			splitLine := strings.Split(line[i], ";")

			callSMS = model.SMSData{
				Country:      splitLine[0],
				Bandwidth:    splitLine[1],
				ResponseTime: splitLine[2],
				Provider:     splitLine[3],
			}
			// проверяем страну
			checkCountry := CheckCountryFunc(splitLine[0])
			//проверяем провайдера
			checkProvider := CheckProviderFunc(splitLine[3])
			// если страна и провайдер прошли проверку, записываем в срез
			if splitLine[0] == checkCountry && splitLine[3] == checkProvider {
				sliceSMS = append(sliceSMS, callSMS)
			}
		}
	}
	return sliceSMS
}

func GetMMS() []model.MMSData {
	var JsonSliceMMS []model.MMSData
	var resultSliceMMS []model.MMSData
	var nilSliceMMS []model.MMSData

	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil || resp.StatusCode != 200 {
		logrus.Println("ошибка", err)
		return nilSliceMMS
	}

	textBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Println(err)
		return nilSliceMMS
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(textBytes, &JsonSliceMMS); err != nil {
		logrus.Println(err, nilSliceMMS)
		return nilSliceMMS
	}

	for _, v := range JsonSliceMMS {
		checkCountry := CheckCountryFunc(v.Country)
		checkProvider := CheckProviderFunc(v.Provider)

		if len(checkCountry) > 0 && len(checkProvider) > 0 {
			resultSliceMMS = append(resultSliceMMS, v)
		}
	}

	return resultSliceMMS
}

func ResultVoiceCall() []model.VoiceCallData {
	file, err := os.Open("voice.data")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	readFile, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	line := strings.Split(string(readFile), "\n")
	var VoiceCall model.VoiceCallData
	var sliceVoiceCall []model.VoiceCallData

	for i := 0; i < len(line); i++ {

		checkLine := strings.Count(line[i], ";")

		if checkLine == 7 {
			splitLine := strings.Split(line[i], ";")

			VoiceCall = model.VoiceCallData{
				Country:             splitLine[0],
				Bandwidth:           splitLine[1],
				ResponseTime:        splitLine[2],
				Provider:            splitLine[3],
				ConnectionStability: convertingFloat32(splitLine[4]),
				TTFB:                convertingInt(splitLine[5]),
				VoicePurity:         convertingInt(splitLine[6]),
				MedianOfCallsTime:   convertingInt(splitLine[7]),
			}
			checkCountry := CheckCountryFunc(splitLine[0])
			checkProvider := CheckProviderFunc(splitLine[3])

			if splitLine[0] == checkCountry && splitLine[3] == checkProvider {
				sliceVoiceCall = append(sliceVoiceCall, VoiceCall)
			}

		}
	}
	return sliceVoiceCall
}

func ResultEmail() []model.EmailData {
	file, err := os.Open("email.data")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	readFile, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	line := strings.Split(string(readFile), "\n")
	var Email model.EmailData
	var sliceEmail []model.EmailData

	for i := 0; i < len(line); i++ {

		checkLine := strings.Count(line[i], ";")

		if checkLine == 2 {
			splitLine := strings.Split(line[i], ";")

			Email = model.EmailData{
				Country:      splitLine[0],
				Provider:     splitLine[1],
				DeliveryTime: convertingInt(splitLine[2]),
			}
			checkCountry := CheckCountryFunc(splitLine[0])
			checkProvider := CheckProviderFunc(splitLine[1])

			if splitLine[0] == checkCountry && splitLine[1] == checkProvider {
				sliceEmail = append(sliceEmail, Email)
			}

		}
	}
	return sliceEmail
}

func ResultBilling() []model.BillingData {
	file, err := os.Open("billing.data")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	readFile, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Fatal(err)
	}

	line := strings.Split(string(readFile), "\n")

	var Billing model.BillingData
	var sliceBilling []model.BillingData

	for i := 0; i < len(line); i++ {

		splitLine := strings.Split(line[i], "")
		lenSplitLine := len(splitLine)

		Billing = model.BillingData{
			CreateCustomer: convertingBool(splitLine[lenSplitLine-1]),
			Purchase:       convertingBool(splitLine[lenSplitLine-2]),
			Payout:         convertingBool(splitLine[lenSplitLine-3]),
			Recurring:      convertingBool(splitLine[lenSplitLine-4]),
			FraudControl:   convertingBool(splitLine[lenSplitLine-5]),
			CheckoutPage:   convertingBool(splitLine[lenSplitLine-6]),
		}

		sliceBilling = append(sliceBilling, Billing)

	}
	return sliceBilling
}

func GetSupport() []model.SupportData {

	var resultSliceSupport []model.SupportData
	var nilSliceSupport []model.SupportData

	resp, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil || resp.StatusCode != 200 {
		logrus.Println(err)
		return nilSliceSupport
	}

	textBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Println(err)
		return nilSliceSupport
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(textBytes, &resultSliceSupport); err != nil {
		logrus.Println(err)
		return nilSliceSupport
	}

	return resultSliceSupport
}

func GetIncident() []model.IncidentData {
	var resultSliceIncident []model.IncidentData
	var nilSliceIncident []model.IncidentData

	resp, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		logrus.Println(err)
		return nilSliceIncident
	}

	textBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Println(err)
		return nilSliceIncident
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(textBytes, &resultSliceIncident); err != nil {
		logrus.Println(err)
		return nilSliceIncident
	}

	return resultSliceIncident
}
