package repository

import v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"

// RepositoryRepository company repository repository related operations
type RepositoryRepository interface {
	Store(repositories []v1.Repository) error
	GetByCompanyId(companyId string, pagination bool, option v1.CompanyQueryOption) ([]v1.Repository, int64)
	DeleteRepository(companyId, repositoryId string) error
	GetById(companyId, repositoryId string) v1.Repository
}
