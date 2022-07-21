package service

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
)

// Company Company related operations.
type Company interface {
	Store(company v1.Company) error
	Delete(companyId string) error
	GetCompanies(option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Company
	GetByCompanyId(id string) v1.Company
	GetByName(name string, status v1.StatusQueryOption) v1.Company
}
