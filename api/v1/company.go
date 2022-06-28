package v1

import (
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
	companyService     service.Company
	repositoryService  service.Repository
	applicationService service.Application
	githubService      service.Git
	bitbucketService   service.Git
	observerList       []service.Observer
}

// Get.. Get applications
// @Summary Get applications by company id and repository type
// @Description Get applications by company id and repository type
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Param repository_type query string true "Repository type"
// @Param companyUpdateOption query string true "Company Update Option"
// @Success 200 {object} common.ResponseDTO{data=[]v1.Application}
// @Router /api/v1/companies/{id}/applications [GET]
func (c companyApi) GetApplicationsByCompanyIdAndRepositoryType(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, nil, "Company Id is required!")
	}
	repositoryType := context.QueryParam("repository_type")
	status := getStatusOption(context)
	option := getQueryOption(context)
	apps := c.companyService.GetApplicationsByCompanyIdAndRepositoryType(id, enums.REPOSITORY_TYPE(repositoryType), option, status)
	if apps == nil {
		return common.GenerateErrorResponse(context, nil, "Company Id is not found!")
	}
	return common.GenerateSuccessResponse(context, apps, nil, "success")
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
		return common.GenerateErrorResponse(context, nil, "Company Id is required!")
	}
	var payload []v1.Repository
	payload = formData.Repositories
	for _, each := range payload {
		for j, eachApp := range each.Applications {
			each.Applications[j].Url = UrlFormatter(eachApp.Url)
		}
	}
	var options v1.RepositoryUpdateOption
	Option := context.QueryParam("companyUpdateOption")
	options.Option = enums.REPOSITORY_UPDATE_OPTION(Option)
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
// @Success 200 {object} common.ResponseDTO{data=[]v1.Company}
// @Router /api/v1/companies [GET]
func (c companyApi) Get(context echo.Context) error {
	option := getQueryOption(context)
	status := getStatusOption(context)
	data := c.companyService.GetCompanies(option, status)
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
	payload := v1.CompanyDto{}
	if err := context.Bind(&payload); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	if payload.MetaData.NumberOfConcurrentProcess == 0 {
		payload.MetaData.NumberOfConcurrentProcess = config.DefaultNumberOfConcurrentProcess
	}
	if payload.MetaData.TotalProcessPerDay == 0 {
		payload.MetaData.TotalProcessPerDay = config.DefaultPerDayTotalProcess
	}
	contextData := generateRepositoryAndApplicationId(payload)
	for _, each := range contextData.Repositories {
		for j, eachApp := range each.Applications {
			each.Applications[j].Url = UrlFormatter(eachApp.Url)
		}
	}
	company := v1.Company{
		MetaData: contextData.MetaData,
		Id:       contextData.Id,
		Name:     contextData.Name,
		Status:   contextData.Status,
	}
	err := c.companyService.Store(company)
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: Company store failed", "Operation Failed!")
	}
	var repositories []v1.Repository
	for _, eachRepo := range contextData.Repositories {
		repositories = append(repositories, v1.Repository{
			Id:        eachRepo.Id,
			CompanyId: contextData.Id,
			Type:      eachRepo.Type,
			Token:     eachRepo.Token,
		})
		var applications []v1.Application
		for _, eachApp := range eachRepo.Applications {
			applications = append(applications, v1.Application{
				MetaData:     eachApp.MetaData,
				RepositoryId: eachRepo.Id,
				CompanyId:    contextData.Id,
				Url:          eachApp.Url,
				Webhook:      eachApp.Webhook,
				Status:       eachApp.Status,
			})
		}
		go c.applicationService.CreateWebHookAndUpdateApplications(eachRepo.Type, eachRepo.Token, applications)
	}
	err = c.repositoryService.Store(repositories)
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: Repositories store failed", "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, contextData,
		nil, "Operation Successful")
}

// Get.. Get company
// @Summary Get company by id
// @Description Gets company by id
// @Tags Company
// @Produce json
// @Param action query string false "action [dashboard_data]"
// @Param id path string true "Company id"
// @Success 200 {object} common.ResponseDTO{data=v1.Company}
// @Router /api/v1/companies/{id} [GET]
func (c companyApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, nil, "Company Id is required!")
	}
	action := context.QueryParam("action")
	if action == "dashboard_data" {
		data := c.companyService.GetDashboardData(id)
		return common.GenerateSuccessResponse(context, data, nil, "Operation Successful")
	}
	company := c.companyService.GetByCompanyId(id)
	companyDto := v1.CompanyDto{
		MetaData: company.MetaData,
		Id:       company.Id,
		Name:     company.Name,
		Status:   company.Status,
	}
	option := getQueryOption(context)
	var repositoriesDto []v1.RepositoryDto
	if option.LoadRepositories {
		repositories := c.repositoryService.GetByCompanyId(id)
		for _, eachRepo := range repositories {
			repositoryDto := v1.RepositoryDto{
				Id:    eachRepo.Id,
				Type:  eachRepo.Type,
				Token: eachRepo.Token,
			}
			if option.LoadApplications {
				applications := c.applicationService.GetByCompanyIdAndRepoId(id, eachRepo.Id)
				var applicationsDto []v1.ApplicationDto
				for _, eachApp := range applications {
					applicationsDto = append(applicationsDto, v1.ApplicationDto{
						MetaData: eachApp.MetaData,
						Url:      eachApp.Url,
						Webhook:  eachApp.Webhook,
						Status:   eachApp.Status,
					})
				}
				repositoryDto.Applications = applicationsDto
			}
			repositoriesDto = append(repositoriesDto, repositoryDto)
		}
	}
	companyDto.Repositories = repositoriesDto
	return common.GenerateSuccessResponse(context, companyDto, nil, "Success!")
}

