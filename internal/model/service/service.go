package service

type SMSData struct { // #2.4; #2.9
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

type MMSData struct { // #3.3
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type VoiceCallData struct { // #4.11
	Country             string
	Bandwidth           string
	ResponseTime        string
	Provider            string
	ConnectionStability float32
	TTFB                int
	VoicePurity         int
	MedianOfCallsTime   int
}

type EmailData struct { // #5.11
	Country      string
	Provider     string
	DeliveryTime int
}

type BillingData struct { // #6.6
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

type SupportData struct { // #7.3
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type IncidentData struct { // #8.9
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

type ResultSetT struct { // #10.3
	SMS       [][]SMSData              `json:"sms"`        // #11.2
	MMS       [][]MMSData              `json:"mms"`        // #11.3
	VoiceCall []VoiceCallData          `json:"voice_call"` // #11.4
	Email     map[string][][]EmailData `json:"email"`      // #11.5
	Billing   BillingData              `json:"billing"`    // #11.6
	Support   []int                    `json:"support"`    // #11.7
	Incidents []IncidentData           `json:"incident"`   // #11.8
}

type ResultT struct { // #10.3
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}
