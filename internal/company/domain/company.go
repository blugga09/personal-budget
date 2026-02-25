package domain

type Company struct {
	ID       int64
	Name     string
	Category string
	Tags     string
}

type CompanyRepository interface {
	Create(company *Company) error
	All() ([]*Company, error)
}
