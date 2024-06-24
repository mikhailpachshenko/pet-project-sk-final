package voice

import (
	"strconv"
	"strings"

	"github.com/mikhailpachshenko/pet-project-sk-final.git/model"
	"github.com/mikhailpachshenko/pet-project-sk-final.git/pkg/signs"
	"github.com/mikhailpachshenko/pet-project-sk-final.git/pkg/utils"
)

func GetVoiceData(sourcePath string) (listVoice []model.VoiceCallData) {
	records := utils.ReadFileToStrings(sourcePath) // #4.1 | #4.2

	if len(records) > 0 {
		for i := 0; i < len(records); i++ {
			columns := strings.Split(string(records[i][0]), ";") // #4.3
			if len(columns) == 8 {                               // #4.4 | #4.5 | #4.8
				if signs.FindCountry(columns[0]) && signs.FindProviderVoice(columns[3]) { // #4.6 | #4.7
					connectionStability, _ := strconv.ParseFloat(columns[4], 32)
					ttfb, _ := strconv.ParseInt(columns[5], 10, 64)
					voicePurity, _ := strconv.ParseInt(columns[6], 10, 64)
					medianOfCallsTime, _ := strconv.ParseInt(columns[7], 10, 64)

					listVoice = append(listVoice, model.VoiceCallData{
						Country:             columns[0],
						Bandwidth:           columns[1],
						ResponseTime:        columns[2],
						Provider:            columns[3],
						ConnectionStability: float32(connectionStability), // #4.10
						TTFB:                int(ttfb),                    // #4.9
						VoicePurity:         int(voicePurity),             // #4.9
						MedianOfCallsTime:   int(medianOfCallsTime),       // #4.9
					})
				}
			}
		}
		/* fmt.Print("\nVOICE: ", listVoice, "\n") */
		return listVoice // #11.4
	}
	return listVoice
}
