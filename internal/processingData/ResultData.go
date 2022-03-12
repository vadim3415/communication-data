package processingData

import (
	"Diplom/internal/model"
)

var sortResult model.ResultSetT

func GetResultData() model.ResultSetT {

	sortResult.SMS = SortSMS()
	//fmt.Println("\n SMS \n", sortResult.SMS)

	sortResult.MMS = SortMMS()
	//fmt.Println("\n MMS \n", sortResult.MMS)

	sortResult.VoiceCall = ResultVoiceCall()
	//fmt.Println("\n voiceCell \n", sortResult.VoiceCall)

	sortResult.Email = SortEmail()
	//fmt.Println("\n Email \n", sortResult.Email)

	sortResult.Billing = ResultBilling()
	//fmt.Println("\n Billing \n", sortResult.Billing)

	sortResult.Support = SortSupport()
	//fmt.Println("\n Support\n", sortResult.Support)

	sortResult.Incidents = SortIncident()
	//fmt.Println("\n Incident \n", sortResult.Incidents)

	//v, _ := json.Marshal(sortResult)
	//fmt.Println("\nJson\n", string(v))
	return sortResult

}
