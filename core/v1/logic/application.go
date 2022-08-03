package logic

import (
	"errors"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"log"
)

type applicationService struct {
	repo   repository.ApplicationRepository
	client service.HttpClient
}

func (a applicationService) SearchAppsByCompanyIdAndName(companyId, name string) []v1.Application {
	return a.repo.SearchAppsByCompanyIdAndName(companyId, name)
}

func (a applicationService) GetByCompanyIdAndUrl(companyId, applicationUrl string) v1.Application {
	return a.repo.GetByCompanyIdAndUrl(companyId, applicationUrl)
}

func (a applicationService) GetByCompanyIdAndRepositoryIdAndUrl(companyId, repositoryId, applicationUrl string) v1.Application {
	return a.repo.GetByCompanyIdAndRepositoryIdAndUrl(companyId, repositoryId, applicationUrl)
}

func (a applicationService) GetById(companyId string, repoId string, applicationId string) v1.Application {
	return a.repo.GetById(companyId, repoId, applicationId)
}

func (a applicationService) GetAll(companyId string, option v1.CompanyQueryOption) ([]v1.Application, int64) {
	return a.repo.GetAll(companyId, option)
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

func (a applicationService) UpdateApplications(repository v1.Repository, apps []v1.Application, applicationUpdateOption v1.ApplicationUpdateOption) error {
	if applicationUpdateOption.Option == enums.APPEND_APPLICATION {
		return a.AppendApplications(repository, apps)
	} else if applicationUpdateOption.Option == enums.SOFT_DELETE_APPLICATION {
		return a.SoftDeleteApplications(repository, apps)
	} else if applicationUpdateOption.Option == enums.DELETE_APPLICATION {
		return a.DeleteApplications(repository, apps)
	} else {
		return errors.New("invalid application update option")
	}
}

func (a applicationService) AppendApplications(repository v1.Repository, apps []v1.Application) error {
	a.CreateWebHookAndUpdateApplications(repository.Type, repository.Token, apps)
	return nil
}

func (a applicationService) SoftDeleteApplications(repository v1.Repository, apps []v1.Application) error {
	var applications []v1.Application
	for _, each := range apps {
		application := a.repo.GetById(repository.CompanyId, repository.Id, each.MetaData.Id)
		applications = append(applications, application)
	}
	for _, each := range applications {
		err := a.repo.SoftDeleteApplication(each)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a applicationService) DeleteApplications(repository v1.Repository, apps []v1.Application) error {
	var applications []v1.Application
	for _, each := range apps {
		application := a.repo.GetById(repository.CompanyId, repository.Id, each.MetaData.Id)
		applications = append(applications, application)
	}
	for _, each := range applications {
		err := a.repo.DeleteApplication(repository.CompanyId, repository.Id, each.MetaData.Id)
		if err != nil {
			return err
		}
	}
	return nil
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
	err = a.repo.Store(app)
	if err != nil {
		return
	}
}

func (a applicationService) CreateBitbucketWebHookAndStoreApplication(token string, app v1.Application) {
	usernameOrorgName, repoName := v1.GetUsernameAndRepoNameFromBitbucketRepositoryUrl(app.Url)
	gitWebhook, err := NewBitBucketService(nil, a.client).CreateRepositoryWebhook(usernameOrorgName, repoName, token, app.CompanyId)
	if err != nil {
		log.Println("ERROR while creating webhook for application: ", err.Error())
	}
	app.Webhook = gitWebhook
	app.MetaData.IsWebhookEnabled = gitWebhook.Active
	app.Status = enums.ACTIVE
	err = a.repo.Store(app)
	if err != nil {
		return
	}
}

func (a applicationService) UpdateWebhook(repository v1.Repository, url, webhookId, action, appId string) error {
	if action == string(enums.WEBHOOK_EANBLE) {
		if repository.Type == enums.GITHUB {
			return a.EnableGithubWebhookAndUpdateApplication(repository.CompanyId, repository.Id, appId, url, repository.Token)
		} else if repository.Type == enums.BIT_BUCKET {
			return a.EnableBitbucketWebhookAndUpdateApplication(repository.CompanyId, repository.Id, appId, url, repository.Token)
		}
	} else if action == string(enums.WEBHOOK_DISABLE) {
		if repository.Type == enums.GITHUB {
			return a.DisableGithubWebhookAndUpdateApplication(repository.CompanyId, repository.Id, appId, url, webhookId, repository.Token)
		} else if repository.Type == enums.BIT_BUCKET {
			return a.DisableBitbucketWebhookAndUpdateApplication(repository.CompanyId, repository.Id, appId, url, webhookId, repository.Token)
		}
	}
	return errors.New("provide valid action")
}

func (a applicationService) EnableBitbucketWebhookAndUpdateApplication(companyId, repoId, appId, url, token string) error {
	username, repositoryName := v1.GetUsernameAndRepoNameFromBitbucketRepositoryUrl(url)
	webhook, err := NewBitBucketService(nil, a.client).CreateRepositoryWebhook(username, repositoryName, token, companyId)
	if err != nil {
		return err
	}
	app := a.repo.GetById(companyId, repoId, appId)
	app.Webhook = webhook
	app.MetaData.IsWebhookEnabled = true
	err = a.repo.Update(companyId, repoId, app)
	if err != nil {
		return err
	}
	return nil
}

func (a applicationService) DisableBitbucketWebhookAndUpdateApplication(companyId, repoId, appId, url, webhookId, token string) error {
	username, repositoryName := v1.GetUsernameAndRepoNameFromBitbucketRepositoryUrl(url)
	err := NewBitBucketService(nil, a.client).DeleteRepositoryWebhookById(username, repositoryName, webhookId, token)
	if err != nil {
		return err
	}
	app := a.repo.GetById(companyId, repoId, appId)
	app.Webhook = v1.GitWebhook{}
	app.MetaData.IsWebhookEnabled = false
	err = a.repo.Update(companyId, repoId, app)
	if err != nil {
		return err
	}
	return nil
}

func (a applicationService) EnableGithubWebhookAndUpdateApplication(companyId, repoId, appId, url, token string) error {
	username, repositoryName := v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(url)
	webhook, err := NewGithubService(nil, a.client).CreateRepositoryWebhook(username, repositoryName, token, companyId)
	if err != nil {
		return err
	}
	app := a.repo.GetById(companyId, repoId, appId)
	app.Webhook = webhook
	app.MetaData.IsWebhookEnabled = true
	err = a.repo.Update(companyId, repoId, app)
	if err != nil {
		return err
	}
	return nil
}

func (a applicationService) DisableGithubWebhookAndUpdateApplication(companyId, repoId, appId, url, webhookId, token string) error {
	username, repositoryName := v1.GetUsernameAndRepoNameFromGithubRepositoryUrl(url)
	err := NewGithubService(nil, a.client).DeleteRepositoryWebhookById(username, repositoryName, webhookId, token)
	if err != nil {
		return err
	}
	app := a.repo.GetById(companyId, repoId, appId)
	app.Webhook = v1.GitWebhook{}
	app.MetaData.IsWebhookEnabled = false
	err = a.repo.Update(companyId, repoId, app)
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
func NewApplicationService(repo repository.ApplicationRepository, client service.HttpClient) service.Application {
	return &applicationService{
		repo:   repo,
		client: client,
	}
}
