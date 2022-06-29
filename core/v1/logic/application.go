package logic

import (
	"errors"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"k8s.io/apimachinery/pkg/util/yaml"
	"log"
)

type applicationService struct {
	repo                       repository.ApplicationRepository
	applicationMetadataService service.ApplicationMetadataService
	client                     service.HttpClient
}

func (a applicationService) GetByCompanyIdAndRepoId(companyId, repoId string, pagination bool, option v1.CompanyQueryOption, statusQuery bool, status v1.StatusQueryOption) ([]v1.Application, int64) {
	return a.repo.GetByCompanyIdAndRepoId(companyId, repoId, pagination, option, statusQuery, status)
}

func (a applicationService) GetApplicationsByCompanyIdAndRepositoryType(companyId string, _type enums.REPOSITORY_TYPE, pagination bool, option v1.CompanyQueryOption, statusQuery bool, status v1.StatusQueryOption) ([]v1.Application, int64) {
	return a.repo.GetApplicationsByCompanyIdAndRepositoryType(companyId, _type, pagination, option, statusQuery, status)
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
	err = a.applicationMetadataService.Update(app.CompanyId, applicationMetadataCollection)
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
	err = a.applicationMetadataService.Update(app.CompanyId, applicationMetadataCollection)
	if err != nil {
		return
	}
	err = a.repo.Store(app)
	if err != nil {
		return
	}
}

func (a applicationService) webHookForGithub(apps []v1.ApplicationDto, companyId string, token string) {
	for i := range apps {
		usernameOrorgName, repoName := v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(apps[i].Url)
		gitWebhook, err := NewGithubService(nil, a.client).CreateRepositoryWebhook(usernameOrorgName, repoName, token, companyId)
		if err != nil {
			log.Println("[ERROR] Failed to create webhook ", err.Error())
		}
		apps[i].Webhook = gitWebhook
		apps[i].MetaData.IsWebhookEnabled = gitWebhook.Active
		apps[i].Status = enums.ACTIVE
		applicationMetadataCollection := v1.ApplicationMetadataCollection{
			MetaData: apps[i].MetaData,
			Status:   apps[i].Status,
		}
		err = a.applicationMetadataService.Store(applicationMetadataCollection)
		if err != nil {
			return
		}
	}
}

func (a applicationService) webHookForBitbucket(apps []v1.ApplicationDto, companyId string, token string) {
	for i := range apps {
		usernameOrOrgName, repoName := v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(apps[i].Url)
		b, err := a.client.Get(enums.BITBUCKET_API_BASE_URL+"repositories/"+usernameOrOrgName+"/"+repoName, nil)
		var repositoryDetails v1.BitbucketRepository
		err = yaml.Unmarshal(b, &repositoryDetails)
		if err != nil {
			log.Println(err.Error())
		}
		bitbucketWebhook, err := NewBitBucketService(nil, a.client).CreateRepositoryWebhook(repositoryDetails.Workspace.Slug, repositoryDetails.Slug, token, companyId)
		if err != nil {
			log.Println("ERROR failed to create webhook", err.Error())
		}
		apps[i].Webhook = bitbucketWebhook
		apps[i].MetaData.IsWebhookEnabled = bitbucketWebhook.Active
		apps[i].Status = enums.ACTIVE
		applicationMetadataCollection := v1.ApplicationMetadataCollection{
			MetaData: apps[i].MetaData,
			Status:   apps[i].Status,
		}
		err = a.applicationMetadataService.Store(applicationMetadataCollection)
		if err != nil {
			return
		}
	}
}

func (a applicationService) UpdateWebhook(repository v1.Repository, url, webhookId, action string) error {
	if action == string(enums.WEBHOOK_EANBLE) {
		if repository.Type == enums.GITHUB {
			return a.EnableGithubWebhookAndUpdateApplication(repository.CompanyId, repository.Id, url, repository.Token)
		} else if repository.Type == enums.BIT_BUCKET {
			return a.EnableBitbucketWebhookAndUpdateApplication(repository.CompanyId, repository.Id, url, repository.Token)
		}
	} else if action == string(enums.WEBHOOK_DISABLE) {
		if repository.Type == enums.GITHUB {
			return a.DisableGithubWebhookAndUpdateApplication(repository.CompanyId, repository.Id, url, webhookId, repository.Token)
		} else if repository.Type == enums.BIT_BUCKET {
			return a.DisableBitbucketWebhookAndUpdateApplication(repository.CompanyId, repository.Id, url, webhookId, repository.Token)
		}
	}
	return errors.New("provide valid action")
}

