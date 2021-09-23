package v1

import (
	"errors"
	"fmt"
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
)

type companyApi struct {
	companyService service.Company
	observerList   []service.Observer
}

func (c companyApi) Save(context echo.Context) error {
	formData := v1.Company{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	var payload = v1.Company{
		MetaData:     formData.MetaData,
		Id:           formData.Id,
		Name:         formData.Name,
		Repositories: formData.Repositories,
		Status:       enums.ACTIVE,
	}
	err := c.companyService.Store(payload)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, payload,
		nil, "saved Successfully")
}

func (c companyApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	limitStr := context.QueryParam("limit")
	var limit int64 = 20
	if limitStr != "" {
		limitInt, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			log.Println(err)
		}
		limit = limitInt
	}
	fmt.Println(limit)
	pageStr := context.QueryParam("page")
	var page int64 = 1
	if pageStr != "" {
		pageInt, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			log.Println(err)
		}
		page = pageInt
	}
	fmt.Println(page)

	loadRepos := context.QueryParam("LoadRepositories")
	var lr bool = true
	if loadRepos != "" {
		LoadRepository, err := strconv.ParseBool(loadRepos)
		if err != nil {
			log.Println(err)
		}
		lr = LoadRepository
	}
	loadApps := context.QueryParam("LoadApplications")
	var la bool = true
	if loadRepos != "" {
		LoadRepository, err := strconv.ParseBool(loadApps)
		if err != nil {
			log.Println(err)
		}
		la = LoadRepository
	}

	data := c.companyService.GetByCompanyId(id, v1.CompanyQueryOption{
		Pagination: struct {
			Page  int64
			Limit int64
		}{Page: page, Limit: limit},
		LoadApplications: la,
		LoadRepositories: lr,
	})

	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (c companyApi) GetRepositoriesById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	limitStr := context.QueryParam("limit")
	var limit int64 = 20
	if limitStr != "" {
		limitInt, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			log.Println(err)
		}
		limit = limitInt
	}
	fmt.Println(limit)
	pageStr := context.QueryParam("page")
	var page int64 = 1
	if pageStr != "" {
		pageInt, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			log.Println(err)
		}
		page = pageInt
	}
	fmt.Println(page)

	loadRepos := context.QueryParam("LoadRepositories")
	var lr bool = true
	if loadRepos != "" {
		LoadRepository, err := strconv.ParseBool(loadRepos)
		if err != nil {
			log.Println(err)
		}
		lr = LoadRepository
	}
	loadApps := context.QueryParam("LoadApplications")
	var la bool = true
	if loadRepos != "" {
		LoadRepository, err := strconv.ParseBool(loadApps)
		if err != nil {
			log.Println(err)
		}
		la = LoadRepository
	}
	data := c.companyService.GetRepositoriesByCompanyId(id, v1.CompanyQueryOption{
		Pagination: struct {
			Page  int64
			Limit int64
		}{Page: page, Limit: limit},
		LoadApplications: la,
		LoadRepositories: lr,
	})
	if data == nil {
		err := common.GenerateSuccessResponse(context, nil, nil, "Database Error!")
		if err != nil {
			return err
		}
	}
	return common.GenerateSuccessResponse(context, data, nil, "Database Error!")
}

func NewCompanyApi(companyService service.Company, observerList []service.Observer) api.Company {
	return &companyApi{
		companyService: companyService,
		observerList:   observerList,
	}
}
