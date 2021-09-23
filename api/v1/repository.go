package v1

import (
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/labstack/echo/v4"
)

type repositoryApi struct {
	companyService service.Company
	observerList []service.Observer
}

func (r repositoryApi) Save(context echo.Context) error {
	panic("implement me")
}

func (r repositoryApi) GetById(context echo.Context) error {
	panic("implement me")
}

func (r repositoryApi) GetApplicationsById(context echo.Context) error {
	panic("implement me")
}

func NewRepositoryApi(companyService service.Company,observerList []service.Observer) api.Repository {
	return &repositoryApi{
		companyService: companyService,
		observerList: observerList,
	}
}
