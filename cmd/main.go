package main

import (
	"database/sql"
	"fmt"
	"log"
	enterprise_repository "personal-budget/internal/enterprise/repository"
	"personal-budget/internal/invoice"
	purchase_repository "personal-budget/internal/invoice/repository"
	"personal-budget/internal/statement"
	moviment_repository "personal-budget/internal/statement/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	purchaseRepository := purchase_repository.NewSqlitePurchaseRepository(db)
	enterpriseRepository := enterprise_repository.NewSqliteEnterpriseRepository(db)
	movimentRepository := moviment_repository.NewSqliteMovimentRepository(db)
	enterprises, err := enterpriseRepository.All()
	if err != nil {
		log.Fatal(err)
	}

	invoice.Generate(purchaseRepository, enterprises)
	statement.Generate(movimentRepository, enterprises)

	fmt.Println("Arquivo consolidado gerado com sucesso:")
}
