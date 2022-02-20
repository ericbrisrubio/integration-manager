package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/labstack/echo/v4"
	"log"
)

type applicationApi struct {
	companyService service.Company
	observerList   []service.Observer
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
func (a applicationApi) GetApplicationByApplicationId(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, nil, "Company Id is required!")
	}
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return context.JSON(404, common.ResponseDTO{
			Message: "company id is required",
		})
	}
	repositoryId := context.QueryParam("repositoryId")
	if repositoryId == "" {
		return context.JSON(404, common.ResponseDTO{
			Message: "repository id is required",
		})
	}
	data := a.companyService.GetApplicationByApplicationId(companyId, repositoryId, id)
	if data.MetaData.Id == "" {
		return common.GenerateErrorResponse(context, nil, "Company not found!")
	}
	return common.GenerateSuccessResponse(context, data,
		nil, "Successfully")
	//testforgitIgnore
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
	id := context.QueryParam("companyId")
	if id == "" {
		return context.JSON(404, common.ResponseDTO{
			Message: "company id is required",
		})
	}
	repoId := context.QueryParam("repositoryId")
	if repoId == "" {
		return context.JSON(404, common.ResponseDTO{
			Message: "repository id is required",
		})
	}
	updateOption := context.QueryParam("companyUpdateOption")
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	var payload []v1.Application
	payload = formData.Applications
	var options v1.CompanyUpdateOption
	options.Option = enums.COMPANY_UPDATE_OPTION(updateOption)
	err := a.companyService.UpdateApplications(id, repoId, payload, options)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	return common.GenerateSuccessResponse(context, payload,
		nil, "saved Successfully")
}

// NewApplicationApi returns Application type api
func NewApplicationApi(companyService service.Company, observerList []service.Observer) api.Application {
	return &applicationApi{
		companyService: companyService,
		observerList:   observerList,
	}
}
