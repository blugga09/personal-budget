package main

import (
	"database/sql"
	"fmt"
	"log"
	company "personal-budget/internal/company"
	company_repository "personal-budget/internal/company/repository"
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
	companyRepository := company_repository.NewSqliteCompanyRepository(db)
	movimentRepository := moviment_repository.NewSqliteMovimentRepository(db)
	companyService := company.NewService(companyRepository)
	if err != nil {
		log.Fatal(err)
	}

	invoice.Generate(purchaseRepository, companyService)
	statement.Generate(movimentRepository, companyService)

	fmt.Println("Arquivo consolidado gerado com sucesso:")
}
