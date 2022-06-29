package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
)

type applicationMetadataService struct {
	repo repository.ApplicationMetadataRepository
}

func (a applicationMetadataService) Store(applicationMetadataCollection v1.ApplicationMetadataCollection) error {
	return a.repo.Store(applicationMetadataCollection)
}

func (a applicationMetadataService) SearchAppsByCompanyIdAndName(companyId, name string) []v1.ApplicationMetadataCollection {
	return a.repo.SearchAppsByCompanyIdAndName(companyId, name)
}

func (a applicationMetadataService) GetById(id, companyId string) v1.ApplicationMetadataCollection {
	return a.repo.GetById(id, companyId)
}

func (a applicationMetadataService) Update(companyId string, data v1.ApplicationMetadataCollection) error {
	return a.repo.Update(companyId, data)
}

func (a applicationMetadataService) Delete(id, companyId string) error {
	return a.repo.Delete(id, companyId)
}

// NewApplicationMetadataService returns Application Metadata type service
func NewApplicationMetadataService(repo repository.ApplicationMetadataRepository) service.ApplicationMetadataService {
	return &applicationMetadataService{
		repo: repo,
	}
}
