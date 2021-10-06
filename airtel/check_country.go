package airtel

import "strings"

func contains(s []string, str string) bool {
	for _, v := range s {
		if strings.ToLower(v) == strings.ToLower(str) {
			return true
		}
	}

	return false
}

func CheckCountry(api string, country string, allCountries map[string][]string) bool {
	a := allCountries[api]
	if a == nil || len(a) <= 0 {
		return false
	}

	return contains(a, country)
}
