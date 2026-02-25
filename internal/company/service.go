package company

import (
	"log"
	"personal-budget/internal/company/domain"
	"strings"
)

type Service struct {
	companies []*domain.Company
}

func NewService(repository domain.CompanyRepository) *Service {
	companies, err := repository.All()
	if err != nil {
		log.Fatal(err)
	}
	return &Service{companies}
}

func (s *Service) searchCategory(company string) string {
	for _, c := range s.companies {
		content := strings.Contains(strings.ToUpper(c.Name), strings.ToUpper(company))
		if content {
			return c.Category
		}
	}

	return ""
}
