package processingData

import (
	"Diplom/internal/model"
	"encoding/json"
	"fmt"
)

var sortResult model.ResultSetT

func GetResultData() {

	smsData := SortSMS()
	sortResult.SMS = append(sortResult.SMS, smsData...)
	fmt.Println("SMS", sortResult.SMS)

	sortResult.VoiceCall = ResultVoiceCall()
	fmt.Println(sortResult)

	a, _ := json.Marshal(sortResult)
	fmt.Println("Json", string(a))

}
