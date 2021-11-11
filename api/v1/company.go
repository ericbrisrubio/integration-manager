package v1

import (
	"errors"
	guuid "github.com/google/uuid"
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	"github.com/klovercloud-ci-cd/integration-manager/config"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
	"strings"
)

type companyApi struct {
	companyService service.Company
	observerList   []service.Observer
}

// Update... Update repositories
// @Summary Update repositories by company id
// @Description updates repositories
// @Tags Company
// @Produce json
// @Param data body v1.RepositoriesDto true "RepositoriesDto data"
// @Param id path string true "Company id"
// @Param companyUpdateOption query string true "Company Update Option"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id}/repositories [PUT]
func (c companyApi) UpdateRepositories(context echo.Context) error {
	var formData v1.RepositoriesDto
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	var payload []v1.Repository
	payload = formData.Repositories
	var options v1.CompanyUpdateOption
	Option := context.QueryParam("companyUpdateOption")
	options.Option = enums.COMPANY_UPDATE_OPTION(Option)
	err := c.companyService.UpdateRepositories(id, payload, options)
	if err != nil {
		log.Println("Update Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	return common.GenerateSuccessResponse(context, formData,
		nil, "Operation Successful")
}

// Get... Get companies
// @Summary Get companies
// @Description Gets companies
// @Tags Company
// @Produce json
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Param loadRepositories query bool false "Loads RepositoriesDto"
// @Param loadApplications query bool false "Loads ApplicationsDto"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies [GET]
func (c companyApi) Get(context echo.Context) error {
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

// Get.. Get RepositoriesDto by company id
// @Summary Get RepositoriesDto by company id
// @Description Gets RepositoriesDto by company id
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Param loadApplications query bool false "Loads ApplicationsDto"
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

func generateRepositoryAndApplicationId(payload v1.Company) v1.Company {
	for i, each := range payload.Repositories {
		payload.Repositories[i].Id = guuid.New().String()
		for j := range each.Applications {
			payload.Repositories[i].Applications[j].MetaData.Id = guuid.New().String()
		}
	}
	return payload
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

// NewCompanyApi returns Company type api
func NewCompanyApi(companyService service.Company, observerList []service.Observer) api.Company {
	return &companyApi{
		companyService: companyService,
		observerList:   observerList,
	}
}
