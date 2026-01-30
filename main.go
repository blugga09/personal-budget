package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	c6 "personal-budget/banks/C6"
	"personal-budget/domain"
	"strings"
)

func main() {
	outFile, err := os.Create("faturas_consolidadas.csv")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	writer.Comma = ';'
	defer writer.Flush()

	writer.Write([]string{
		"Data compra", "Mês fatura", "Descrição", "Parcela", "Total",
		"Banco", "Número", "Categoria", "Valor", "Tags",
	})

	err = filepath.Walk("invoices", func(path string, info os.FileInfo, err error) error {
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
			purchases, err = c6.ReadFile(bank, path, info)
		}

		for _, p := range purchases {

			writer.Write(p.ToArray())
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo consolidado gerado com sucesso:")
}
