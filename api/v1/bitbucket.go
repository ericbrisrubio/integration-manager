package v1

import (
	_ "encoding/json"
	"errors"
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/labstack/echo/v4"
	"github.com/twinj/uuid"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"log"
	"strings"
)

type v1BitbucketApi struct {
	gitService                   service.Git
	companyService               service.Company
	processInventoryEventService service.ProcessInventoryEvent
	observerList                 []service.Observer
}

// DisableWebhook... Disable Webhook
// @Summary Disable Webhook
// @Description Disable Webhook
// @Tags Bitbucket
// @Produce json
// @Param userName query string true "User Name"
// @Param repoName query string true "Repository Name"
// @Param companyId query string true "Company Id"
// @Param repoId query string true "Repository Id"
// @Param webhookId query string true "Webhook Id"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/bitbuckets/webhook [DELETE]
func (b v1BitbucketApi) DisableWebhook(context echo.Context) error {
	username := context.QueryParam("userName")
	if username == "" {
		return errors.New("userName is required")
	}
	repoName := context.QueryParam("repoName")
	if repoName == "" {
		return errors.New("repository name is required")
	}
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return errors.New("companyId is required")
	}
	repoId := context.QueryParam("repoId")
	if repoId == "" {
		return errors.New("repoId is required")
	}
	webhookId := context.QueryParam("webhookId")
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
		LoadToken:        true,
	}
	repo := b.companyService.GetRepositoryByRepositoryId(companyId, repoId, option)
	err := b.gitService.DeleteRepositoryWebhookById(username, repoName, webhookId, repo.Token)
	if err != nil {
		return common.GenerateErrorResponse(context, err, err.Error())
	}
	return common.GenerateSuccessResponse(context, nil, nil, "successfully disable webhook")
}

// EnableWebhook... Enable Webhook
// @Summary Enable Webhook
// @Description Enable Webhook
// @Tags Bitbucket
// @Produce json
// @Param userName query string true "User Name"
// @Param repoName query string true "Repository Name"
// @Param companyId query string true "Company Id"
// @Param repoId query string true "Repository Id"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/bitbuckets/webhook [PUT]
func (b v1BitbucketApi) EnableWebhook(context echo.Context) error {
	username := context.QueryParam("userName")
	if username == "" {
		return errors.New("userName is required")
	}
	repoName := context.QueryParam("repoName")
	if repoName == "" {
		return errors.New("repository name is required")
	}
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return errors.New("companyId is required")
	}
	repoId := context.QueryParam("repoId")
	if repoId == "" {
		return errors.New("repoId is required")
	}
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
		LoadToken:        true,
	}
	repo := b.companyService.GetRepositoryByRepositoryId(companyId, repoId, option)
	_, err := b.gitService.CreateRepositoryWebhook(username, repoName, repo.Token, companyId)
	if err != nil {
		return common.GenerateErrorResponse(context, err, err.Error())
	}
	return common.GenerateSuccessResponse(context, nil, nil, "successfully enable webhook")
}

// GetCommitByBranch... Get Commit By Branch
// @Summary Get Commit By Branch
// @Description Get Commit By Branch
// @Tags Bitbucket
// @Produce json
// @Param userName query string true "User Name"
// @Param repoName query string true "Repository Name"
// @Param companyId query string true "Company Id"
// @Param repoId query string true "Repository Id"
// @Param branch query string true "Branch"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/bitbucket/commits [GET]
func (b v1BitbucketApi) GetCommitByBranch(context echo.Context) error {
	repoId := context.QueryParam("repoId")
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
		LoadToken:        true,
	}
	id := context.QueryParam("companyId")
	repo := b.companyService.GetRepositoryByRepositoryId(id, repoId, option)
	userName := context.QueryParam("userName")
	if userName == "" {
		return errors.New("userName is required")
	}
	repoName := context.QueryParam("repoName")
	if repoName == "" {
		return errors.New("repoName is required")
	}
	branch := context.QueryParam("branch")
	commits, err := b.gitService.GetCommitByBranch(userName, repoName, branch, repo.Token)
	if err != nil {
		return common.GenerateErrorResponse(context, err, err.Error())
	}
	return common.GenerateSuccessResponse(context, commits, nil, "success")
}

