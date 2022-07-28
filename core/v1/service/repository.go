package service

import v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"

// Repository related operations.
type Repository interface {
	Store(repositories []v1.Repository) error
	GetByCompanyId(companyId string, pagination bool, option v1.CompanyQueryOption) ([]v1.Repository, int64)
	UpdateRepositories(companyId string, repositoriesDto []v1.RepositoryDto, repositoryUpdateOption v1.RepositoryUpdateOption) error
	GetById(companyId, repositoryId string) v1.Repository
	GetByCompanyIdAndApplicationUrl(companyId, url string) v1.Repository
	SearchByNameAndCompanyId(name, companyId string) []v1.Repository
}