func (a applicationService) EnableBitbucketWebhookAndUpdateApplication(companyId, repoId, url, token string) error {
	username, repositoryName := v1.GetUsernameAndRepoNameFromBitbucketRepositoryUrl(url)
	webhook, err := NewBitBucketService(nil, a.client).CreateRepositoryWebhook(username, repositoryName, token, companyId)
	if err != nil {
		return err
	}
	app := a.repo.GetByCompanyIdAndRepositoryIdAndUrl(companyId, repoId, url)
	app.Webhook = webhook
	app.MetaData.IsWebhookEnabled = true
	err = a.repo.Update(companyId, repoId, app)
	if err != nil {
		return err
	}
	applicationMetadata := v1.ApplicationMetadataCollection{
		MetaData: app.MetaData,
		Status:   app.Status,
	}
	err = a.applicationMetadataService.Update(companyId, applicationMetadata)
	if err != nil {
		return err
	}
	return nil
}

func (a applicationService) DisableBitbucketWebhookAndUpdateApplication(companyId, repoId, url, webhookId, token string) error {
	username, repositoryName := v1.GetUsernameAndRepoNameFromBitbucketRepositoryUrl(url)
	err := NewBitBucketService(nil, a.client).DeleteRepositoryWebhookById(username, repositoryName, webhookId, token)
	if err != nil {
		return err
	}
	app := a.repo.GetByCompanyIdAndRepositoryIdAndUrl(companyId, repoId, url)
	app.Webhook = v1.GitWebhook{}
	app.MetaData.IsWebhookEnabled = false
	err = a.repo.Update(companyId, repoId, app)
	if err != nil {
		return err
	}
	applicationMetadata := v1.ApplicationMetadataCollection{
		MetaData: app.MetaData,
		Status:   app.Status,
	}
	err = a.applicationMetadataService.Update(companyId, applicationMetadata)
	if err != nil {
		return err
	}
	return nil
}

func (a applicationService) EnableGithubWebhookAndUpdateApplication(companyId, repoId, url, token string) error {
	username, repositoryName := v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(url)
	webhook, err := NewGithubService(nil, a.client).CreateRepositoryWebhook(username, repositoryName, token, companyId)
	if err != nil {
		return err
	}
	app := a.repo.GetByCompanyIdAndRepositoryIdAndUrl(companyId, repoId, url)
	app.Webhook = webhook
	app.MetaData.IsWebhookEnabled = true
	err = a.repo.Update(companyId, repoId, app)
	if err != nil {
		return err
	}
	applicationMetadata := v1.ApplicationMetadataCollection{
		MetaData: app.MetaData,
		Status:   app.Status,
	}
	err = a.applicationMetadataService.Update(companyId, applicationMetadata)
	if err != nil {
		return err
	}
	return nil
}

func (a applicationService) DisableGithubWebhookAndUpdateApplication(companyId, repoId, url, webhookId, token string) error {
	username, repositoryName := v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(url)
	err := NewGithubService(nil, a.client).DeleteRepositoryWebhookById(username, repositoryName, webhookId, token)
	if err != nil {
		return err
	}
	app := a.repo.GetByCompanyIdAndRepositoryIdAndUrl(companyId, repoId, url)
	app.Webhook = v1.GitWebhook{}
	app.MetaData.IsWebhookEnabled = false
	err = a.repo.Update(companyId, repoId, app)
	if err != nil {
		return err
	}
	applicationMetadata := v1.ApplicationMetadataCollection{
		MetaData: app.MetaData,
		Status:   app.Status,
	}
	err = a.applicationMetadataService.Update(companyId, applicationMetadata)
	if err != nil {
		return err
	}
	return nil
}

func (a applicationService) SoftDeleteApplication(application v1.Application) error {
	return a.repo.SoftDeleteApplication(application)
}

func (a applicationService) DeleteApplication(companyId, repositoryId, applicationId string) error {
	return a.repo.DeleteApplication(companyId, repositoryId, applicationId)
}

// NewApplicationService returns Application type service
func NewApplicationService(repo repository.ApplicationRepository, applicationMetadataService service.ApplicationMetadataService, client service.HttpClient) service.Application {
	return &applicationService{
		repo:                       repo,
		applicationMetadataService: applicationMetadataService,
		client:                     httpClientService{},
	}
}
