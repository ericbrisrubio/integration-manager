package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type applicationApi struct {
	applicationService service.Application
	observerList       []service.Observer
	pipelineService    service.Pipeline
}

// GetAll.. Get All Applications
// @Summary Get All Applications
// @Description Get All Applications
// @Tags Application
// @Produce json
// @Param companyId query string true "company id"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Param loadRepositories query bool false "Loads Repositories"
// @Param loadApplications query bool false "Loads Applications"
// @Success 200 {object} common.ResponseDTO{data=[]v1.Application}
// @Router /api/v1/applications [GET]
func (a applicationApi) Get(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company id is required", "Operation failed")
	}
	return a.GetAll(context, companyId)
}

func (a applicationApi) GetAll(context echo.Context, companyId string) error {
	option := getQueryOption(context)
	data, total := a.applicationService.GetAll(companyId, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(data)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, data,
		&metadata, "Successful")
}

// Get.. Get Application by Application id
// @Summary Get Application by Application id
// @Description Gets Application by Application id
// @Tags Application
// @Produce json
// @Param id path string true "Application id"
// @Param companyId query string true "company id"
// @Param repositoryId query string true "repository id"
// @Success 200 {object} common.ResponseDTO{data=v1.Application}
// @Router /api/v1/applications/{id} [GET]
func (a applicationApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Application id is required", "Operation failed")
	}
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company id is required", "Operation failed")
	}
	repositoryId := context.QueryParam("repositoryId")
	if repositoryId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Repository id is required", "Operation failed")
	}
	data := a.applicationService.GetByApplicationId(companyId, repositoryId, id)
	if data.MetaData.Id == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Application not found", "Operation failed")
	}
	return common.GenerateSuccessResponse(context, data,
		nil, "Successful")
}

// NewApplicationApi returns Application type api
func NewApplicationApi(applicationService service.Application, observerList []service.Observer, pipelineService service.Pipeline) api.Application {
	return &applicationApi{
		applicationService: applicationService,
		observerList:       observerList,
		pipelineService:    pipelineService,
	}
}
