package v1

import (
	_ "encoding/json"
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/labstack/echo/v4"
	"github.com/twinj/uuid"
	"log"
	"strings"
)

type v1BitbucketApi struct {
	gitService                   service.Git
	companyService               service.Company
	processInventoryEventService service.ProcessInventoryEvent
	observerList                 []service.Observer
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
	repoName := resource.Repository.Name
	owner := resource.Actor.DisplayName
	revision := resource.Push.Changes[len(resource.Push.Changes)-1].New.Target.Hash
	companyId := resource.Repository.Owner.DisplayName
	repository := b.companyService.GetRepositoryByCompanyIdAndApplicationUrl(companyId, resource.Repository.Links.HTML.Href)
	application := b.companyService.GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, resource.Repository.UUID, resource.Repository.Links.HTML.Href)
	if !application.MetaData.IsWebhookEnabled {
		return common.GenerateForbiddenResponse(context, "[Forbidden]: Web hook is disabled!", "Operation Failed!")
	}
	data, err := b.gitService.GetPipeline(repoName, owner, revision, repository.Token)
	if err != nil {
		log.Println("[ERROR]:Failed to trigger pipeline process! ", err.Error())
		return common.GenerateErrorResponse(context, err.Error(), "Failed to trigger pipeline process!")
	}
	for _, step := range data.Steps {
		if step.Type == enums.BUILD && step.Params[enums.REVISION] != "" {
			branch := resource.Push.Changes[len(resource.Push.Changes)-1].New.Name
			if step.Params[enums.REVISION] != branch {
				return common.GenerateForbiddenResponse(context, "[Forbidden]: Branch wasn't matched!", "Operation Failed!")
			}
		}
	}
	if data != nil {
		for i := range data.Steps {
			if data.Steps[i].Type == enums.DEPLOY {
				if val, ok := data.Steps[i].Params["env"]; ok {
					contentsData, err := b.gitService.GetDescriptors(repoName, owner, revision, repository.Token, enums.PIPELINE_DESCRIPTORS_BASE_DIRECTORY+"/", val)
					if err != nil {
						return common.GenerateErrorResponse(context, err.Error(), "Failed to trigger pipeline process!")
					}
					if contentsData != nil {
						data.Steps[i].Descriptors = &contentsData
					}
				}
			} else if data.Steps[i].Type == enums.BUILD {
				if images, ok := data.Steps[i].Params["images"]; ok {
					imageRevision := revision
					if data.Steps[i].Params[enums.REVISION] != "" {
						imageRevision = data.Steps[i].Params[enums.REVISION]
					}
					images := strings.Split(images, ",")
					for i, image := range images {
						strs := strings.Split(image, ":")
						if len(strs) == 1 {
							images[i] = images[i] + ":" + imageRevision
						}
					}
					data.Steps[i].Params["images"] = strings.Join(images, ",")
				}

			} else if data.Steps[i].Type == enums.INTERMEDIARY {
				if images, ok := data.Steps[i].Params["images"]; ok {
					images := strings.Split(images, ",")
					imageRevision := revision
					if data.Steps[i].Params[enums.REVISION] != "" {
						imageRevision = data.Steps[i].Params[enums.REVISION]
					}
					for i, image := range images {
						strs := strings.Split(image, ":")
						if len(strs) == 1 {
							images[i] = images[i] + ":" + imageRevision
						}
					}
					data.Steps[i].Params["images"] = strings.Join(images, ",")
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
