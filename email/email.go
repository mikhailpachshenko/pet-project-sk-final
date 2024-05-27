package email

import (
	"project/model"
	"project/pkg/signs"
	"project/pkg/utils"
	"sort"
	"strconv"
	"strings"
)

func GetMailData(sourcePath string) map[string][][]model.EmailData {
	var listEmail []model.EmailData

	records := utils.ReadFileToStrings(sourcePath) /* #5.1; #5.2 */

	if len(records) > 0 {
		for i := 0; i < len(records); i++ {
			columns := strings.Split(string(records[i][0]), ";") /* #5.3 */
			if len(columns) == 3 {                               /* #5.4; #5.5; #5.9 */
				if signs.FindCountry(columns[0]) && signs.FindProviderEmail(columns[1]) { /* #5.6; #5.7; #5.8 */
					deliveryTime, _ := strconv.ParseInt(columns[2], 10, 64) /* #5.10 */
					listEmail = append(listEmail, model.EmailData{          /* #5.11 */
						Country:      columns[0],
						Provider:     columns[1],
						DeliveryTime: int(deliveryTime),
					})
				}
			}
		}

		var out = make(map[string][][]model.EmailData) /* #11.5 */
		var arrIII [][]model.EmailData
		country := signs.GetCountriesList()

		for _, sign := range country {
			arrI := []model.EmailData{}
			for _, eData := range listEmail {
				if eData.Country == sign {
					arrI = append(arrI, eData)
				}
			}

			sort.SliceStable(arrI, func(i, j int) bool { return arrI[i].DeliveryTime < arrI[j].DeliveryTime })
			fastest := arrI[0:3]
			slowest := arrI[len(arrI)-3:]
			sort.SliceStable(slowest, func(i, j int) bool { return slowest[i].DeliveryTime > slowest[j].DeliveryTime })
			arrIII = append(arrIII, fastest, slowest)
			out[sign] = arrIII
			arrI = nil
			arrIII = nil
		}
		/* fmt.Print("\nEMAIL: ", out, "\n") */
		return out

	} else {
		var out = make(map[string][][]model.EmailData)

		return out
	}
}
