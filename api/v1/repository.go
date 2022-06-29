package v1

import (
	"errors"
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type repositoryApi struct {
	repositoryService  service.Repository
	applicationService service.Application
	observerList       []service.Observer
}

// Get.. Get Repository by id
// @Summary Get Repository by id
// @Description Gets Repository by id
// @Tags Repository
// @Produce json
// @Param id path string true "Repository id"
// @Param companyId query string true "company id"
// @Success 200 {object} common.ResponseDTO{data=v1.RepositoryDto}
// @Router /api/v1/repositories/{id} [GET]
func (r repositoryApi) GetById(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return errors.New("company id is required")
	}
	repoId := context.Param("id")
	if repoId == "" {
		return errors.New("repository id is required")
	}
	option := getQueryOption(context)
	repository := r.repositoryService.GetById(companyId, repoId)
	repositoryDto := v1.RepositoryDto{
		Id:   repository.Id,
		Type: repository.Type,
	}
	if option.LoadToken {
		repositoryDto.Token = repository.Token
	}
	if option.LoadApplications {
		applications, _ := r.applicationService.GetByCompanyIdAndRepoId(companyId, repoId, false, v1.CompanyQueryOption{}, false, v1.StatusQueryOption{})
		var applicationsDto []v1.ApplicationDto
		for _, eachApp := range applications {
			applicationsDto = append(applicationsDto, v1.ApplicationDto{
				MetaData: eachApp.MetaData,
				Url:      eachApp.Url,
				Webhook:  eachApp.Webhook,
				Status:   eachApp.Status,
			})
		}
		repositoryDto.Applications = applicationsDto
	}
	return common.GenerateSuccessResponse(context, repositoryDto, nil, "Success!")
}

// Get.. Get Applications by repository id
// @Summary Get Applications by repository id
// @Description Gets Applications by repository id
// @Tags Repository
// @Produce json
// @Param id path string true "Repository id"
// @Param companyy query string true "Company id"
// @Success 200 {object} common.ResponseDTO{data=[]v1.ApplicationDto}
// @Router /api/v1/repositories/{id}/applications [GET]
func (r repositoryApi) GetApplicationsById(context echo.Context) error {
	repoId := context.Param("id")
	if repoId == "" {
		return errors.New("repository id is required")
	}
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return errors.New("company id is required")
	}
	option := getQueryOption(context)
	status := getStatusOption(context)
	applications, total := r.applicationService.GetByCompanyIdAndRepoId(companyId, repoId, true, option, true, status)
	var applicationsDto []v1.ApplicationDto
	for _, eachApp := range applications {
		applicationsDto = append(applicationsDto, v1.ApplicationDto{
			MetaData: eachApp.MetaData,
			Url:      eachApp.Url,
			Webhook:  eachApp.Webhook,
			Status:   eachApp.Status,
		})
	}
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(applicationsDto)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?companyId=" + context.QueryParam("companyId") + "&status=" + context.QueryParam("status") + "&loadApplications=" + context.QueryParam("loadApplications") + "&loadRepositories=" + context.QueryParam("loadRepositories") + "&loadToken=" + context.QueryParam("loadToken") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?companyId=" + context.QueryParam("companyId") + "&status=" + context.QueryParam("status") + "&loadApplications=" + context.QueryParam("loadApplications") + "&loadRepositories=" + context.QueryParam("loadRepositories") + "&loadToken=" + context.QueryParam("loadToken") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})

	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?companyId=" + context.QueryParam("companyId") + "&status=" + context.QueryParam("status") + "&loadApplications=" + context.QueryParam("loadApplications") + "&loadRepositories=" + context.QueryParam("loadRepositories") + "&loadToken=" + context.QueryParam("loadToken") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, applicationsDto, &metadata, "")
}

// NewRepositoryApi returns Repository type api
func NewRepositoryApi(repositoryService service.Repository, applicationService service.Application, observerList []service.Observer) api.Repository {
	return &repositoryApi{
		repositoryService:  repositoryService,
		applicationService: applicationService,
		observerList:       observerList,
	}
}
