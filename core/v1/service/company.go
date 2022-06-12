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
	UpdateApplication(companyId string, repositoryId string, app v1.Application) error
	Delete(companyId string) error
	GetCompanies(option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Company
	GetByCompanyId(id string, option v1.CompanyQueryOption) (v1.Company, int64)
	GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64)
	GetRepositoryByRepositoryId(id, repoId string, option v1.CompanyQueryOption) v1.Repository
	GetAllApplications(companyId string, option v1.CompanyQueryOption) ([]v1.Application, int64)
	GetApplicationsByRepositoryId(repositoryId string, companyId string, option v1.CompanyQueryOption, status v1.StatusQueryOption) ([]v1.Application, int64)
	GetCompanyByApplicationUrl(url string) v1.Company
	GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Application
	GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository
	GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application
	GetApplicationByApplicationId(companyId string, repoId string, applicationId string) v1.Application
	GetWebhookCount(companyId string) v1.ApplicationWebhookCount
	AppendRepositories(companyId string, repositories []v1.Repository) error
	SoftDeleteRepositories(companyId string, company v1.Company, repositories []v1.Repository) error
	DeleteRepositories(companyId string, company v1.Company, repositories []v1.Repository) error
	AppendApplications(companyId, repositoryId string, apps []v1.Application, option v1.CompanyQueryOption) error
	SoftDeleteApplications(companyId, repositoryId string, company v1.Company, apps []v1.Application) error
	DeleteApplications(companyId, repositoryId string, company v1.Company, apps []v1.Application, option v1.CompanyQueryOption) error
	CreateGithubWebHookAndUpdateApplication(companyId string, repoId string, token string, app v1.Application)
	CreateBitbucketWebHookAndUpdateApplication(companyId string, repoId string, token string, app v1.Application)
	UpdateWebhook(companyId, repoId, url, webhookId, action string) error
	EnableBitbucketWebhookAndUpdateApplication(companyId, repoId, url, token string) error
	EnableGithubWebhookAndUpdateApplication(companyId, repoId, url, token string) error
	DisableBitbucketWebhookAndUpdateApplication(companyId, repoId, url, webhookId, token string) error
	DisableGithubWebhookAndUpdateApplication(companyId, repoId, url, webhookId, token string) error
}
