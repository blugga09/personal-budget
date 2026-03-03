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

func (s *Service) SearchCategory(company string) *domain.Company {
	for _, c := range s.companies {
		content := strings.Contains(strings.ToUpper(company), strings.ToUpper(c.Name))
		if content {
			return c
		}
	}

	return nil
}
