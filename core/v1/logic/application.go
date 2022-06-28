package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"log"
)

type applicationService struct {
	repo                          repository.ApplicationRepository
	applicationMetadataRepository repository.ApplicationMetadataRepository
	client                        service.HttpClient
}

func (a applicationService) GetByCompanyIdAndRepoId(companyId, repoId string) []v1.Application {
	return a.repo.GetByCompanyIdAndRepoId(companyId, repoId)
}

func (a applicationService) StoreAll(applications []v1.Application) error {
	return a.repo.StoreAll(applications)
}

func (a applicationService) CreateWebHookAndUpdateApplications(repoType enums.REPOSITORY_TYPE, token string, apps []v1.Application) {
	if repoType == enums.GITHUB {
		for _, each := range apps {
			go a.CreateGithubWebHookAndStoreApplication(token, each)
		}
	} else if repoType == enums.BIT_BUCKET {
		for _, each := range apps {
			go a.CreateBitbucketWebHookAndStoreApplication(token, each)
		}
	}
}

func (a applicationService) CreateGithubWebHookAndStoreApplication(token string, app v1.Application) {
	usernameOrorgName, repoName := v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(app.Url)
	gitWebhook, err := NewGithubService(nil, a.client).CreateRepositoryWebhook(usernameOrorgName, repoName, token, app.CompanyId)
	if err != nil {
		log.Println("ERROR while creating webhook for application: ", err.Error())
	}
	app.Webhook = gitWebhook
	app.MetaData.IsWebhookEnabled = gitWebhook.Active
	app.Status = enums.ACTIVE
	applicationMetadataCollection := v1.ApplicationMetadataCollection{
		MetaData: app.MetaData,
		Status:   app.Status,
	}
	err = a.applicationMetadataRepository.Update(app.CompanyId, applicationMetadataCollection)
	if err != nil {
		return
	}
	err = a.repo.Store(app)
	if err != nil {
		return
	}
}

func (a applicationService) CreateBitbucketWebHookAndStoreApplication(token string, app v1.Application) {
	usernameOrorgName, repoName := v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(app.Url)
	gitWebhook, err := NewBitBucketService(nil, a.client).CreateRepositoryWebhook(usernameOrorgName, repoName, token, app.CompanyId)
	if err != nil {
		log.Println("ERROR while creating webhook for application: ", err.Error())
	}
	app.Webhook = gitWebhook
	app.MetaData.IsWebhookEnabled = gitWebhook.Active
	app.Status = enums.ACTIVE
	applicationMetadataCollection := v1.ApplicationMetadataCollection{
		MetaData: app.MetaData,
		Status:   app.Status,
	}
	err = a.applicationMetadataRepository.Update(app.CompanyId, applicationMetadataCollection)
	if err != nil {
		return
	}
	err = a.repo.Store(app)
	if err != nil {
		return
	}
}

// NewApplicationService returns Application type service
func NewApplicationService(repo repository.ApplicationRepository, applicationMetadataRepository repository.ApplicationMetadataRepository, client service.HttpClient) service.Application {
	return &applicationService{
		repo:                          repo,
		applicationMetadataRepository: applicationMetadataRepository,
		client:                        httpClientService{},
	}
}
