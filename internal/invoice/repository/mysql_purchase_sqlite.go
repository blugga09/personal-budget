package repository

import (
	"database/sql"
	"log"
	"personal-budget/internal/invoice/domain"
)

type sqlitePurchaseRepository struct {
	db *sql.DB
}

func NewSqlitePurchaseRepository(db *sql.DB) *sqlitePurchaseRepository {
	return &sqlitePurchaseRepository{
		db: db,
	}
}

func (repository *sqlitePurchaseRepository) Create(purchase *domain.Purchase) error {
	stmt, err := repository.db.Prepare(`
	insert into purchases (date, month, description, current_installment, total_installment, bank, number, category, value, tags, content)
	values (?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = stmt.Exec(
		purchase.Date,
		purchase.Month,
		purchase.Description,
		purchase.CurrentInstallment,
		purchase.TotalInstallment,
		purchase.Bank,
		purchase.Number,
		purchase.Category,
		purchase.Value,
		purchase.Tags,
		purchase.Content,
	)

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (repository *sqlitePurchaseRepository) Find(description string) (*domain.Purchase, error) {
	stmt, err := repository.db.Prepare(`select * from purchases where description like "%?%"`)
	if err != nil {
		return nil, err
	}

	var purchase domain.Purchase
	err = stmt.QueryRow(description).Scan(&purchase.ID, &purchase.Date, &purchase.Month, &purchase.Description, &purchase.CurrentInstallment, &purchase.TotalInstallment, &purchase.Bank, &purchase.Number, &purchase.Category, &purchase.Value, &purchase.Tags, &purchase.Content)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &purchase, nil
	}
}
