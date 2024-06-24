package incident

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"

	"github.com/mikhailpachshenko/pet-project-sk-final.git/model"
)

func GetIncidentData(sourcePath string) []model.IncidentData {
	var listIncidents []model.IncidentData

	resp, err := http.Get(sourcePath) /* #8.7 */
	if err != nil {
		return listIncidents /* #8.12 */
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 { /* #8.10 */
		content, err := io.ReadAll(resp.Body) /* #8.8 */
		if err != nil {
			return listIncidents /* #8.12 */
		}

		if err := json.Unmarshal(content, &listIncidents); err != nil { /* #8.11 */
			return listIncidents
		}
	} else if resp.StatusCode == 500 { /* #8.10 */
		return listIncidents /* #8.12 */
	}

	sort.SliceStable(listIncidents, func(i, j int) bool { return listIncidents[i].Status < listIncidents[j].Status }) /* #11.8 */
	/* fmt.Print("\nINCIDENT: ", listIncidents, "\n") */
	return listIncidents
}
