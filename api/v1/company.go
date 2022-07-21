package v1

import (
	uuid "github.com/google/uuid"
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
// @Success 200 {object} common.ResponseDTO{data=[]v1.ApplicationDto}
// @Router /api/v1/companies/{id}/applications [GET]
func (c companyApi) GetApplicationsByCompanyIdAndRepositoryType(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, nil, "Company Id is required!")
	}
	repositoryType := context.QueryParam("repository_type")
	status := getStatusOption(context)
	option := getQueryOption(context)
	applications, total := c.applicationService.GetApplicationsByCompanyIdAndRepositoryType(id, enums.REPOSITORY_TYPE(repositoryType), false, option, false, status)
	if total == 0 {
		return common.GenerateErrorResponse(context, nil, "Company Id is not found!")
	}
	return common.GenerateSuccessResponse(context, applications, nil, "success")
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
	var payload []v1.RepositoryDto
	payload = formData.Repositories
	for _, each := range payload {
		for j, eachApp := range each.Applications {
			each.Applications[j].Url = UrlFormatter(eachApp.Url)
		}
	}
	var options v1.RepositoryUpdateOption
	Option := context.QueryParam("companyUpdateOption")
	options.Option = enums.REPOSITORY_UPDATE_OPTION(Option)
	company := c.companyService.GetByCompanyId(id)
	if company.Id == "" {
		return common.GenerateErrorResponse(context, "[ERROR] Company does not exist", "Operation Failed")
	}
	err := c.repositoryService.UpdateRepositories(id, payload, options)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	return common.GenerateSuccessResponse(context, nil,
		nil, "Operation Successful")
}

