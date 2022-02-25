package processingData

import (
	"Diplom/pkg/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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
