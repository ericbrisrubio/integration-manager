package inmemory

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
)

type companyRepository struct {
}

func (c companyRepository) Store(company v1.Company) error {
	//TODO implement me
	panic("implement me")
}

func (c companyRepository) Delete(companyId string) error {
	//TODO implement me
	panic("implement me")
}

func (c companyRepository) GetCompanies(option v1.CompanyQueryOption, status v1.StatusQueryOption) ([]v1.Company, int64) {
	//TODO implement me
	panic("implement me")
}

func (c companyRepository) GetByCompanyId(id string) v1.Company {
	//TODO implement me
	panic("implement me")
}

func (c companyRepository) GetByName(name string, status v1.StatusQueryOption) v1.Company {
	//TODO implement me
	panic("implement me")
}

// NewCompanyRepository returns CompanyRepository type object
func NewCompanyRepository(timeout int) repository.CompanyRepository {
	return &companyRepository{}

}
