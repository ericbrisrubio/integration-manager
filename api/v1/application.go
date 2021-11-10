package v1

import (
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"log"
)

type applicationApi struct {
	companyService service.Company
	observerList   []service.Observer
}

// Update ... Update Application
// @Summary  Update Application
// @Description Update Application by company id and  repository id
// @Tags Application
// @Accept json
// @Produce json
// @Param data body v1.ListOfApplications true "ListOfApplications Data"
// @Param company_id query string true "company id"
// @Param repository_Id query string true "repository id"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/applications [POST]
func (a applicationApi) UpdateApplication(context echo.Context) error {
	var formData v1.ListOfApplications
	id := context.QueryParam("company_id")
	repoId := context.QueryParam("repository_Id")
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

func NewApplicationApi(companyService service.Company, observerList []service.Observer) api.Application {
	return &applicationApi{
		companyService: companyService,
		observerList:   observerList,
	}
}
