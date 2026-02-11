package banks

import (
	"fmt"
	"personal-budget/internal/helper"
	"personal-budget/internal/statement/domain"
	"strings"
)

func ImportPicpay(row []string) domain.Moviment {
	rawCategory := strings.TrimSpace(row[2])
	category := strings.TrimSpace(strings.Split(rawCategory, "/")[0])
	description := strings.TrimSpace(row[3])
	date, month := formatDate(row[0])

	return domain.Moviment{
		Date:        fmt.Sprintf("%s %s", date, strings.TrimSpace(row[1])),
		Month:       month,
		Description: description,
		Category:    helper.ConvertCategory(category, description),
		Bank:        "Picpay",
		Method:      method(row[2]),
		Value:       formatMoney(row[4]),
		Tags:        strings.Join(strings.Split(rawCategory, " / "), ","),
	}
}

func formatDate(value string) (string, string) {
	if len(value) < 7 {
		return "indefinido", "indefinido"
	}

	date := strings.Split(value, "-")
	return fmt.Sprintf("%s/%s/%s", date[2], date[1], date[0]), fmt.Sprintf("01/%s/%s", date[1], date[0])
}

func formatMoney(value string) string {
	replacer := strings.NewReplacer("+R$ ", "", "−R$ ", "-")
	return replacer.Replace(strings.TrimSpace(value))
}

func method(value string) string {
	value = strings.TrimSpace(value)

	if strings.Contains(value, "Pix") {
		return "PIX"
	}

	if strings.Contains(value, "TED") {
		return "TED"
	}

	return "Conta corrente"
}
