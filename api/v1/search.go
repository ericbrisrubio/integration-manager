package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/labstack/echo/v4"
)

type searchApi struct {
	searchService	service.Search
}

func (s searchApi) SearchReposAndAppsByCompanyIdAndName(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	name := context.QueryParam("name")
	reposOption := false
	if context.QueryParam("reposOption") == "true" {
		reposOption = true
	}
	appsOption := false
	if context.QueryParam("appsOption") == "true" {
		appsOption = true
	}
	data := s.searchService.SearchReposAndAppsByCompanyIdAndName(companyId, name, reposOption, appsOption)
	if len(data.Repositories) == 0 && len(data.Applications) == 0 {
		return common.GenerateErrorResponse(context, nil, "No data found!")
	}
	return common.GenerateSuccessResponse(context, data,
		nil, "Operation Successful")
}

// NewSearchApi returns search type api
func NewSearchApi(searchService	service.Search) api.Search {
	return &searchApi{
		searchService: searchService,
	}
}
