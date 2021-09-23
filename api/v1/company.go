package v1

import (
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/labstack/echo/v4"
)

type companyApi struct {
	companyService service.Company
	observerList []service.Observer
}

func (c companyApi) Save(context echo.Context) error {
	panic("implement me")
}

func (c companyApi) GetById(context echo.Context) error {
	panic("implement me")
}

func (c companyApi) GetRepositoriesById(context echo.Context) error {
	panic("implement me")
}

func NewCompanyApi(companyService service.Company,observerList []service.Observer) api.Company {
	return &companyApi{
		companyService: companyService,
		observerList: observerList,
	}
}
