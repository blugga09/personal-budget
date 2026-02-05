package main

import (
	"fmt"
	"personal-budget/internal/invoice"
)

func main() {

	invoice.Generate()

	fmt.Println("Arquivo consolidado gerado com sucesso:")
}
