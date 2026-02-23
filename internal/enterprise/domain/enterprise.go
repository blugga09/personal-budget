package domain

type Enterprise struct {
	ID       int64
	Name     string
	Category string
	Tags     string
}

type EnterpriseRepository interface {
	Create(enterprise *Enterprise) error
	All() ([]*Enterprise, error)
}
