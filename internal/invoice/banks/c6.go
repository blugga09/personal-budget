package banks

import (
	"encoding/csv"
	"fmt"
	"os"
	company "personal-budget/internal/company"
	"personal-budget/internal/helper"
	"personal-budget/internal/invoice/domain"
	"strconv"
	"strings"
)

type C6 struct {
	CompanyService *company.Service
}

const BANK = "c6"

func (c C6) Import(path string, date string) ([]domain.Purchase, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	month := c.extractMonth(date)

	purchases := []domain.Purchase{}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for i, row := range lines {
		if i == 0 || len(row) < 9 {
			continue
		}

		num, _ := strconv.ParseFloat(strings.TrimSpace(row[8]), 64)

		if num < 0 {
			continue
		}

		currentInstallment, totalInstallment := c.extractInstallment(strings.TrimSpace(row[5]))
		description, category, tags := c.extractInfo(strings.TrimSpace(row[3]), strings.TrimSpace(row[4]))

		purchases = append(purchases, domain.Purchase{
			Date:               strings.TrimSpace(row[0]),
			Month:              month,
			Description:        description,
			CurrentInstallment: currentInstallment,
			TotalInstallment:   totalInstallment,
			Bank:               BANK,
			Number:             strings.TrimSpace(row[2]),
			Category:           category,
			Value:              strings.ReplaceAll(strings.TrimSpace(row[8]), ".", ","),
			Tags:               tags,
			Content:            strings.Join(row, ";"),
		})
	}

	return purchases, nil
}

func (c C6) formatMonth(date string) string {
	if len(date) < 7 {
		return "indefinido"
	}
	year := date[:4]
	month := date[5:7]
	return fmt.Sprintf("01/%s/%s", month, year)
}

func (c C6) extractMonth(fileName string) string {
	name := strings.TrimSuffix(fileName, ".csv")
	part := strings.Split(name, "_")
	if len(part) < 2 {
		return ""
	}
	date := part[1]
	return c.formatMonth(date)
}

func (c C6) extractInstallment(installment string) (string, string) {
	currentInstallment := "1"
	totalInstallment := "1"
	if strings.Contains(installment, "/") {
		part := strings.Split(installment, "/")
		if len(part) == 2 {
			currentInstallment = strings.TrimSpace(part[0])
			totalInstallment = strings.TrimSpace(part[1])
		}
	} else if strings.Contains(strings.ToLower(installment), "única") {
		currentInstallment = "1"
		totalInstallment = "1"
	}
	return currentInstallment, totalInstallment
}

func (c C6) extractInfo(category string, description string) (string, string, string) {
	comp := c.CompanyService.SearchCategory(description)
	if comp != nil {
		return comp.Name, comp.Category, comp.Tags
	}

	category = helper.ConvertCategory(strings.TrimSpace(strings.Split(category, "/")[0]), description)
	return description, category, strings.Join(strings.Split(category, " / "), ",")
}
