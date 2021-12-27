package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	"github.com/klovercloud-ci-cd/integration-manager/config"
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

type v1GithubApi struct {
	gitService                   service.Git
	companyService               service.Company
	processInventoryEventService service.ProcessInventoryEvent
	observerList                 []service.Observer
}

// Listen ... Listen Github Web hook event
// @Summary  Listen Github Web hook event
// @Description Listens Github Web hook events. Register this endpoint as github web hook endpoint
// @Tags Github
// @Accept json
// @Produce json
// @Param data body v1.GithubWebHookEvent true "GithubWebHookEvent Data"
// @Success 200 {object} common.ResponseDTO{data=string}
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/githubs [POST]
func (g v1GithubApi) ListenEvent(context echo.Context) error {
	resource := new(v1.GithubWebHookEvent)
	if err := context.Bind(resource); err != nil {
		log.Println(err.Error())
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	companyId :=context.QueryParam("companyId")
	if companyId==""{
		return common.GenerateErrorResponse(context,"[ERROR] no companyId is provided","Please provide companyId")
	}
	repoName := resource.Repository.Name
	owner := resource.Repository.Owner.Login
	revision := resource.After
	repository := g.companyService.GetRepositoryByCompanyIdAndApplicationUrl(companyId, resource.Repository.URL)
	application := g.companyService.GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repository.Id, resource.Repository.URL)
	if !application.MetaData.IsWebhookEnabled {
		return common.GenerateForbiddenResponse(context, "[Forbidden]: Web hook is disabled!", "Operation Failed!")
	}
	data, err := g.gitService.GetPipeline(repoName, owner, revision, repository.Token)
	if err != nil {
		log.Println("[ERROR]:Failed to trigger pipeline process! ", err.Error())
		return common.GenerateErrorResponse(context, err.Error(), "Failed to trigger pipeline process!")
	}
	checkingFlag := branchExists(data.Steps, resource)
	if !checkingFlag {
		return common.GenerateErrorResponse(context, "Branch does not exist!", "Operation Failed!")
	}
	if data != nil {
		for i := range data.Steps {
			if data.Steps[i].Type == enums.BUILD {
				if images, ok := data.Steps[i].Params["images"]; ok {
					data.Steps[i].Params["images"] = setImageVersionForBuild(data.Steps[i], revision, images)
				}

			} else if data.Steps[i].Type == enums.DEPLOY {
				data.Steps[i].Params["images"] = setDeploymentVersion(data.Steps[i], revision)
				descriptor := g.setDescriptors(data.Steps[i], repoName, owner, revision, repository.Token)
				if descriptor != nil {
					data.Steps[i].Descriptors = descriptor
				} else {
					return common.GenerateErrorResponse(context, err.Error(), "Failed to trigger pipeline process!")
				}

			} else if data.Steps[i].Type == enums.INTERMEDIARY {
				if images, ok := data.Steps[i].Params["images"]; ok {
					data.Steps[i].Params["images"] = setImageVersionForIntermediary(data.Steps[i], revision, images)
				}
			}
		}
	}
	data.ProcessId = uuid.NewV4().String()

	company, _ := g.companyService.GetByCompanyId(companyId, v1.CompanyQueryOption{v1.Pagination{}, false, false})
	todaysRanProcess := g.processInventoryEventService.CountTodaysRanProcessByCompanyId(companyId)
	data.MetaData = v1.PipelineMetadata{
		CompanyId:       companyId,
		CompanyMetadata: company.MetaData,
	}
	subject := v1.Subject{
		Log:                   "Pipeline triggered",
		CoreRequestQueryParam: map[string]string{"url": resource.Repository.URL, "revision": revision, "purging": config.PipelinePurging},
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

	go g.notifyAll(subject)
	return common.GenerateSuccessResponse(context, data.ProcessId, nil, "Pipeline triggered!")
}

// setImageVersionForBuild returns image version for build step
func setImageVersionForBuild(step v1.Step, revision string, images string) string {
	imageRevision := revision
	if step.Params[enums.REVISION] != "" {
		imageRevision = step.Params[enums.REVISION]
	}
	listOfImages := strings.Split(images, ",")
	for i, image := range listOfImages {
		strs := strings.Split(image, ":")
		if len(strs) == 1 {
			listOfImages[i] = listOfImages[i] + ":" + imageRevision
		}
	}
	return strings.Join(listOfImages, ",")
}

// setImageVersionForIntermediary returns image version for Intermediary step
func setImageVersionForIntermediary(step v1.Step, revision string, img string) string {
	images := strings.Split(img, ",")
	imageRevision := revision
	if step.Params[enums.REVISION] != "" {
		imageRevision = step.Params[enums.REVISION]
	}
	for i, image := range images {
		strs := strings.Split(image, ":")
		if len(strs) == 1 {
			images[i] = images[i] + ":" + imageRevision
		}
	}
	return strings.Join(images, ",")
}

// branchExists returns boolean for branch existence
func branchExists(steps []v1.Step, resource *v1.GithubWebHookEvent) bool {
	for _, step := range steps {
		if step.Type == enums.BUILD && step.Params[enums.REVISION] != "" {
			branch := strings.Split(resource.Ref, "/")[2]
			if step.Params[enums.REVISION] != branch {
				log.Println("[Forbidden]: Branch wasn't matched!")
				return false
			}
		}
	}
	return true
}

// setDeploymentVersion returns image version for deployment
func setDeploymentVersion(step v1.Step, revision string) string {
	var deploymentVersion string
	if images, ok := step.Params["images"]; ok {
		images := strings.Split(images, ",")
		for i, image := range images {
			strs := strings.Split(image, ":")
			if len(strs) == 1 {
				images[i] = images[i] + ":" + revision
			}
		}
		deploymentVersion = strings.Join(images, ",")
	}
	return deploymentVersion
}

// setDescriptors returns descriptors for deployment
func (g v1GithubApi) setDescriptors(step v1.Step, repoName string, owner string, revision string, token string) *[]unstructured.Unstructured {
	var descriptor *[]unstructured.Unstructured
	if val, ok := step.Params["env"]; ok {
		contentsData, err := g.gitService.GetDescriptors(repoName, owner, revision, token, enums.PIPELINE_DESCRIPTORS_BASE_DIRECTORY+"/", val)
		if err != nil {
			return nil
		}
		if contentsData != nil {
			descriptor = &contentsData
		}
	}
	return descriptor
}
func (g v1GithubApi) notifyAll(listener v1.Subject) {
	for _, observer := range g.observerList {
		go observer.Listen(listener)
	}
}

// NewGithubApi returns Git type api
func NewGithubApi(gitService service.Git, companyService service.Company, processInventoryEventService service.ProcessInventoryEvent, observerList []service.Observer) api.Git {
	return &v1GithubApi{
		gitService:                   gitService,
		companyService:               companyService,
		observerList:                 observerList,
		processInventoryEventService: processInventoryEventService,
	}
}
