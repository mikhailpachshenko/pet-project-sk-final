package billing

import (
	"project/model"
	"project/pkg/utils"
)

func GetBillingData(sourcePath string) model.BillingData {
	var billing model.BillingData
	byteValue := utils.ReadFileToByte(sourcePath) /* #6.1; #6.2 */

	if len(byteValue) == 6 {
		billing = model.BillingData{
			CreateCustomer: bitToBool(byteValue[5]),
			Purchase:       bitToBool(byteValue[4]),
			Payout:         bitToBool(byteValue[3]),
			Recurring:      bitToBool(byteValue[2]),
			FraudControl:   bitToBool(byteValue[1]),
			CheckoutPage:   bitToBool(byteValue[0]),
		}
		/* fmt.Print("\nBILLING: ", billing, "\n") */
		return billing /* #11.6 */
	}
	return billing
}

func bitToBool(line byte) bool {
	if line == 49 { /* #6.5 */
		return true
	} else {
		return false
	}
}

/*
func toByte(line []byte) bool {
	var amount uint8
	for j := len(line) - 1; j < 0; j-- { #6.3
		if line[j] == 49 {
			numb, _ := strconv.ParseUint(string(line[j]), 10, 8)
			numb8 := uint8(numb)
			amount += func(i uint8, j int) uint8 { return uint8(math.Pow(float64(i), float64(j))) }(numb8, j)  #6.4
		}
	}
	return false
} */
