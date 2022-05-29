package logic

import (
	"errors"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

type pipelineService struct {
	githubService    service.Git
	bitbucketService service.Git
	companyService   service.Company
}

func (p pipelineService) GetPipelineForValidation(companyId, repositoryId, url, revision string) (v1.PipelineForValidation, error) {
	option := v1.CompanyQueryOption{LoadRepositories: true, LoadApplications: true, LoadToken: true}
	repo := p.companyService.GetRepositoryByRepositoryId(companyId, repositoryId, option)
	var username, repoName string
	var pipelineForValidation v1.PipelineForValidation
	if repo.Type == enums.GITHUB {
		username, repoName = v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(url)
		pipeline, err := p.githubService.GetPipeline(repoName, username, revision, repo.Token)
		if err != nil {
			return v1.PipelineForValidation{}, err
		}
		pipelineForValidation = pipeline.GetPipelineForValidationFromPipeline()
		return pipelineForValidation, nil
	} else if repo.Type == enums.BIT_BUCKET {
		username, repoName = v1.GetUsernameAndRepoNameFromBitbucketRepositoryUrl(url)
		pipeline, err := p.bitbucketService.GetPipeline(repoName, username, revision, repo.Token)
		if err != nil {
			return v1.PipelineForValidation{}, err
		}
		pipelineForValidation = pipeline.GetPipelineForValidationFromPipeline()
		return pipelineForValidation, nil
	}
	return v1.PipelineForValidation{}, errors.New("invalid repository type")
}

// NewPipelineService returns Pipeline type service
func NewPipelineService(githubService service.Git, bitbucketService service.Git, companyService service.Company) service.Pipeline {
	return &pipelineService{
		githubService:    githubService,
		bitbucketService: bitbucketService,
		companyService:   companyService,
	}
}
