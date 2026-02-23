package repository

import (
	"database/sql"
	"errors"
	"personal-budget/internal/enterprise/domain"
)

type sqliteEnterpriseRepository struct {
	db *sql.DB
}

func NewSqliteEnterpriseRepository(db *sql.DB) *sqliteEnterpriseRepository {
	return &sqliteEnterpriseRepository{
		db: db,
	}
}

func (repository *sqliteEnterpriseRepository) Create(enterprise *domain.Enterprise) error {
	stmt, err := repository.db.Prepare(`insert into enterprises (description, category, tags) values (?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		enterprise.Name,
		enterprise.Category,
		enterprise.Tags,
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

func (repository *sqliteEnterpriseRepository) All() ([]*domain.Enterprise, error) {
	stmt, err := repository.db.Prepare(`select name, category, tags from enterprises`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	var enterprises []*domain.Enterprise
	for rows.Next() {
		var enterprise domain.Enterprise
		err = rows.Scan(&enterprise.Name, &enterprise.Category, &enterprise.Tags)
		if err != nil {
			return nil, err
		}

		enterprises = append(enterprises, &enterprise)
	}

	if len(enterprises) == 0 {
		return nil, errors.New("Not Found")
	}

	return enterprises, nil
}
