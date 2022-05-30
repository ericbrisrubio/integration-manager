package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/labstack/echo/v4"
)

type pipelineApi struct {
	pipelineService service.Pipeline
}

// Update... Update pipeline
// @Summary  Update Pipeline
// @Description Update Pipeline by company id, repository id, application url
// @Tags Pipeline
// @Accept json
// @Produce json
// @Param companyId query string true "Company id"
// @Param repositoryId query string true "Repository id"
// @Param url query string true "Application url"
// @Success 200 {object} common.ResponseDTO{data=v1.DirectoryContentCreateAndUpdateResponse}
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/pipelines [PUT]
func (p pipelineApi) Update(context echo.Context) error {
	var payload v1.DirectoryContentUpdatePayload
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company id is required", "Operation failed")
	}
	repoId := context.QueryParam("repositoryId")
	if repoId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Repository id is required", "Operation failed")
	}
	url := context.QueryParam("url")
	if url == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Url is required", "Operation failed")
	}
	if err := context.Bind(&payload); err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation failed")
	}
	res, err := p.pipelineService.Update(companyId, repoId, url, payload)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation failed")
	}
	return common.GenerateSuccessResponse(context, res, nil, "Successful")

}

// Create... Create pipeline
// @Summary  Create Pipeline
// @Description Create Pipeline by company id, repository id, application url
// @Tags Pipeline
// @Accept json
// @Produce json
// @Param companyId query string true "Company id"
// @Param repositoryId query string true "Repository id"
// @Param url query string true "Application url"
// @Success 200 {object} common.ResponseDTO{data=v1.DirectoryContentCreateAndUpdateResponse}
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/pipelines [POST]
func (p pipelineApi) Create(context echo.Context) error {
	var payload v1.DirectoryContentCreatePayload
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company id is required", "Operation failed")
	}
	repoId := context.QueryParam("repositoryId")
	if repoId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Repository id is required", "Operation failed")
	}
	url := context.QueryParam("url")
	if url == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Url is required", "Operation failed")
	}
	if err := context.Bind(&payload); err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation failed")
	}
	res, err := p.pipelineService.Create(companyId, repoId, url, payload)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation failed")
	}
	return common.GenerateSuccessResponse(context, res, nil, "Successful")
}

// Get... Get pipeline for validation
// @Summary  Get Pipeline for validation
// @Description Get Pipeline for validation by company id, repository id, application url and revision
// @Tags Pipeline
// @Accept json
// @Produce json
// @Param action query string true "action [GET_PIPELINE_FOR_VALIDATION]"
// @Param companyId query string true "company id"
// @Param repositoryId query string true "repository id"
// @Param url query string true "application url"
// @Param revision query string true "commit id or branch name"
// @Success 200 {object} common.ResponseDTO{data=v1.PipelineForValidation}
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/pipelines [GET]
func (p pipelineApi) Get(context echo.Context) error {
	if context.QueryParam("action") == string(enums.GET_PIPELINE_FOR_VALIDATION) {
		companyId := context.QueryParam("companyId")
		if companyId == "" {
			return common.GenerateErrorResponse(context, "[ERROR]: Company id is required", "Operation failed")
		}
		repoId := context.QueryParam("repositoryId")
		if repoId == "" {
			return common.GenerateErrorResponse(context, "[ERROR]: Repository id is required", "Operation failed")
		}
		url := context.QueryParam("url")
		if url == "" {
			return common.GenerateErrorResponse(context, "[ERROR]: Url is required", "Operation failed")
		}
		revision := context.QueryParam("revision")
		if revision == "" {
			revision = "master"
		}
		pipelineForValidation, err := p.pipelineService.GetPipelineForValidation(companyId, repoId, url, revision)
		if err != nil {
			return common.GenerateErrorResponse(context, err.Error(), "Operation failed")
		}
		return common.GenerateSuccessResponse(context, pipelineForValidation, nil, "Successful")
	}
	return common.GenerateErrorResponse(context, "[ERROR]: Action is required", "Operation failed")
}

// NewPipelineApi returns Pipeline type api
func NewPipelineApi(pipelineService service.Pipeline) api.Pipeline {
	return &pipelineApi{
		pipelineService: pipelineService,
	}
}
