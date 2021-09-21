package v1

import (
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"log"
)

type githubApi struct {
	gitService     service.Git
	companyService service.Company
	observerList []service.Observer
}

func (g githubApi) ListenEvent(context echo.Context) error {
	resource := new(v1.GithubWebHook)
	if err := context.Bind(resource); err != nil {
		log.Println(err.Error())
	}
	repoName := resource.Repository.Name
	owner := resource.Repository.Owner.Login
	revision := resource.After
	companyId := ""
	if resource.Repository.Owner.Type == "Organization" {
		companyId = resource.Repository.Name
	} else {
		companyId = resource.Repository.Owner.Email
	}
	repository := g.companyService.GetRepositoryByCompanyIdAndApplicationUrl(companyId, enums.GITHUB_BASE_URL+owner+"/"+repoName)

	data := g.gitService.GetPipeline(repoName, owner, revision, repository.Token)

	if data != nil {
		for i,_:=range data.Steps{
			for j,_:=range data.Steps[i].Outputs{
				if (data.Steps[i].Outputs[j].DeploymentResource!=nil){
					if(data.Steps[i].Outputs[j].DeploymentResource.MountPath!=nil){
						//read files from mount path and set to descriptors
					}
				}
			}
		}
	}
	return common.GenerateErrorResponse(context, nil, "Failed to trigger pipeline process!")
}
func (g githubApi)notifyAll(listener v1.Subject){
	for _, observer := range g.observerList {
		go observer.Listen(listener)
	}
}
func NewGithubApi(gitService service.Git, companyService service.Company,observerList []service.Observer) api.Github {
	return &githubApi{
		gitService:     gitService,
		companyService: companyService,
		observerList: observerList,
	}
}
