package main

import (
	"database/sql"
	"fmt"
	"log"
	enterprise_repository "personal-budget/internal/enterprise/repository"
	"personal-budget/internal/invoice"
	invoice_repository "personal-budget/internal/invoice/repository"
	"personal-budget/internal/statement"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	purchaseRepository := invoice_repository.NewSqlitePurchaseRepository(db)
	enterpriseRepository := enterprise_repository.NewSqliteEnterpriseRepository(db)
	enterprises, err := enterpriseRepository.All()
	if err != nil {
		log.Fatal(err)
	}

	invoice.Generate(purchaseRepository, enterprises)
	statement.Generate()

	fmt.Println("Arquivo consolidado gerado com sucesso:")
}
