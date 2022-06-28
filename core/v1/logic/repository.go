package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
)

type repositoryService struct {
	repo repository.RepositoryRepository
}

func (r repositoryService) GetByCompanyId(companyId string) []v1.Repository {
	return r.repo.GetByCompanyId(companyId)
}

func (r repositoryService) Store(repositories []v1.Repository) error {
	return r.repo.Store(repositories)
}

// NewRepositoryService returns Repository type service
func NewRepositoryService(repo repository.RepositoryRepository) service.Repository {
	return &repositoryService{
		repo: repo,
	}
}
