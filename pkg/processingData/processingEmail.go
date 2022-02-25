package processingData

import (
	"Diplom/pkg/model"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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
