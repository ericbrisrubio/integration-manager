package service

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

// Application Application related operations.
type Application interface {
	GetAll(companyId string, option v1.CompanyQueryOption) ([]v1.Application, int64)
	GetByApplicationId(companyId string, repoId string, applicationId string) v1.Application
	StoreAll(applications []v1.Application) error
	CreateWebHookAndUpdateApplications(repoType enums.REPOSITORY_TYPE, token string, apps []v1.Application)
	GetByCompanyIdAndRepoId(companyId, repoId string, pagination bool, option v1.CompanyQueryOption, statusQuery bool, status v1.StatusQueryOption) ([]v1.Application, int64)
	GetApplicationsByCompanyIdAndRepositoryType(companyId string, _type enums.REPOSITORY_TYPE, pagination bool, option v1.CompanyQueryOption, statusQuery bool, status v1.StatusQueryOption) ([]v1.Application, int64)
	GetByCompanyIdAndRepositoryIdAndUrl(companyId, repositoryId, applicationUrl string) v1.Application
	GetByCompanyIdAndUrl(companyId, applicationUrl string) v1.Application
	SoftDeleteApplication(application v1.Application) error
	DeleteApplication(companyId, repositoryId, applicationId string) error
	UpdateApplications(repository v1.Repository, apps []v1.Application, applicationUpdateOption v1.ApplicationUpdateOption) error
	UpdateWebhook(repository v1.Repository, url, webhookId, action string) error
}
