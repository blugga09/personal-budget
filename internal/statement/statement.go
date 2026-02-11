package statement

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"personal-budget/internal/statement/banks"
	"personal-budget/internal/statement/domain"
	"strings"
)

func Generate() {
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

		moviments, err = readFile(bank, path)

		for _, m := range moviments {
			writer.Write(m.ToArray())
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Extratos exportadas")
}

func readFile(bank string, path string) ([]domain.Moviment, error) {
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

	for i, row := range lines {
		if i == 0 {
			continue
		}

		switch bank {
		case "picpay":
			moviments = append(moviments, banks.ImportPicpay(row))
		}
	}

	return moviments, nil
}
