package domain

type Moviment struct {
	Date        string
	Month       string
	Description string
	Category    string
	Bank        string
	Method      string
	Value       string
	Tags        string
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
	}
}
