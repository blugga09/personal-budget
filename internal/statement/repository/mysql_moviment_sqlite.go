package repository

import (
	"database/sql"
	"log"
	"personal-budget/internal/statement/domain"
)

type sqliteMovimentRepository struct {
	db *sql.DB
}

func NewSqliteMovimentRepository(db *sql.DB) *sqliteMovimentRepository {
	return &sqliteMovimentRepository{
		db: db,
	}
}

func (repository *sqliteMovimentRepository) Create(moviment *domain.Moviment) error {
	stmt, err := repository.db.Prepare(`
	insert into moviments (date, month, description, bank, category, value, tags, content)
	values (?,?,?,?,?,?,?,?)`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = stmt.Exec(
		moviment.Date,
		moviment.Month,
		moviment.Description,
		moviment.Bank,
		moviment.Category,
		moviment.Value,
		moviment.Tags,
		moviment.Content,
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

func (repository *sqliteMovimentRepository) Find(description string) (*domain.Moviment, error) {
	stmt, err := repository.db.Prepare(`select * from moviments where description like "%?%"`)
	if err != nil {
		return nil, err
	}

	var moviment domain.Moviment
	err = stmt.QueryRow(description).Scan(&moviment.ID, &moviment.Date, &moviment.Month, &moviment.Description, &moviment.Bank, &moviment.Category, &moviment.Value, &moviment.Tags, &moviment.Content)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &moviment, nil
	}
}
