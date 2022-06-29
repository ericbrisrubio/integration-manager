package service

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

// Application Application related operations.
type Application interface {
	StoreAll(applications []v1.Application) error
	CreateWebHookAndUpdateApplications(repoType enums.REPOSITORY_TYPE, token string, apps []v1.Application)
	GetByCompanyIdAndRepoId(companyId, repoId string, pagination bool, option v1.CompanyQueryOption, statusQuery bool, status v1.StatusQueryOption) ([]v1.Application, int64)
	GetApplicationsByCompanyIdAndRepositoryType(companyId string, _type enums.REPOSITORY_TYPE, pagination bool, option v1.CompanyQueryOption, statusQuery bool, status v1.StatusQueryOption) ([]v1.Application, int64)
	SoftDeleteApplication(application v1.Application) error
	DeleteApplication(companyId, repositoryId, applicationId string) error
	UpdateWebhook(repository v1.Repository, url, webhookId, action string) error
}
