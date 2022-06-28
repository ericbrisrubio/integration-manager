package service

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

// Application Application related operations.
type Application interface {
	StoreAll(applications []v1.Application) error
	CreateWebHookAndUpdateApplications(repoType enums.REPOSITORY_TYPE, token string, apps []v1.Application)
	GetByCompanyIdAndRepoId(companyId, repoId string) []v1.Application
}
