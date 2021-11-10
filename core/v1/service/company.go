package service

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/enums"
)

type Company interface {
	Store(company v1.Company) error
	UpdateRepositories(companyId string, repositories []v1.Repository, companyUpdateOption v1.CompanyUpdateOption) error
	UpdateApplications(companyId string, repositoryId string, apps []v1.Application, companyUpdateOption v1.CompanyUpdateOption) error
	Delete(companyId string) error
	GetCompanies(option v1.CompanyQueryOption) []v1.Company
	GetByCompanyId(id string, option v1.CompanyQueryOption) (v1.Company, int64)
	GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64)
	GetRepositoryByRepositoryId(id string) v1.Repository
	GetApplicationsByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Application, int64)
	GetCompanyByApplicationUrl(url string) v1.Company
	GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption) []v1.Application
	GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository
	GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application
}
