package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/labstack/echo/v4"
)

type pipelineApi struct {
	pipelineService service.Pipeline
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
