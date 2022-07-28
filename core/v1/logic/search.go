package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
)

type searchService struct {
	applicationService service.Application
}

func (s searchService) SearchReposAndAppsByCompanyIdAndName(companyId, name string, reposOption, appsOption bool) v1.SearchData {
	var searchData v1.SearchData
	if appsOption {
		searchData.Applications = s.applicationService.SearchAppsByCompanyIdAndName(companyId, name)
	}
	return searchData
}

// NewSearchService returns Search type service
func NewSearchService(applicationService service.Application) service.Search {
	return &searchService{
		applicationService: applicationService,
	}
}
