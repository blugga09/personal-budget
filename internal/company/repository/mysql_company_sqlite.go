package repository

import (
	"database/sql"
	"errors"
	"personal-budget/internal/company/domain"
)

type sqliteCompanyRepository struct {
	db *sql.DB
}

func NewSqliteCompanyRepository(db *sql.DB) *sqliteCompanyRepository {
	return &sqliteCompanyRepository{
		db: db,
	}
}

func (repository *sqliteCompanyRepository) Create(company *domain.Company) error {
	stmt, err := repository.db.Prepare(`insert into companies (description, category, tags) values (?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		company.Name,
		company.Category,
		company.Tags,
	)

	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (repository *sqliteCompanyRepository) All() ([]*domain.Company, error) {
	stmt, err := repository.db.Prepare(`select name, category, tags from companies`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	var companies []*domain.Company
	for rows.Next() {
		var company domain.Company
		err = rows.Scan(&company.Name, &company.Category, &company.Tags)
		if err != nil {
			return nil, err
		}

		companies = append(companies, &company)
	}

	if len(companies) == 0 {
		return nil, errors.New("Not Found")
	}

	return companies, nil
}