// Get.. Get RepositoriesDto by company id
// @Summary Get RepositoriesDto by company id
// @Description Gets RepositoriesDto by company id
// @Tags Company
// @Produce json
// @Param id path string true "Company id"
// @Param loadApplications query bool false "Loads ApplicationsDto"
// @Success 200 {object} common.ResponseDTO{data=[]v1.Repository}
// @Router /api/v1/companies/{id}/repositories [GET]
func (c companyApi) GetRepositoriesById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, nil, "Company Id is required!")
	}
	option := getQueryOption(context)
	data, total := c.companyService.GetRepositoriesByCompanyId(id, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(data)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?loadApplications=" + context.QueryParam("loadApplications") + "&loadRepositories=" + context.QueryParam("loadRepositories") + "&loadToken=" + context.QueryParam("loadToken") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?loadApplications=" + context.QueryParam("loadApplications") + "&loadRepositories=" + context.QueryParam("loadRepositories") + "&loadToken=" + context.QueryParam("loadToken") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})

	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?loadApplications=" + context.QueryParam("loadApplications") + "&loadRepositories=" + context.QueryParam("loadRepositories") + "&loadToken=" + context.QueryParam("loadToken") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, data, &metadata, "")
}

// UpdateWebhook... Update Webhook
// @Summary Update Webhook to Enable or Disable
// @Description Update Webhook
// @Tags Github
// @Produce json
// @Param id path string true "Company id"
// @Param repoId path string true "Repository id"
// @Param url query string true "Url"
// @Param webhookId query string false "Webhook Id to disable webhook"
// @Param action query string true "action type [enable/disable]"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/companies/{id}/repositories/{repoId}/webhooks [PATCH]
func (c companyApi) UpdateWebhook(context echo.Context) error {
	action := context.QueryParam("action")
	if action != string(enums.WEBHOOK_EANBLE) && action != string(enums.WEBHOOK_DISABLE) {
		return common.GenerateErrorResponse(context, "[ERROR]: invalid action provided", "Provide valid action. [enable/disable]")
	}
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: no company id is not provided", "Please provide company id")
	}
	repoId := context.Param("repoId")
	if repoId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: repository id is not provided", "Please provide repository id")
	}
	url := context.QueryParam("url")
	if url == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: application url is not provided", "Please provide application url")
	}
	webhookId := context.QueryParam("webhookId")
	if webhookId == "" && action == string(enums.WEBHOOK_DISABLE) {
		return common.GenerateErrorResponse(context, "[ERROR]: webhook id is not provided", "Please provide webook id")
	}
	err := c.companyService.UpdateWebhook(id, repoId, url, webhookId, action)
	if err != nil {
		return common.GenerateErrorResponse(context, err, err.Error())
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Webhook updated sucessfully")
}

func generateRepositoryAndApplicationId(payload v1.CompanyDto) v1.CompanyDto {
	for i, each := range payload.Repositories {
		payload.Repositories[i].Id = guuid.New().String()
		for j := range each.Applications {
			payload.Repositories[i].Applications[j].MetaData.Id = guuid.New().String()
		}
	}
	return payload
}
func getStatusOption(context echo.Context) v1.StatusQueryOption {
	status := v1.StatusQueryOption{}
	option := context.QueryParam("status")
	if option != "" {
		status.Option = enums.ACTIVE
	}
	status.Option = enums.COMPANY_STATUS(option)
	return status
}

func getQueryOption(context echo.Context) v1.CompanyQueryOption {
	option := v1.CompanyQueryOption{}
	page := context.QueryParam("page")
	limit := context.QueryParam("limit")
	loadApplications := context.QueryParam("loadApplications")
	loadRepositories := context.QueryParam("loadRepositories")
	loadToken := context.QueryParam("loadToken")
	if page == "" {
		option.Pagination.Page = 0
		option.Pagination.Limit = 10
		option.LoadApplications, _ = strconv.ParseBool(loadApplications)
		option.LoadRepositories, _ = strconv.ParseBool(loadRepositories)
		option.LoadToken, _ = strconv.ParseBool(loadToken)
	} else {
		option.Pagination.Page, _ = strconv.ParseInt(page, 10, 64)
		option.Pagination.Limit, _ = strconv.ParseInt(limit, 10, 64)
		option.LoadApplications, _ = strconv.ParseBool(loadApplications)
		option.LoadRepositories, _ = strconv.ParseBool(loadRepositories)
		option.LoadToken, _ = strconv.ParseBool(loadToken)
	}
	return option
}

// NewCompanyApi returns Company type api
func NewCompanyApi(companyService service.Company, repositoryService service.Repository, applicationService service.Application, githubService service.Git, bitbucketService service.Git, observerList []service.Observer) api.Company {
	return &companyApi{
		companyService:     companyService,
		repositoryService:  repositoryService,
		applicationService: applicationService,
		githubService:      githubService,
		bitbucketService:   bitbucketService,
		observerList:       observerList,
	}
}
