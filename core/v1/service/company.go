package service

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

// Company Company related operations.
type Company interface {
	Store(company v1.Company) error
	UpdateRepositories(companyId string, repositories []v1.Repository, companyUpdateOption v1.RepositoryUpdateOption) error
	UpdateApplications(companyId string, repositoryId string, apps []v1.Application, companyUpdateOption v1.ApplicationUpdateOption) error
	Delete(companyId string) error
	GetCompanies(option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Company
	GetByCompanyId(id string, option v1.CompanyQueryOption) (v1.Company, int64)
	GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64)
	GetRepositoryByRepositoryId(id string, companyId string, option v1.CompanyQueryOption) v1.Repository
	GetApplicationsByRepositoryId(repositoryId string, companyId string, option v1.CompanyQueryOption, status v1.StatusQueryOption) ([]v1.Application, int64)
	GetCompanyByApplicationUrl(url string) v1.Company
	GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Application
	GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository
	GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application
	GetApplicationByApplicationId(companyId string, repoId string, applicationId string) v1.Application
}
