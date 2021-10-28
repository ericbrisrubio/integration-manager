package v1

import (
	"errors"
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/labstack/echo/v4"
	"log"
)

type applicationApi struct {
	companyService service.Company
	observerList   []service.Observer
}


func (a applicationApi) UpdateApplication(context echo.Context) error {
	var formData v1.ApplicationWithUpdateOption
	id := context.QueryParam("company_id")
	if id == "" {
		return errors.New("Id required!")
	}
	repoId := context.QueryParam("repository_Id")
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	var payload []v1.Application
	payload = formData.Applications
	var options v1.CompanyUpdateOption
	options.Option = formData.Option
	err := a.companyService.UpdateApplications(id, repoId, payload, options)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, "Database error!")
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