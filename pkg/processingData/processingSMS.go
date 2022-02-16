package processingData

import (
	"Diplom/pkg/model"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ResultSMS() []model.SMSData {
	//читаем файл
	file, err := os.Open("simulator/skillbox-diploma/sms.data")
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
