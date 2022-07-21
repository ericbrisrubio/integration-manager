package logic

import (
	"errors"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

type pipelineService struct {
	githubService     service.Git
	bitbucketService  service.Git
	repositoryService service.Repository
}

func (p pipelineService) Create(companyId, repositoryId, url string, payload v1.DirectoryContentCreatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error) {
	repo := p.repositoryService.GetById(companyId, repositoryId)
	if repo.Id == "" {
		return v1.DirectoryContentCreateAndUpdateResponse{}, errors.New("repository not found")
	}
	var username, repoName, path string
	path = "klovercloud/pipeline/pipeline.yml"
	if repo.Type == enums.GITHUB {
		username, repoName = v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(url)
		pipeline, err := p.githubService.CreateDirectoryContent(repoName, username, repo.Token, path, payload)
		if err != nil {
			return v1.DirectoryContentCreateAndUpdateResponse{}, err
		}
		return pipeline, nil
	} else if repo.Type == enums.BIT_BUCKET {
		username, repoName = v1.GetUsernameAndRepoNameFromBitbucketRepositoryUrl(url)
		pipeline, err := p.bitbucketService.CreateDirectoryContent(repoName, username, repo.Token, path, payload)
		if err != nil {
			return v1.DirectoryContentCreateAndUpdateResponse{}, err
		}
		return pipeline, nil
	}
	return v1.DirectoryContentCreateAndUpdateResponse{}, nil
}

func (p pipelineService) Update(companyId, repositoryId, url string, payload v1.DirectoryContentUpdatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error) {
	repo := p.repositoryService.GetById(companyId, repositoryId)
	if repo.Id == "" {
		return v1.DirectoryContentCreateAndUpdateResponse{}, errors.New("repository not found")
	}
	var username, repoName, path string
	path = "klovercloud/pipeline/pipeline.yml"
	if repo.Type == enums.GITHUB {
		username, repoName = v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(url)
		content, err := p.githubService.GetContent(repoName, username, repo.Token, path)
		if err != nil {
			return v1.DirectoryContentCreateAndUpdateResponse{}, err
		}
		payload.Sha = content.Sha
		pipeline, err := p.githubService.UpdateDirectoryContent(repoName, username, repo.Token, path, payload)
		if err != nil {
			return v1.DirectoryContentCreateAndUpdateResponse{}, err
		}
		return pipeline, nil
	} else if repo.Type == enums.BIT_BUCKET {
		username, repoName = v1.GetUsernameAndRepoNameFromBitbucketRepositoryUrl(url)
		content, err := p.bitbucketService.GetContent(repoName, username, repo.Token, path)
		if err != nil {
			return v1.DirectoryContentCreateAndUpdateResponse{}, err
		}
		payload.Sha = content.Sha
		pipeline, err := p.bitbucketService.UpdateDirectoryContent(repoName, username, repo.Token, path, payload)
		if err != nil {
			return v1.DirectoryContentCreateAndUpdateResponse{}, err
		}
		return pipeline, nil
	}
	return v1.DirectoryContentCreateAndUpdateResponse{}, nil
}

func (p pipelineService) GetPipelineForValidation(companyId, repositoryId, url, revision string) (v1.PipelineForValidation, error) {
	repo := p.repositoryService.GetById(companyId, repositoryId)
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
func NewPipelineService(githubService service.Git, bitbucketService service.Git, repositoryService service.Repository) service.Pipeline {
	return &pipelineService{
		githubService:     githubService,
		bitbucketService:  bitbucketService,
		repositoryService: repositoryService,
	}
}
