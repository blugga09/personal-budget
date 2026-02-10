package domain

type Purchase struct {
	Date               string
	Month              string
	Description        string
	CurrentInstallment string
	TotalInstallment   string
	Bank               string
	Number             string
	Category           string
	Value              string
	Tags               string
}

func (p Purchase) ToArray() []string {
	return []string{
		p.Date,
		p.Month,
		p.Description,
		p.CurrentInstallment,
		p.TotalInstallment,
		p.Bank,
		p.Number,
		p.Category,
		p.Value,
		p.Tags,
	}
}
