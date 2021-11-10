package v1

import (
	"errors"
	"github.com/klovercloud-ci/api/common"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type repositoryApi struct {
	companyService service.Company
	observerList   []service.Observer
}

// Save... Save repositories
// @Summary Save repositories
// @Description Saves repositories
// @Tags Repository
// @Produce json
// @Param data body v1.CompanyWithUpdateOption true "CompanyWithUpdateOption data"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/repositories [POST]
//func (r repositoryApi) Save(context echo.Context) error {
//	//formData := v1.CompanyWithUpdateOption{}
//	//if err := context.Bind(&formData); err != nil {
//	//	log.Println("Input Error:", err.Error())
//	//	return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
//	//}
//	//
//	//var payload = v1.Company{
//	//	MetaData:     formData.MetaData,
//	//	Id:           formData.Id,
//	//	Name:         formData.Name,
//	//	Repositories: formData.Repositories,
//	//	Status:       enums.ACTIVE,
//	//}
//	//var options v1.CompanyUpdateOption
//	//options.Option = formData.Option
//	//err := r.companyService.UpdateRepositories(payload, options)
//	//if err != nil {
//	//	return common.GenerateErrorResponse(context, nil, "Operation Failed!")
//	//}
//	//return common.GenerateSuccessResponse(context, payload,
//	//	nil, "Operation Successful")
//}

// Get.. Get Repository by id
// @Summary Get Repository by id
// @Description Gets Repository by id
// @Tags Repository
// @Produce json
// @Param id path string true "Repository id"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/repositories/{id} [GET]
func (r repositoryApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	data := r.companyService.GetRepositoryByRepositoryId(id)
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

// Get.. Get Applications by repository id
// @Summary Get Applications by repository id
// @Description Gets Applications by repository id
// @Tags Repository
// @Produce json
// @Param id path string true "Repository id"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/repositories/{id}/applications [GET]
func (r repositoryApi) GetApplicationsById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	option := getQueryOption(context)
	data, total := r.companyService.GetApplicationsByCompanyId(id, option)
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

func NewRepositoryApi(companyService service.Company, observerList []service.Observer) api.Repository {
	return &repositoryApi{
		companyService: companyService,
		observerList:   observerList,
	}
}
