package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/billing"
	"project/email"
	"project/incident"
	"project/mms"
	"project/model"
	"project/pkg/utils"
	"project/sms"
	"project/support"
	"project/voice"

	"github.com/gorilla/mux"
)

type SourcesData struct {
	ServiceName string `json:"service"`
	ServicePath string `json:"source"`
}

func main() {
	router := mux.NewRouter()                // #9.2
	router.HandleFunc("/", handleConnection) // #9.2
	srv := &http.Server{Handler: router, Addr: "localhost:8080"}
	fmt.Print("The server has been running\n\n")
	log.Fatal((srv.ListenAndServe())) // #9.5

}

func handleConnection(w http.ResponseWriter, r *http.Request) { // #9.3
	var result model.ResultT
	var check bool = true

	data := getResultData()
	check = getCheckServices(data)

	if check {
		result.Status = true
		result.Data = data
		result.Error = ""
	} else {
		result.Status = false
		result.Error = "Error on collet data"
	}

	structToJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(structToJson))
}

func loadConfig(path string) []SourcesData {
	var sourcesData []SourcesData

	byteValue := utils.ReadFileToByte(path)

	json.Unmarshal(byteValue, &sourcesData)

	return sourcesData
}

func getResultData() model.ResultSetT {

	var resultSet model.ResultSetT

	service := loadConfig("service_config.json")

	sms := sms.GetSMSData(findServicePath("sms", service))
	mms := mms.GetMMSData(findServicePath("mms", service))
	voice := voice.GetVoiceData(findServicePath("voice", service))
	email := email.GetMailData(findServicePath("email", service))
	billing := billing.GetBillingData(findServicePath("billing", service))
	support := support.GetSupportData(findServicePath("support", service))
	incident := incident.GetIncidentData(findServicePath("incident", service))

	resultSet = model.ResultSetT{
		SMS:       sms,
		MMS:       mms,
		VoiceCall: voice,
		Email:     email,
		Billing:   billing,
		Support:   support,
		Incidents: incident,
	}

	return resultSet
}

func findServicePath(name string, sourceData []SourcesData) string {
	var path string

	for _, path := range sourceData {
		if path.ServiceName == name {
			return path.ServicePath
		}
	}
	return path
}

func getCheckServices(data model.ResultSetT) bool {

	var check bool = true

	if len(data.SMS) == 0 {
		check = false
	}

	if len(data.MMS) == 0 {
		check = false
	}

	if len(data.VoiceCall) == 0 {
		check = false
	}

	if len(data.Email) == 0 {
		check = false
	}

	if data.Billing == (model.BillingData{}) {
		check = false
	}

	if len(data.Support) == 0 {
		check = false
	}

	if len(data.Incidents) == 0 {
		check = false
	}

	return check
}
