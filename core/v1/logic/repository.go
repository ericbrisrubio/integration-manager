package logic

import (
	"errors"
	"github.com/google/uuid"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

type repositoryService struct {
	repo               repository.RepositoryRepository
	applicationService service.Application
}

func (r repositoryService) GetById(companyId, repositoryId string) v1.Repository {
	return r.repo.GetById(companyId, repositoryId)
}

func (r repositoryService) GetByCompanyId(companyId string, pagination bool, option v1.CompanyQueryOption) ([]v1.Repository, int64) {
	return r.repo.GetByCompanyId(companyId, pagination, option)
}

func (r repositoryService) Store(repositories []v1.Repository) error {
	return r.repo.Store(repositories)
}

func (r repositoryService) UpdateRepositories(companyId string, repositoriesDto []v1.RepositoryDto, repositoryUpdateOption v1.RepositoryUpdateOption) error {
	if repositoryUpdateOption.Option == enums.APPEND_REPOSITORY {
		return r.AppendRepositories(companyId, repositoriesDto)
	} else if repositoryUpdateOption.Option == enums.SOFT_DELETE_REPOSITORY {
		return r.SoftDeleteRepositories(companyId, repositoriesDto)
	} else if repositoryUpdateOption.Option == enums.DELETE_REPOSITORY {
		return r.DeleteRepositories(companyId, repositoriesDto)
	} else {
		return errors.New("invalid repository update option")
	}
}

func (r repositoryService) AppendRepositories(companyId string, repositoriesDto []v1.RepositoryDto) error {
	var repositories []v1.Repository
	for i, eachRepo := range repositoriesDto {
		repositoriesDto[i].Id = uuid.New().String()
		var applications []v1.Application
		for j, _ := range eachRepo.Applications {
			eachRepo.Applications[j].MetaData.Id = uuid.New().String()
			applications = append(applications, v1.Application{
				MetaData:       eachRepo.Applications[j].MetaData,
				RepositoryId:   eachRepo.Id,
				RepositoryType: eachRepo.Type,
				CompanyId:      companyId,
				Url:            eachRepo.Applications[j].Url,
				Webhook:        eachRepo.Applications[j].Webhook,
				Status:         eachRepo.Applications[j].Status,
			})
		}
		go r.applicationService.CreateWebHookAndUpdateApplications(eachRepo.Type, eachRepo.Token, applications)
		repositories = append(repositories, v1.Repository{
			Id:        eachRepo.Id,
			CompanyId: companyId,
			Type:      eachRepo.Type,
			Token:     eachRepo.Token,
		})
	}
	err := r.repo.Store(repositories)
	if err != nil {
		return err
	}
	return nil
}

func (r repositoryService) SoftDeleteRepositories(companyId string, repositoriesDto []v1.RepositoryDto) error {
	for _, eachRepo := range repositoriesDto {
		for _, eachApp := range eachRepo.Applications {
			application := v1.Application{
				MetaData:       eachApp.MetaData,
				RepositoryId:   eachRepo.Id,
				RepositoryType: eachRepo.Type,
				CompanyId:      companyId,
				Url:            eachApp.Url,
				Webhook:        eachApp.Webhook,
				Status:         enums.INACTIVE,
			}
			err := r.applicationService.SoftDeleteApplication(application)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r repositoryService) DeleteRepositories(companyId string, repositoriesDto []v1.RepositoryDto) error {
	for _, eachRepo := range repositoriesDto {
		for _, eachApp := range eachRepo.Applications {
			err := r.applicationService.DeleteApplication(companyId, eachRepo.Id, eachApp.MetaData.Id)
			if err != nil {
				return err
			}
		}
		err := r.repo.DeleteRepository(companyId, eachRepo.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewRepositoryService returns Repository type service
func NewRepositoryService(repo repository.RepositoryRepository, applicationService service.Application) service.Repository {
	return &repositoryService{
		repo:               repo,
		applicationService: applicationService,
	}
}
