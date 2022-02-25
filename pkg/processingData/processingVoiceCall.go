package processingData

import (
	"Diplom/pkg/model"

	"io/ioutil"
	"log"
	"os"
	"strings"
)

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
