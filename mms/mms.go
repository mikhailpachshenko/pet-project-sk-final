package mms

import (
	"encoding/json"
	"io"
	"net/http"
	"project/model"
	"project/pkg/signs"
	"slices"
	"sort"
)

func GetMMSData(sourcePath string) (sortLists [][]model.MMSData) {
	resp, err := http.Get(sourcePath) // #3.1
	if err != nil {
		return sortLists
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 { // #3.4
		content, err := io.ReadAll(resp.Body) // #3.2
		if err != nil {
			return sortLists
		}
		var mms []model.MMSData
		if err := json.Unmarshal(content, &mms); err != nil { // #3.3; #3.5
			return sortLists // #3.8
		}

		var listMMS []model.MMSData
		for _, val := range mms {
			if signs.FindCountry(val.Country) && signs.FindProvider(val.Provider) { // #3.6; #3.7
				val.Country = signs.CodeToCountry(val.Country)
				listMMS = append(listMMS, val)
			}
		}

		sort.SliceStable(listMMS, func(i, j int) bool { return listMMS[i].Provider < listMMS[j].Provider }) // #11.3
		providers := slices.Clone(listMMS)

		sort.SliceStable(listMMS, func(i, j int) bool { return listMMS[i].Country < listMMS[j].Provider }) // #11.3
		countrys := slices.Clone(listMMS)

		sortLists = append(sortLists, providers, countrys)
		/* fmt.Print("\nMMS: ", sortLists, "\n") */
		return sortLists
	}
	return sortLists
}
