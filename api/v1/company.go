package v1

import (
	"errors"
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
// Get... Get companies
// @Summary Get companies
// @Description Gets companies
// @Tags Company
// @Produce json
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Param loadRepositories query int64 false "Loads Repositories"
// @Param loadApplications query int64 false "Loads Applications"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/ [GET]
func (c companyApi) GetCompanies(context echo.Context) error {
	option := getQueryOption(context)
	data := c.companyService.GetCompanies(option)
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

// Save... Save company
// @Summary Save company
// @Description Saves company
// @Tags Company
// @Produce json
// @Param data body v1.Company true "Company data"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies [POST]
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
	if payload.MetaData.NumberOfConcurrentProcess == 0 {
		payload.MetaData.NumberOfConcurrentProcess = config.DefaultNumberOfConcurrentProcess
	}
	if payload.MetaData.TotalProcessPerDay == 0 {
		payload.MetaData.TotalProcessPerDay = config.DefaultPerDayTotalProcess
	}
	contextData := generateRepositoryAndApplicationId(payload)
	err := c.companyService.Store(contextData)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, contextData,
		nil, "Operation Successful")
}

// Get.. Get company
// @Summary Get company by id
// @Description Gets company by id
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id} [GET]
func (c companyApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	option := getQueryOption(context)

	data, _ := c.companyService.GetByCompanyId(id, option)
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}
// Get.. Get Repositories by company id
// @Summary Get Repositories by company id
// @Description Gets Repositories by company id
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id}/repositories [GET]
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


func generateRepositoryAndApplicationId(payload v1.Company) v1.Company{
	comp := v1.Company{}
	comp = payload
	for i, each := range payload.Repositories {
		comp.Repositories[i].Id = guuid.New().String()
		for j, _ := range each.Applications {
			comp.Repositories[i].Applications[j].MetaData.Id = guuid.New().String()
		}
	}
	return comp
}

func getQueryOption(context echo.Context) v1.CompanyQueryOption {
	option := v1.CompanyQueryOption{}
	page := context.QueryParam("page")
	limit := context.QueryParam("limit")
	loadApplications := context.QueryParam("loadApplications")
	loadRepositories := context.QueryParam("loadRepositories")
	if page == "" {
		option.Pagination.Page = 0
		option.Pagination.Limit = 10
		option.LoadApplications, _ = strconv.ParseBool(loadApplications)
		option.LoadRepositories, _ = strconv.ParseBool(loadRepositories)
	} else {
		option.Pagination.Page, _ = strconv.ParseInt(page, 10, 64)
		option.Pagination.Limit, _ = strconv.ParseInt(limit, 10, 64)
		option.LoadApplications, _ = strconv.ParseBool(loadApplications)
		option.LoadRepositories, _ = strconv.ParseBool(loadRepositories)
	}
	return option
}

func NewCompanyApi(companyService service.Company, observerList []service.Observer) api.Company {
	return &companyApi{
		companyService: companyService,
		observerList:   observerList,
	}
}
