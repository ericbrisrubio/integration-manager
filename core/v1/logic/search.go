package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
)

type searchService struct {
	applicationMetadataRepository	repository.ApplicationMetadataRepository
}

func (s searchService) SearchReposAndAppsByCompanyIdAndName(companyId, name string, reposOption, appsOption bool) v1.SearchData {
	var searchData v1.SearchData
	if appsOption {
		searchData.Applications = s.applicationMetadataRepository.SearchAppsByCompanyIdAndName(companyId, name)
	}
	return searchData
}

// NewSearchService returns Search type service
func NewSearchService(applicationMetadataRepository	repository.ApplicationMetadataRepository) service.Search {
	return &searchService{
		applicationMetadataRepository: applicationMetadataRepository,
	}
}