// Update... Update repositories
// @Summary Update repositories by company id
// @Description updates repositories
// @Tags Company
// @Produce json
// @Param data body v1.RepositoriesDto true "RepositoriesDto data"
// @Param id path string true "Company id"
// @Param repoId path string true "Repository id"
// @Param companyUpdateOption query string true "Company Update Option"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/companies/{id}/repositories/{repoId}/applications [PUT]
func (c companyApi) UpdateApplications(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company id is required", "Operation failed")
	}
	repoId := context.Param("repoId")
	if repoId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Repository id is required", "Operation failed")
	}
	company := c.companyService.GetByCompanyId(id)
	if company.Id == "" {
		return common.GenerateErrorResponse(context, "[ERROR] Company does not exist", "Operation failed")
	}
	repository := c.repositoryService.GetById(id, repoId)
	if repository.Id == "" {
		return common.GenerateErrorResponse(context, "[ERROR] Repository does not exist", "Operation failed")
	}
	updateOption := context.QueryParam("companyUpdateOption")
	var options v1.ApplicationUpdateOption
	options.Option = enums.APPLICATION_UPDATE_OPTION(updateOption)
	var formData v1.Applications
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, err.Error(), "Operation failed")
	}
	var payload []v1.Application
	payload = formData.Applications
	if options.Option == enums.APPEND_APPLICATION {
		for i, _ := range payload {
			payload[i].MetaData.Id = uuid.New().String()
			payload[i].Url = UrlFormatter(payload[i].Url)
			if payload[i].MetaData.Labels == nil {
				payload[i].MetaData.Labels = make(map[string]string)
			}
			payload[i].MetaData.Labels["companyId"] = id
			payload[i].CompanyId = id
			payload[i].RepositoryId = repoId
			payload[i].RepositoryType = repository.Type
		}
	}
	err := c.applicationService.UpdateApplications(repository, payload, options)
	if err != nil {
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
// @Success 200 {object} common.ResponseDTO{data=[]v1.CompanyDto}
// @Router /api/v1/companies [GET]
func (c companyApi) Get(context echo.Context) error {
	option := getQueryOption(context)
	status := getStatusOption(context)
	var companies []v1.CompanyDto
	data := c.companyService.GetCompanies(option, status)
	for _, eachCompany := range data {
		company := v1.CompanyDto{
			MetaData: eachCompany.MetaData,
			Id:       eachCompany.Id,
			Name:     eachCompany.Name,
			Status:   eachCompany.Status,
		}
		if option.LoadRepositories {
			repositories, _ := c.repositoryService.GetByCompanyId(eachCompany.Id, false, v1.CompanyQueryOption{})
			for _, eachRepo := range repositories {
				repository := v1.RepositoryDto{
					Id:    eachRepo.Id,
					Type:  eachRepo.Type,
					Token: eachRepo.Token,
				}
				if option.LoadApplications {
					applications, _ := c.applicationService.GetByCompanyIdAndRepoId(eachCompany.Id, eachRepo.Id, false, v1.CompanyQueryOption{}, false, v1.StatusQueryOption{})
					for _, eachApp := range applications {
						repository.Applications = append(repository.Applications, v1.ApplicationDto{
							MetaData: eachApp.MetaData,
							Url:      eachApp.Url,
							Webhook:  eachApp.Webhook,
							Status:   eachApp.Status,
						})
					}
				}
				company.Repositories = append(company.Repositories, repository)
			}
		}
		companies = append(companies, company)
	}
	return common.GenerateSuccessResponse(context, companies, nil, "Success!")
}

// Save... Save company
// @Summary Save company
// @Description Saves company
// @Tags Company
// @Produce json
// @Param data body v1.CompanyDto true "Company data"
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
		Status:   enums.ACTIVE,
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
			if eachApp.MetaData.Labels == nil {
				eachApp.MetaData.Labels = make(map[string]string)
			}
			eachApp.MetaData.Labels["companyId"] = company.Id
			applications = append(applications, v1.Application{
				MetaData:       eachApp.MetaData,
				RepositoryId:   eachRepo.Id,
				RepositoryType: eachRepo.Type,
				CompanyId:      contextData.Id,
				Url:            eachApp.Url,
				Webhook:        eachApp.Webhook,
				Status:         enums.ACTIVE,
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
// @Param loadRepositories query bool false "Loads RepositoriesDto"
// @Param loadApplications query bool false "Loads ApplicationsDto"
// @Param loadToken query bool false "Loads TokenDto"
// @Param id path string true "Company id"
// @Success 200 {object} common.ResponseDTO{data=v1.CompanyDto}
// @Router /api/v1/companies/{id} [GET]
func (c companyApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, nil, "Company Id is required!")
	}
	action := context.QueryParam("action")
	if action == "dashboard_data" {
		repositories, repoCount := c.repositoryService.GetByCompanyId(id, false, v1.CompanyQueryOption{})
		var enabled, disabled int64
		for _, eachRepo := range repositories {
			applications, _ := c.applicationService.GetByCompanyIdAndRepoId(id, eachRepo.Id, false, v1.CompanyQueryOption{}, false, v1.StatusQueryOption{})
			for _, eachApp := range applications {
				if eachApp.Webhook.Active {
					enabled++
				} else {
					disabled++
				}
			}
		}
		data := v1.DashboardData{
			Repository: struct {
				Count int64 `json:"count"`
			}{Count: repoCount},
			Application: struct {
				Webhook struct {
					Enabled  int64 `json:"enabled"`
					Disabled int64 `json:"disabled"`
				} `json:"webhook"`
			}{Webhook: struct {
				Enabled  int64 `json:"enabled"`
				Disabled int64 `json:"disabled"`
			}(struct {
				Enabled  int64
				Disabled int64
			}{Enabled: enabled, Disabled: disabled})},
		}
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
		repositories, _ := c.repositoryService.GetByCompanyId(id, false, v1.CompanyQueryOption{})
		for _, eachRepo := range repositories {
			repositoryDto := v1.RepositoryDto{
				Id:   eachRepo.Id,
				Type: eachRepo.Type,
			}
			if option.LoadToken {
				repositoryDto.Token = eachRepo.Token
			}
			if option.LoadApplications {
				applications, _ := c.applicationService.GetByCompanyIdAndRepoId(id, eachRepo.Id, false, v1.CompanyQueryOption{}, false, v1.StatusQueryOption{})
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
// @Param loadToken query bool false "Loads TokenDto"
// @Success 200 {object} common.ResponseDTO{data=[]v1.RepositoryDto}
// @Router /api/v1/companies/{id}/repositories [GET]
func (c companyApi) GetRepositoriesById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, nil, "Company Id is required!")
	}
	option := getQueryOption(context)
	var repositoriesDto []v1.RepositoryDto
	repositories, total := c.repositoryService.GetByCompanyId(id, true, option)
	for _, eachRepo := range repositories {
		repositoryDto := v1.RepositoryDto{
			Id:   eachRepo.Id,
			Type: eachRepo.Type,
		}
		if option.LoadToken {
			repositoryDto.Token = eachRepo.Token
		}
		if option.LoadApplications {
			applications, _ := c.applicationService.GetByCompanyIdAndRepoId(id, eachRepo.Id, false, v1.CompanyQueryOption{}, false, v1.StatusQueryOption{})
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
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(repositoriesDto)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?loadApplications=" + context.QueryParam("loadApplications") + "&loadRepositories=" + context.QueryParam("loadRepositories") + "&loadToken=" + context.QueryParam("loadToken") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?loadApplications=" + context.QueryParam("loadApplications") + "&loadRepositories=" + context.QueryParam("loadRepositories") + "&loadToken=" + context.QueryParam("loadToken") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})

	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?loadApplications=" + context.QueryParam("loadApplications") + "&loadRepositories=" + context.QueryParam("loadRepositories") + "&loadToken=" + context.QueryParam("loadToken") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, repositoriesDto, &metadata, "")
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
	repository := c.repositoryService.GetById(id, repoId)
	if repository.Id == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: repository not found", "Please provide valid repository id")
	}
	err := c.applicationService.UpdateWebhook(repository, url, webhookId, action)
	if err != nil {
		return common.GenerateErrorResponse(context, err, err.Error())
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Webhook updated sucessfully")
}

func generateRepositoryAndApplicationId(payload v1.CompanyDto) v1.CompanyDto {
	for i, each := range payload.Repositories {
		payload.Repositories[i].Id = uuid.New().String()
		for j := range each.Applications {
			payload.Repositories[i].Applications[j].MetaData.Id = uuid.New().String()
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
