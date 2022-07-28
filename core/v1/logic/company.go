package logic

import (
	"errors"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

type companyService struct {
	repo   repository.CompanyRepository
	client service.HttpClient
}

func (c companyService) Store(company v1.Company) error {
	if company.Id == "" {
		return errors.New("[ERROR]: No company id is given")
	}
	data := c.GetByName(company.Name, v1.StatusQueryOption{Option: enums.ACTIVE})
	if data.Id != "" {
		return errors.New("[ERROR]: Company name already exists")
	}
	return c.repo.Store(company)
}

func (c companyService) Delete(companyId string) error {
	err := c.repo.Delete(companyId)
	if err != nil {
		return err
	}
	return nil
}

func (c companyService) GetCompanies(option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Company {
	companies, _ := c.repo.GetCompanies(option, status)
	return companies
}

func (c companyService) GetByCompanyId(id string) v1.Company {
	return c.repo.GetByCompanyId(id)
}

func (c companyService) GetByName(name string, status v1.StatusQueryOption) v1.Company {
	return c.repo.GetByName(name, status)
}

// NewCompanyService returns Company type service
func NewCompanyService(repo repository.CompanyRepository, client service.HttpClient) service.Company {
	return &companyService{
		repo:   repo,
		client: client,
	}
}
