package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
)

var data []v1.Company

func InitCompanyData() []v1.Company {
	var data []v1.Company
	//add object to data
	return data
}
func NewMockCompanyRepository() repository.CompanyRepository {
	manager := GetMockDmManager()
	manager.Db.Drop(context.Background())
	return &companyRepository{
		manager: GetMockDmManager(),
		timeout: 3000,
	}

}
