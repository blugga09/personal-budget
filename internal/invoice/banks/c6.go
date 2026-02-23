package banks

import (
	"encoding/csv"
	"fmt"
	"os"
	"personal-budget/internal/helper"
	"personal-budget/internal/invoice/domain"
	"strconv"
	"strings"
)

func ImportC6(bank string, path string, info os.FileInfo) ([]domain.Purchase, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	month := extractMonth(info.Name())

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

		rawCategory := strings.TrimSpace(row[3])
		category := strings.TrimSpace(strings.Split(rawCategory, "/")[0])
		description := strings.TrimSpace(row[4])
		currentInstallment, totalInstallment := extractInstallment(strings.TrimSpace(row[5]))

		purchases = append(purchases, domain.Purchase{
			Date:               strings.TrimSpace(row[0]),
			Month:              month,
			Description:        description,
			CurrentInstallment: currentInstallment,
			TotalInstallment:   totalInstallment,
			Bank:               bank,
			Number:             strings.TrimSpace(row[2]),
			Category:           helper.ConvertCategory(category, description),
			Value:              strings.ReplaceAll(strings.TrimSpace(row[8]), ".", ","),
			Tags:               strings.Join(strings.Split(rawCategory, " / "), ","),
			Content:            strings.Join(row, ";"),
		})
	}

	return purchases, nil
}

func formatMonth(date string) string {
	if len(date) < 7 {
		return "indefinido"
	}
	year := date[:4]
	month := date[5:7]
	return fmt.Sprintf("01/%s/%s", month, year)
}

func extractMonth(fileName string) string {
	name := strings.TrimSuffix(fileName, ".csv")
	part := strings.Split(name, "_")
	if len(part) < 2 {
		return ""
	}
	date := part[1]
	return formatMonth(date)
}

func extractInstallment(installment string) (string, string) {
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
