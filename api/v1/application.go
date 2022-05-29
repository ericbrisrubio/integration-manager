package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
	"strings"
)

type applicationApi struct {
	companyService  service.Company
	observerList    []service.Observer
	pipelineService service.Pipeline
}

// GetAll.. Get All Applications
// @Summary Get All Applications
// @Description Get All Applications
// @Tags Application
// @Produce json
// @Param companyId query string true "company id"
// @Success 200 {object} common.ResponseDTO{data=[]v1.Application}
// @Router /api/v1/applications [GET]

func (a applicationApi) GetAll(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company id is required", "Operation failed")
	}
	option := getQueryOption(context)
	option.LoadRepositories = true
	option.LoadApplications = true
	data, total := a.companyService.GetAllApplications(companyId, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(data)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
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
	data := a.companyService.GetApplicationByApplicationId(companyId, repositoryId, id)
	if data.MetaData.Id == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company not found", "Operation failed")
	}
	return common.GenerateSuccessResponse(context, data,
		nil, "Successful")
}

// Update... Update Application
// @Summary  Update Application
// @Description Update Application by company id and  repository id
// @Tags Application
// @Accept json
// @Produce json
// @Param data body v1.ApplicationsDto true "ApplicationsDto Data"
// @Param companyId query string true "company id"
// @Param repositoryId query string true "repository id"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/applications [POST]
func (a applicationApi) Update(context echo.Context) error {
	var formData v1.ApplicationsDto
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company id is required", "Operation failed")
	}
	repoId := context.QueryParam("repositoryId")
	if repoId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Repository id is required", "Operation failed")
	}
	updateOption := context.QueryParam("companyUpdateOption")
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, err.Error(), "Operation failed")
	}
	var payload []v1.Application
	payload = formData.Applications
	for i, _ := range payload {
		payload[i].Url = UrlFormatter(payload[i].Url)
		if payload[i].MetaData.Labels == nil {
			payload[i].MetaData.Labels = make(map[string]string)
		}
		payload[i].MetaData.Labels["CompanyId"] = companyId
	}
	var options v1.ApplicationUpdateOption
	options.Option = enums.APPLICATION_UPDATE_OPTION(updateOption)
	err := a.companyService.UpdateApplications(companyId, repoId, payload, options)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	return common.GenerateSuccessResponse(context, payload,
		nil, "Operation Successful")
}

// Get... Get pipeline for validation
// @Summary  Get Pipeline for validation
// @Description Get Pipeline for validation by company id, repository id, application url and revision
// @Tags Application
// @Accept json
// @Produce json
// @Param companyId query string true "company id"
// @Param repositoryId query string true "repository id"
// @Param url query string true "application url"
// @Param revision query string true "commit id or branch name"
// @Success 200 {object} common.ResponseDTO{data=v1.PipelineForValidation}
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/applications/{id}/pipelines [GET]
func (a applicationApi) GetPipelineForValidation(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company id is required", "Operation failed")
	}
	repoId := context.QueryParam("repositoryId")
	if repoId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Repository id is required", "Operation failed")
	}
	appId := context.Param("applicationId")
	if appId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Application id is required", "Operation failed")
	}
	revision := context.QueryParam("revision")
	if revision == "" {
		revision = "master"
	}
	application := a.companyService.GetApplicationByApplicationId(companyId, repoId, appId)
	pipelineForValidation, err := a.pipelineService.GetPipelineForValidation(companyId, repoId, application.Url, revision)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation failed")
	}
	return common.GenerateSuccessResponse(context, pipelineForValidation, nil, "Successful")
}

// NewApplicationApi returns Application type api
func NewApplicationApi(companyService service.Company, observerList []service.Observer, pipelineService service.Pipeline) api.Application {
	return &applicationApi{
		companyService:  companyService,
		observerList:    observerList,
		pipelineService: pipelineService,
	}
}
