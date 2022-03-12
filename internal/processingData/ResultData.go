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
	fmt.Println("\n SMS \n", sortResult.SMS)

	mmsData := SortMMS()
	sortResult.MMS = append(sortResult.MMS, mmsData...)
	fmt.Println("\n MMS \n", sortResult.MMS)

	sortResult.VoiceCall = ResultVoiceCall()
	fmt.Println("\n voiceCell \n", sortResult.VoiceCall)

	sortResult.Billing = ResultBilling()
	fmt.Println("\n Billing \n", sortResult.Billing)

	sortResult.Incidents = SortIncident()
	fmt.Println("\n Incident \n", sortResult.Incidents)

	sortResult.Email = SortEmail()
	fmt.Println("\n Email \n", sortResult.Email)

	v, _ := json.Marshal(sortResult)
	fmt.Println("\nJson\n", string(v))

}
