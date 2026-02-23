package invoice

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	enterprise "personal-budget/internal/enterprise/domain"
	"personal-budget/internal/invoice/banks"
	"personal-budget/internal/invoice/domain"
	"strings"
)

func Generate(repository domain.PurchaseRepository, enterprises []*enterprise.Enterprise) {
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

		switch bank {
		case "c6":
			purchases, err = banks.ImportC6(bank, path, info)
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
