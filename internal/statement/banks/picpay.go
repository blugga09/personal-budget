package banks

import (
	"fmt"
	company "personal-budget/internal/company"
	"personal-budget/internal/helper"
	"personal-budget/internal/statement/domain"
	"strings"
)

type Picpay struct {
	CompanyService *company.Service
}

const BANK = "picpay"

func (p Picpay) Import(row []string) domain.Moviment {

	date, month := p.formatDate(row[0])
	description, category, tags := p.extractInfo(strings.TrimSpace(row[2]), strings.TrimSpace(row[3]))

	return domain.Moviment{
		Date:        fmt.Sprintf("%s %s", date, strings.TrimSpace(row[1])),
		Month:       month,
		Description: description,
		Category:    category,
		Bank:        BANK,
		Method:      p.method(row[2]),
		Value:       p.formatMoney(row[4]),
		Tags:        tags,
	}
}

func (p Picpay) formatDate(value string) (string, string) {
	if len(value) < 7 {
		return "indefinido", "indefinido"
	}

	date := strings.Split(value, "-")
	return fmt.Sprintf("%s/%s/%s", date[2], date[1], date[0]), fmt.Sprintf("01/%s/%s", date[1], date[0])
}

func (p Picpay) formatMoney(value string) string {
	replacer := strings.NewReplacer("+R$ ", "", "−R$ ", "-")
	return replacer.Replace(strings.TrimSpace(value))
}

func (p Picpay) method(value string) string {
	value = strings.TrimSpace(value)

	if strings.Contains(value, "Pix") {
		return "PIX"
	}

	if strings.Contains(value, "TED") {
		return "TED"
	}

	return "Conta corrente"
}

func (p Picpay) extractInfo(category string, description string) (string, string, string) {
	comp := p.CompanyService.SearchCategory(description)
	if comp != nil {
		return comp.Name, comp.Category, comp.Tags
	}

	category = helper.ConvertCategory(strings.TrimSpace(strings.Split(category, "/")[0]), description)
	return description, category, strings.Join(strings.Split(category, " / "), ",")
}
