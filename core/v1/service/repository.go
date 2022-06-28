package service

import v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"

// Repository related operations.
type Repository interface {
	Store(repositories []v1.Repository) error
	GetByCompanyId(companyId string) []v1.Repository
}
