package mongo

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/enums"
	"time"
)

var (
	CompanyCollection = "CompanyCollection"
)

type companyRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (c companyRepository) GetCompanyByApplicationUrl(url string) v1.Company {
	panic("implement me")
}

func (c companyRepository) GetCompanies(option v1.CompanyQueryOption) ([]v1.Company, int64) {
	panic("implement me")
}

func (c companyRepository) GetByCompanyId(id string, option v1.CompanyQueryOption) v1.Company {
	panic("implement me")
}

func (c companyRepository) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) []v1.Repository {
	panic("implement me")
}

func (c companyRepository) GetApplicationsByCompanyId(id string, option v1.CompanyQueryOption) []v1.Application {
	panic("implement me")
}

func (c companyRepository) GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption) []v1.Application {
	panic("implement me")
}

func (c companyRepository) Store(company v1.Company) (error, int64) {
	panic("implement me")
}

func (c companyRepository) Update(company v1.Company, companyUpdateOption ...v1.CompanyUpdateOption) {
	panic("implement me")
}

func (c companyRepository) Delete(companyId string) error {
	panic("implement me")
}

func NewCompanyRepository(timeout int) repository.CompanyRepository {
	return &companyRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}

}
