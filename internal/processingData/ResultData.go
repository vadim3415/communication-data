package processingData

import (
	"Diplom/internal/model"
	"fmt"
)

var sortResult model.ResultSetT

func GetResultData() {

	smsData := SortSMS()
	sortResult.SMS = append(sortResult.SMS, smsData...)
	fmt.Println()

}
