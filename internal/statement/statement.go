package statement

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	company "personal-budget/internal/company"
	"personal-budget/internal/statement/banks"
	"personal-budget/internal/statement/domain"
	"strings"
)

func Generate(repository domain.MovimentRepository, companyService *company.Service) {
	outFile, err := os.Create("./export/statements/extratos_consolidados.csv")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	writer.Comma = ';'
	defer writer.Flush()

	writer.Write([]string{
		"Data", "Mês", "Descrição", "Categoria", "Banco", "Forma de pagamento", "Valor", "Tags",
	})

	err = filepath.Walk("./import/statements", func(path string, info os.FileInfo, err error) error {
		moviments := []domain.Moviment{}

		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".csv") {
			return nil
		}

		bank := filepath.Base(filepath.Dir(path))
		moviments, err = readFile(bank, path, companyService)

		for _, m := range moviments {
			writer.Write(m.ToArray())
			repository.Create(&m)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Extratos exportadas")
}

func readFile(bank string, path string, companyService *company.Service) ([]domain.Moviment, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	moviments := []domain.Moviment{}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.Comma = ','
	lines, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var m domain.Moviment

	for i, row := range lines {
		if i == 0 {
			continue
		}

		switch bank {
		case "picpay":
			m = banks.Picpay{CompanyService: companyService}.Import(row)
		case "c6":
			m = banks.C6{CompanyService: companyService}.Import(row)
		}

		moviments = append(moviments, m)
	}

	return moviments, nil
}
