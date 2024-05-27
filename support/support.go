package support

import (
	"encoding/json"
	"io"
	"net/http"
	"project/model"
)

func GetSupportData(sourcePath string) []int {
	var empty []int

	resp, err := http.Get(sourcePath) // #7.1
	if err != nil {
		return empty // #7.4
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 { // #7.4
		content, err := io.ReadAll(resp.Body) // #7.2
		if err != nil {
			return empty // #7.4
		}

		var listSupport []model.SupportData                           // #7.3
		if err := json.Unmarshal(content, &listSupport); err != nil { // #7.5
			return empty // #7.4
		}

		var general []int // #11.7
		var sumTicket int = 0
		var ticketPer int = (60 / 18)

		support := listSupport
		for _, sData := range support {
			sumTicket += sData.ActiveTickets
		}

		if sumTicket <= 8 {
			general = append(general, 1)
		} else if sumTicket >= 9 && sumTicket <= 16 {
			general = append(general, 2)
		} else if sumTicket > 16 {
			general = append(general, 3)
		}

		general = append(general, ticketPer*sumTicket)

		/* fmt.Print("\nSUPPORT: ", general, "\n") */
		return general
	}
	return empty // #7.4
}
