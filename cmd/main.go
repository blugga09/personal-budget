package main

import (
	"fmt"
	"personal-budget/internal/invoice"
	"personal-budget/internal/statement"
)

func main() {

	invoice.Generate()
	statement.Generate()

	fmt.Println("Arquivo consolidado gerado com sucesso:")
}
