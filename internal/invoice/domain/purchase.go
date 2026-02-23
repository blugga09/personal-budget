package domain

type Purchase struct {
	ID                 int64
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
	Content            string
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
		p.Content,
	}
}

type PurchaseRepository interface {
	Create(purchase *Purchase) error
	Find(description string) (*Purchase, error)
}
