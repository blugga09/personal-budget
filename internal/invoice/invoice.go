package invoice

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	company "personal-budget/internal/company"
	"personal-budget/internal/invoice/banks"
	"personal-budget/internal/invoice/domain"
	"strings"
)

func Generate(repository domain.PurchaseRepository, companyService *company.Service) {
	outFile, err := os.Create("./export/invoices/faturas_consolidadas.csv")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	writer.Comma = ';'
	defer writer.Flush()

	writer.Write([]string{
		"Data compra", "Mês fatura", "Descrição", "Parcela", "Total", "Banco", "Número", "Categoria", "Valor", "Tags",
	})

	err = filepath.Walk("./import/invoices", func(path string, info os.FileInfo, err error) error {
		purchases := []domain.Purchase{}

		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasPrefix(info.Name(), "Fatura_") || !strings.HasSuffix(info.Name(), ".csv") {
			return nil
		}

		bank := filepath.Base(filepath.Dir(path))
		date := info.Name()

		switch bank {
		case "c6":
			purchases, err = banks.C6{CompanyService: companyService}.Import(path, date)
		}

		for _, p := range purchases {
			writer.Write(p.ToArray())
			repository.Create(&p)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Faturas exportadas")
}
