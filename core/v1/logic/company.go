package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
)

type companyService struct {
	repo repository.CompanyRepository
}

func (c companyService) GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository {
	return c.repo.GetRepositoryByCompanyIdAndApplicationUrl(id,url)
}

func (c companyService) GetCompanyByApplicationUrl(url string) v1.Company {
	panic("implement me")
}

func (c companyService) Store(company v1.Company) error {
	panic("implement me")
}

func (c companyService) Update(company v1.Company, companyUpdateOption ...v1.CompanyUpdateOption) {
	panic("implement me")
}

func (c companyService) Delete(companyId string) error {
	panic("implement me")
}

func (c companyService) GetCompanies(option v1.CompanyQueryOption) []v1.Company {
	panic("implement me")
}

func (c companyService) GetByCompanyId(id string, option v1.CompanyQueryOption) v1.Company {
	panic("implement me")
}

func (c companyService) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) []v1.Repository {
	panic("implement me")
}

func (c companyService) GetApplicationsByCompanyId(id string, option v1.CompanyQueryOption) []v1.Application {
	panic("implement me")
}

func (c companyService) GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption) []v1.Application {
	panic("implement me")
}

func NewCompanyService(repo repository.CompanyRepository) service.Company {
	return &companyService{
		repo: repo,
	}
}
