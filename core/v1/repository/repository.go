package repository

import v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"

// RepositoryRepository company repository repository related operations
type RepositoryRepository interface {
	Store(repositories []v1.Repository) error
	GetByCompanyId(companyId string) []v1.Repository
}
