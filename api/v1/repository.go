package v1

import (
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
	"strings"
)

type repositoryApi struct {
	companyService service.Company
	observerList   []service.Observer
}

func (r repositoryApi) Save(context echo.Context) error {
	formData := v1.CompanyWithUpdateOption{}
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
	var options v1.CompanyUpdateOption
	options.Option = formData.Option
	contextData, er := validator(payload)
	if er != nil {
		return common.GenerateErrorResponse(context, nil, "invalid repository type!")
	}
	err := r.companyService.UpdateRepositories(payload, options)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, "Database error!")
	}
	return common.GenerateSuccessResponse(context, contextData,
		nil, "saved Successfully")
}
func validator(payload v1.Company) (v1.Company, error) {
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

func (r repositoryApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return errors.New("Id required!")
	}
	data := r.companyService.GetRepositoryByRepositoryId(id)
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

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
