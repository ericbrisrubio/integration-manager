package repository

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/enums"
)

type CompanyRepository interface {
	Store(company v1.Company) error
	AppendRepositories(companyId string, repos []v1.Repository) error
	DeleteRepositories(companyId string, repos []v1.Repository, isSoftDelete bool) error
	AppendApplications(companyId, repositoryId string, apps []v1.Application) error
	DeleteApplications(companyId, repositoryId string, apps []v1.Application, isSoftDelete bool) error
	Delete(companyId string) error
	GetCompanies(option v1.CompanyQueryOption) ([]v1.Company, int64)
	GetByCompanyId(id string, option v1.CompanyQueryOption) v1.Company
	GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) []v1.Repository
	GetApplicationsByCompanyId(id string, option v1.CompanyQueryOption) []v1.Application
	GetCompanyByApplicationUrl(url string) v1.Company
	GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption) []v1.Application
	GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository
	GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId,repositoryId,applicationUrl string)v1.Application
}
