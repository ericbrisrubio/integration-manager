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
	gitService service.Git
	companyService service.Company
}

func (g githubApi) ListenEvent(context echo.Context) error {

	resource := new(v1.GithubWebHook)
	if err := context.Bind(resource); err != nil {
	log.Println(err.Error())
	}
	repoName:=resource.Repository.Name
	owner:=resource.Repository.Owner.Login
	revision:=resource.After

	repository:=g.companyService.GetRepositoryByCompanyIdAndApplicationUrl(resource.Repository.Owner.Email,enums.GITHUB_BASE_URL+owner+"/"+repoName)

	data:=g.gitService.GetPipeline(repoName,owner,revision,repository.Token)

	if data!=nil {

	}
	return common.GenerateErrorResponse(context,nil,"Failed to trigger pipeline process!")
}

func NewGithubApi(gitService service.Git) api.Github {
	return &githubApi{
		gitService: gitService,
	}
}