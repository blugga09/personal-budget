package banks

import (
	"fmt"
	"log"
	company "personal-budget/internal/company"
	"personal-budget/internal/statement/domain"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type C6 struct {
	CompanyService *company.Service
}

func (c C6) Import(row []string) domain.Moviment {

	printer := message.NewPrinter(language.BrazilianPortuguese)

	date, month := c.formatDate(row[0])
	description, category, tags := c.extractInfo(strings.TrimSpace(row[2]))

	return domain.Moviment{
		Date:        date,
		Month:       month,
		Description: description,
		Category:    category,
		Bank:        "c6",
		Method:      c.method(row[3]),
		Value:       c.formatMoney(row[4], row[5], printer),
		Tags:        tags,
	}
}

func (c C6) formatDate(value string) (string, string) {
	if len(value) < 7 {
		return "indefinido", "indefinido"
	}

	date := strings.Split(value, "/")
	return value + " 12:00", fmt.Sprintf("01/%s/%s", date[1], date[2])
}

func (c C6) formatMoney(credit string, debit string, printer *message.Printer) string {
	cf, err := strconv.ParseFloat(strings.TrimSpace(credit), 64)
	if err != nil {
		log.Println("Erro ao converter")
	}

	df, err := strconv.ParseFloat(strings.TrimSpace(debit), 64)
	if err != nil {
		log.Println("Erro ao converter")
	}

	value := func() float64 {
		if cf > 0 {
			return cf
		}
		return -df
	}()

	return printer.Sprintf("%.2f", value)
}

func (c C6) method(value string) string {
	value = strings.TrimSpace(value)

	if strings.Contains(value, "Pix") {
		return "PIX"
	}

	if strings.Contains(value, "TED") {
		return "TED"
	}

	return "Conta corrente"
}

func (c C6) extractInfo(description string) (string, string, string) {
	comp := c.CompanyService.SearchCategory(description)
	if comp != nil {
		return description, comp.Category, comp.Tags
	}

	return description, description, "-"
}
