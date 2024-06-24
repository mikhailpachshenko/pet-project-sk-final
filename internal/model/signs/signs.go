package signs

func getAlphaTwoCountryList() [][]string {
	return [][]string{
		{"RU", "Russian Federation"},
		{"US", "United States of America"},
		{"GB", "Unitedd Kingdom of Great Britain and Northen Ireland"},
		{"FR", "France"},
		{"BL", "Saint Barthelemy"},
		{"AT", "Austria"},
		{"BG", "Bulgaria"},
		{"DK", "Denmark"},
		{"CA", "Canada"},
		{"ES", "Spain"},
		{"CH", "Switzerland"},
		{"TR", "Turkiye"},
		{"PE", "Peru"},
		{"NZ", "New Zealand"},
		{"MC", "Monaco"},
	}
}

func getCountriesList() []string {
	return []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
}

func GetCountriesList() []string {
	return []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
}

func getProviderList() []string {
	return []string{"Topolo", "Rond", "Kildy"}
}

func getVoiceProviderList() []string {
	return []string{"TransparentCalls", "E-Voice", "JushPhone"}
}

func getEmailProviderList() []string {
	return []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX", "Protonmail", "Yandex", "Mail.ru"}
}

func GetSmsProviderByCountry(country string) string {
	smsProviderMap := map[string]string{
		"RU": "Topolo",
		"US": "Rond",
		"GB": "Topolo",
		"FR": "Topolo",
		"BL": "Kildy",
		"AT": "Topolo",
		"BG": "Rond",
		"DK": "Topolo",
		"CA": "Rond",
		"ES": "Topolo",
		"CH": "Topolo",
		"TR": "Rond",
		"PE": "Topolo",
		"NZ": "Kildy",
		"MC": "Kildy",
	}
	return smsProviderMap[country]
}

func FindCountry(country string) bool {
	var line = getCountriesList()

	for _, val := range line {
		if val == country {
			return true
		}
	}
	return false
}

func FindProvider(provider string) bool {
	var list = getProviderList()

	for _, val := range list {
		if val == provider {
			return true
		}
	}
	return false
}

func FindProviderVoice(provider string) bool {
	var line = getVoiceProviderList()
	for _, val := range line {
		if val == provider {
			return true
		}
	}
	return false
}

func FindProviderEmail(provider string) bool {
	var line = getEmailProviderList()
	for _, val := range line {
		if val == provider {
			return true
		}
	}
	return false
}

func CodeToCountry(code string) (fullName string) {
	country := getAlphaTwoCountryList()

	for _, val := range country {
		if val[0] == code {
			fullName = val[1]
		}
	}
	return fullName
}
