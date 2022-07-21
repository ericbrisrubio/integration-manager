package service

import v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"

// ApplicationMetadataService application metadata service related operations
type ApplicationMetadataService interface {
	Store(applicationMetadataCollection v1.ApplicationMetadataCollection) error
	SearchAppsByCompanyIdAndName(companyId, name string) []v1.ApplicationMetadataCollection
	GetById(id, companyId string) v1.ApplicationMetadataCollection
	Update(companyId string, data v1.ApplicationMetadataCollection) error
	Delete(id, companyId string) error
}