// GetBranches... Get Branches
// @Summary Get Branches
// @Description Gets Branches
// @Tags Bitbucket
// @Produce json
// @Param repoId query string true "Repository Id"
// @Param companyId query string true "company Id"
// @Param url query string true "Url"
// @Param loadApplications query bool false "Loads ApplicationsDto"
// @Param loadToken query bool true "Loads Token"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/bitbuckets/branches [GET]
func (b v1BitbucketApi) GetBranches(context echo.Context) error {
	repoId := context.QueryParam("repoId")
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
		LoadToken:        true,
	}
	id := context.QueryParam("companyId")
	repo := b.companyService.GetRepositoryByRepositoryId(id, repoId, option)
	url := context.QueryParam("url")
	username, repositoryName := getUsernameAndRepoNameFromBitbucketRepositoryUrl(url)
	branches, err := b.gitService.GetBranches(username, repositoryName, repo.Token)
	if err != nil {
		return common.GenerateErrorResponse(context, err, err.Error())
	}
	return common.GenerateSuccessResponse(context, branches, nil, "success")
}

// Listen ... Listen Bitbucket Web hook event
// @Summary  Listen Bitbucket Web hook event
// @Description Listens Bitbucket Web hook events. Register this endpoint as Bitbucket web hook endpoint
// @Tags Bitbucket
// @Accept json
// @Produce json
// @Param data body v1.BitbucketWebHookEvent true "GithubWebHookEvent Data"
// @Success 200 {object} common.ResponseDTO{data=string}
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/bitbuckets [POST]
func (b v1BitbucketApi) ListenEvent(context echo.Context) error {
	resource := new(v1.BitbucketWebHookEvent)
	if err := context.Bind(resource); err != nil {
		log.Println(err.Error())
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no companyId is provided", "Please provide companyId")
	}
	repoName := resource.Repository.Name
	owner := resource.Repository.Workspace.Slug
	revision := resource.Push.Changes[len(resource.Push.Changes)-1].New.Target.Hash
	repository := b.companyService.GetRepositoryByCompanyIdAndApplicationUrl(companyId, resource.Repository.Links.HTML.Href)
	application := b.companyService.GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repository.Id, resource.Repository.Links.HTML.Href)
	if !application.MetaData.IsWebhookEnabled {
		return common.GenerateForbiddenResponse(context, "[Forbidden]: Web hook is disabled!", "Operation Failed!")
	}
	data, err := b.gitService.GetPipeline(repoName, owner, revision, repository.Token)
	if err != nil {
		log.Println("[ERROR]:Failed to trigger pipeline process! ", err.Error())
		return common.GenerateErrorResponse(context, err.Error(), "Failed to trigger pipeline process!")
	}
	checkingFlag := BranchExists(data.Steps, resource.Push.Changes[len(resource.Push.Changes)-1].New.Name, "BIT_BUCKET")
	if !checkingFlag {
		return common.GenerateErrorResponse(context, "Branch does not exist!", "Operation Failed!")
	}
	if data != nil {
		stepsCount := len(data.Steps)
		for i := 0; i < stepsCount; i++ {
			if data.Steps[i].Type == enums.BUILD {
				if images, ok := data.Steps[i].Params["images"]; ok {
					data.Steps[i].Params["images"] = setImageVersionForBuild(data.Steps[i], revision, images)
				}
				if storage, ok := data.Steps[i].Params[enums.STORAGE]; ok {
					data.Steps[i].Params[enums.STORAGE] = storage
				} else {
					data.Steps[i].Params[enums.STORAGE] = "500Mi"
				}
				if accessMode, ok := data.Steps[i].Params[enums.ACCESS]; ok {
					data.Steps[i].Params[enums.ACCESS] = setAccessModeForBuild(accessMode)
				}
				if buildType, ok := data.Steps[i].Params[enums.BUILD_TYPE]; ok {
					data.Steps[i].Params[enums.BUILD_TYPE] = buildType
				}

			} else if data.Steps[i].Type == enums.DEPLOY {

				isThisStepValidForThisCommit := false
				if data.Steps[i].Params[enums.REVISION] != "" {
					allowedRevisions := strings.Split(data.Steps[i].Params[enums.REVISION], ",")
					branch := resource.Push.Changes[len(resource.Push.Changes)-1].New.Name
					for _, each := range allowedRevisions {
						if each == branch {
							isThisStepValidForThisCommit = true
							break
						}
					}
				}
				if isThisStepValidForThisCommit {
					data.Steps[i].Params["images"] = setDeploymentVersion(data.Steps[i], revision)
					descriptor := b.setDescriptors(data.Steps[i], repoName, owner, revision, repository.Token)
					if descriptor != nil {
						data.Steps[i].Descriptors = descriptor
					} else {
						return common.GenerateErrorResponse(context, err.Error(), "Failed to trigger pipeline process!")
					}
				} else {
					data.Steps = append(data.Steps[:i], data.Steps[i+1:]...)
					stepsCount = stepsCount - 1
					i = i - 1
				}
			} else if data.Steps[i].Type == enums.INTERMEDIARY {
				if images, ok := data.Steps[i].Params["images"]; ok {
					data.Steps[i].Params["images"] = setImageVersionForIntermediary(data.Steps[i], revision, images)
				}
			}
		}
	}
	data.ProcessId = uuid.NewV4().String()

	company, _ := b.companyService.GetByCompanyId(companyId, v1.CompanyQueryOption{})
	todaysRanProcess := b.processInventoryEventService.CountTodaysRanProcessByCompanyId(companyId)
	data.MetaData = v1.PipelineMetadata{
		CompanyId:       companyId,
		CompanyMetadata: company.MetaData,
		CommitId:        revision,
	}
	subject := v1.Subject{
		Log:                   "Pipeline triggered",
		CoreRequestQueryParam: map[string]string{"url": resource.Repository.Links.HTML.Href, "revision": revision, "purging": "ENABLE"},
		EventData:             map[string]interface{}{},
		Pipeline:              *data,
		App: struct {
			CompanyId    string
			AppId        string
			RepositoryId string
		}{
			CompanyId:    companyId,
			AppId:        application.MetaData.Id,
			RepositoryId: repository.Id,
		},
	}
	if todaysRanProcess >= company.MetaData.TotalProcessPerDay {
		subject.Log = "No More process today, you've touched today's limit!"
		if subject.EventData == nil {
			subject.EventData = make(map[string]interface{})
		}
		subject.EventData["trigger"] = false
		subject.EventData["log"] = subject.Log
	}

	go b.notifyAll(subject)
	return common.GenerateSuccessResponse(context, data.ProcessId, nil, "Pipeline triggered!")
}

// setDescriptors returns descriptors for deployment
func (b v1BitbucketApi) setDescriptors(step v1.Step, repoName string, owner string, revision string, token string) *[]unstructured.Unstructured {
	var descriptor *[]unstructured.Unstructured
	if val, ok := step.Params["env"]; ok {
		contentsData, err := b.gitService.GetDescriptors(repoName, owner, revision, token, enums.PIPELINE_DESCRIPTORS_BASE_DIRECTORY+"/", val)
		if err != nil {
			return nil
		}
		if contentsData != nil {
			descriptor = &contentsData
		}
	}
	return descriptor
}
func (b v1BitbucketApi) notifyAll(listener v1.Subject) {
	for _, observer := range b.observerList {
		go observer.Listen(listener)
	}
}

// NewBitbucketApi returns Git type api
func NewBitbucketApi(gitService service.Git, companyService service.Company, processInventoryEventService service.ProcessInventoryEvent, observerList []service.Observer) api.Git {
	return &v1BitbucketApi{
		gitService:                   gitService,
		companyService:               companyService,
		observerList:                 observerList,
		processInventoryEventService: processInventoryEventService,
	}
}
