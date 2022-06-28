package repository

import v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"

// ApplicationRepository company application repository related operations
type ApplicationRepository interface {
	Store(application v1.Application) error
	StoreAll(applications []v1.Application) error
	GetByCompanyIdAndRepoId(companyId, repoId string) []v1.Application
}
