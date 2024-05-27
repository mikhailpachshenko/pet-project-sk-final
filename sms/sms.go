package sms

import (
	"project/model"
	"project/pkg/signs"
	"project/pkg/utils"
	"slices"
	"sort"
	"strings"
)

func GetSMSData(sourcePath string) [][]model.SMSData {
	listSMS := make([]model.SMSData, 0)
	var list [][]model.SMSData // #11.2

	records := utils.ReadFileToStrings(sourcePath) // #2.1 | #2.2

	if len(records) > 0 {
		for i := 0; i < len(records); i++ {
			columns := strings.Split(string(records[i][0]), ";") // #2.3
			if len(columns) == 4 {                               // #2.4 | #2.5 | #2.8
				if signs.FindCountry(columns[0]) && signs.FindProvider(columns[3]) { // #2.6 | #2.7
					listSMS = append(listSMS, model.SMSData{
						Country:      signs.CodeToCountry(columns[0]), // #11.2
						Bandwidth:    columns[1],
						ResponseTime: columns[2],
						Provider:     columns[3],
					})
				}
			}
		}
		sort.SliceStable(listSMS, func(i, j int) bool { return listSMS[i].Provider < listSMS[j].Provider }) // #11.2
		provider := slices.Clone(listSMS)

		sort.SliceStable(listSMS, func(i, j int) bool { return listSMS[i].Country < listSMS[j].Country }) // #11.2
		country := slices.Clone(listSMS)
		list = append(list, provider, country) // #11.2

		/* fmt.Print("\nSMS: ", list, "\n") */
		return list
	}
	return list
}
