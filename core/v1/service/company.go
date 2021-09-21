package service

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/enums"
)

type Company interface {
	Store(company v1.Company) error
	Update(company v1.Company, companyUpdateOption ...v1.CompanyUpdateOption)
	Delete(companyId string) error
	GetCompanies(option v1.CompanyQueryOption) []v1.Company
	GetByCompanyId(id string, option v1.CompanyQueryOption) v1.Company
	GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) []v1.Repository
	GetApplicationsByCompanyId(id string, option v1.CompanyQueryOption) []v1.Application
	GetCompanyByApplicationUrl(url string) v1.Company
	GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption) []v1.Application
	GetRepositoryByCompanyIdAndApplicationUrl(id,url string)v1.Repository
}
