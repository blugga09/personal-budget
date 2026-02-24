package domain

type Moviment struct {
	ID          int64
	Date        string
	Month       string
	Description string
	Category    string
	Bank        string
	Method      string
	Value       string
	Tags        string
	Content     string
}

func (m Moviment) ToArray() []string {
	return []string{
		m.Date,
		m.Month,
		m.Description,
		m.Category,
		m.Bank,
		m.Method,
		m.Value,
		m.Tags,
		m.Content,
	}
}

type MovimentRepository interface {
	Create(purchase *Moviment) error
	Find(description string) (*Moviment, error)
}
