package v1

import (
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/klovercloud-ci/api/common"
	"github.com/klovercloud-ci/config"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
	"strings"
)

type companyApi struct {
	companyService service.Company
	observerList   []service.Observer
}

func (c companyApi) GetCompanies(context echo.Context) error {
	option := getQueryOption(context)
	data := c.companyService.GetCompanies(option)
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
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
	if payload.MetaData.NumberOfConcurrentBuild == 0 {
		PerDayConcurrentBuild, _ := strconv.ParseInt(config.PerDayConcurrentBuild, 10, 64)
		payload.MetaData.NumberOfConcurrentBuild = PerDayConcurrentBuild
	}
	if payload.MetaData.TotalBuildPerDay == 0 {
		PerDayConcurrentBuild, _ := strconv.ParseInt(config.PerDayConcurrentBuild, 10, 64)
		payload.MetaData.TotalBuildPerDay = PerDayConcurrentBuild
	}
	contextData, er := validate(payload)
	if er != nil {
		return common.GenerateErrorResponse(context, nil, "invalid repository type!")
	}
	err := c.companyService.Store(contextData)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, contextData,
		nil, "saved Successfully")
}

func validate(payload v1.Company) (v1.Company, error) {
	comp := v1.Company{}
	comp = payload
	for i, each := range payload.Repositories {
		if each.Type == enums.BIT_BUCKET || each.Type == enums.GITHUB {
			comp.Repositories[i].Id = guuid.New().String()
			//each.Id = uuid.New().String()
			for j, eachApp := range each.Applications {
				comp.Repositories[i].Applications[j].MetaData.Id = guuid.New().String()
				//eachApp.MetaData.Id = uuid.New().String()
				fmt.Println("meta data----", eachApp)
			}
		} else {
			return comp, errors.New("Ivalid repository type!")
		}
		fmt.Println("object----", comp)
	}
	return comp, nil
}

func (c companyApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	option := getQueryOption(context)

	data, _ := c.companyService.GetByCompanyId(id, option)
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}
func getQueryOption(context echo.Context) v1.CompanyQueryOption {
	option := v1.CompanyQueryOption{}
	page := context.QueryParam("page")
	limit := context.QueryParam("limit")
	la := context.QueryParam("LoadApplications")
	lr := context.QueryParam("LoadRepositories")
	if page == "" {
		option.Pagination.Page = 0
		option.Pagination.Limit = 10
		option.LoadApplications, _ = strconv.ParseBool(la)
		option.LoadRepositories, _ = strconv.ParseBool(lr)
	} else {
		option.Pagination.Page, _ = strconv.ParseInt(page, 10, 64)
		option.Pagination.Limit, _ = strconv.ParseInt(limit, 10, 64)
		option.LoadApplications, _ = strconv.ParseBool(la)
		option.LoadRepositories, _ = strconv.ParseBool(lr)
	}
	return option
}

func (c companyApi) GetRepositoriesById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	option := getQueryOption(context)
	data, total := c.companyService.GetRepositoriesByCompanyId(id, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(data)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})

	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, data, &metadata, "")
}

func NewCompanyApi(companyService service.Company, observerList []service.Observer) api.Company {
	return &companyApi{
		companyService: companyService,
		observerList:   observerList,
	}
}
