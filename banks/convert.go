package banks

import (
	"personal-budget/config"
	"strings"
)

func fromTo(category string) string {
	text, ok := config.Categories[category]

	if ok {
		return text
	}

	return category
}

func ConvertCompany(company string) string {
	for key, value := range config.Companies {
		content := strings.Contains(strings.ToUpper(company), strings.ToUpper(key))
		if content {
			return value
		}
	}

	return "-"
}

func ConvertCategory(category string, description string) string {
	cat := ConvertCompany(description)
	if cat == "-" {
		cat = fromTo(category)
	}
	return cat
}
