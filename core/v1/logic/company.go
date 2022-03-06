package logic

import (
	"errors"
	"github.com/google/uuid"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"k8s.io/apimachinery/pkg/util/yaml"
	"log"
	"strings"
)

type companyService struct {
	repo   repository.CompanyRepository
	client service.HttpClient
}

func (c companyService) GetApplicationsByRepositoryId(repositoryId string, companyId string, option v1.CompanyQueryOption, status v1.StatusQueryOption) ([]v1.Application, int64) {
	return c.repo.GetApplicationsByRepositoryId(repositoryId, companyId, option, status)
}

func (c companyService) GetApplicationByApplicationId(companyId string, repoId string, applicationId string) v1.Application {
	return c.repo.GetApplicationByApplicationId(companyId, repoId, applicationId)
}

func (c companyService) GetRepositoryByRepositoryId(id string, companyId string, option v1.CompanyQueryOption) v1.Repository {
	return c.repo.GetRepositoryByRepositoryId(id, companyId, option)
}

func (c companyService) CreateWebHookAndUpdateApplications(companyId string, repoType enums.REPOSITORY_TYPE, repoId string, token string, apps []v1.Application) {
	if repoType == enums.GITHUB {
		for _, each := range apps {
			go c.CreateGithubWebHookAndUpdateApplication(companyId, repoId, token, each)
		}
	} else if repoType == enums.BIT_BUCKET {
		for _, each := range apps {
			go c.CreateBitbucketWebHookAndUpdateApplication(companyId, repoId, token, each)
		}
	}
}

func (c companyService) CreateGithubWebHookAndUpdateApplication(companyId string, repoId string, token string, app v1.Application) {
	usernameOrorgName, repoName := getUsernameAndRepoNameFromGithubRepositoryUrl(app.Url)
	gitWebhook, err := NewGithubService(c, nil, c.client).CreateRepositoryWebhook(usernameOrorgName, repoName, token, companyId)
	if err != nil {
		app.Webhook = gitWebhook
		app.MetaData.IsWebhookEnabled = false
	} else {
		app.Webhook = gitWebhook
		app.MetaData.IsWebhookEnabled = true
	}
	err = c.repo.UpdateApplication(companyId, repoId, app.MetaData.Id, app)
	if err != nil {
		return
	}
}

func (c companyService) CreateBitbucketWebHookAndUpdateApplication(companyId string, repoId string, token string, app v1.Application) {
	usernameOrorgName, repoName := getUsernameAndRepoNameFromGithubRepositoryUrl(app.Url)
	gitWebhook, err := NewBitBucketService(c, nil, c.client).CreateRepositoryWebhook(usernameOrorgName, repoName, token, companyId)
	if err != nil {
		app.Webhook = gitWebhook
		app.MetaData.IsWebhookEnabled = false
	} else {
		app.Webhook = gitWebhook
		app.MetaData.IsWebhookEnabled = true
	}
	err = c.repo.UpdateApplication(companyId, repoId, app.MetaData.Id, app)
	if err != nil {
		return
	}
}

func (c companyService) GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application {
	return v1.Application{
		MetaData: v1.ApplicationMetadata{
			Id:               "1001",
			IsWebhookEnabled: true,
		},
	}
	//return c.repo.GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl)
}
func (c companyService) UpdateRepositories(companyId string, repositories []v1.Repository, companyUpdateOption v1.RepositoryUpdateOption) error {
	if companyUpdateOption.Option == enums.APPEND_REPOSITORY {
		for i, each := range repositories {
			repositories[i].Id = uuid.New().String()
			for j := range each.Applications {
				each.Applications[j].MetaData.Id = uuid.New().String()
			}
			if each.Type == enums.GITHUB {
				c.webHookForGithub(each.Applications, companyId, each.Token)
			} else if each.Type == enums.BIT_BUCKET {
				c.webHookForBitbucket(each.Applications, companyId, each.Token)
			}
		}
		err := c.repo.AppendRepositories(companyId, repositories)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if companyUpdateOption.Option == enums.SOFT_DELETE_REPOSITORY {
		err := c.repo.DeleteRepositories(companyId, repositories, true)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if companyUpdateOption.Option == enums.DELETE_REPOSITORY {
		err := c.repo.DeleteRepositories(companyId, repositories, false)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func getUsernameAndRepoNameFromGithubRepositoryUrl(url string) (username string, repoName string) {
	trim := strings.TrimSuffix(url, ".git")
	urlArray := strings.Split(trim, "/")
	if len(urlArray) < 3 {
		return "", ""
	}
	repositoryName := urlArray[len(urlArray)-1]
	usernameOrorgName := urlArray[len(urlArray)-2]
	return usernameOrorgName, repositoryName
}

func getUsernameAndRepoNameFromBitbucketRepositoryUrl(url string) (username string, repoName string) {
	trim := strings.TrimSuffix(url, ".git")
	urlArray := strings.Split(trim, "/")
	if len(urlArray) < 3 {
		return "", ""
	}
	repositoryName := urlArray[len(urlArray)-4]
	usernameOrOrgName := urlArray[len(urlArray)-5]
	return usernameOrOrgName, repositoryName
}
func (c companyService) webHookForGithub(apps []v1.Application, companyId string, token string) {
	for i := range apps {
		usernameOrorgName, repoName := getUsernameAndRepoNameFromGithubRepositoryUrl(apps[i].Url)
		gitWebhook, err := NewGithubService(c, nil, c.client).CreateRepositoryWebhook(usernameOrorgName, repoName, token, companyId)
		if err != nil {
			apps[i].Webhook = gitWebhook
			apps[i].MetaData.IsWebhookEnabled = false
		} else {
			apps[i].Webhook = gitWebhook
			apps[i].MetaData.IsWebhookEnabled = true
			apps[i].Status = enums.ACTIVE
		}

	}
}

func (c companyService) webHookForBitbucket(apps []v1.Application, companyId string, token string) {
	for i := range apps {
		usernameOrOrgName, repoName := getUsernameAndRepoNameFromBitbucketRepositoryUrl(apps[i].Url)
		b, err := c.client.Get(enums.BITBUCKET_API_BASE_URL+"repositories/"+usernameOrOrgName+"/"+repoName, nil)
		var repositoryDetails v1.BitbucketRepository
		err = yaml.Unmarshal(b, &repositoryDetails)
		if err != nil {
			log.Println(err.Error())
		}
		bitbucketWebhook, err := NewBitBucketService(c, nil, c.client).CreateRepositoryWebhook(repositoryDetails.Workspace.Slug, repositoryDetails.Slug, token, companyId)
		if err != nil {
			apps[i].Webhook = bitbucketWebhook
			apps[i].MetaData.IsWebhookEnabled = false
		} else {
			apps[i].Webhook = bitbucketWebhook
			apps[i].MetaData.IsWebhookEnabled = true
			apps[i].Status = enums.ACTIVE
		}

	}
}

func (c companyService) UpdateApplications(companyId string, repositoryId string, apps []v1.Application, companyUpdateOption v1.ApplicationUpdateOption) error {
	option := v1.CompanyQueryOption{LoadApplications: true, LoadToken: true}
	if companyUpdateOption.Option == enums.APPEND_APPLICATION {
		for i := range apps {
			apps[i].MetaData.Id = uuid.New().String()
		}
		repo := c.GetRepositoryByRepositoryId(repositoryId, companyId, option)
		if repo.Type == enums.GITHUB {
			c.webHookForGithub(apps, companyId, repo.Token)
		} else if repo.Type == enums.BIT_BUCKET {
			c.webHookForBitbucket(apps, companyId, repo.Token)
		}
		err := c.repo.AppendApplications(companyId, repositoryId, apps)
		if err != nil {
			return err
		}
	}
	if companyUpdateOption.Option == enums.SOFT_DELETE_APPLICATION {
		err := c.repo.DeleteApplications(companyId, repositoryId, apps, true)
		if err != nil {
			return err
		}
	}
	if companyUpdateOption.Option == enums.DELETE_APPLICATION {
		repo := c.GetRepositoryByRepositoryId(repositoryId, companyId, option)
		if repo.Type == enums.GITHUB {
			for i := range apps {
				usernameOrorgName, repoName := getUsernameAndRepoNameFromGithubRepositoryUrl(apps[i].Url)
				err := NewGithubService(c, nil, c.client).DeleteRepositoryWebhookById(usernameOrorgName, repoName, apps[i].Webhook.ID, repo.Token)
				if err != nil {
					return err
				}
				apps[i].MetaData.IsWebhookEnabled = false
			}
		} else if repo.Type == enums.BIT_BUCKET {
			for i := range apps {
				usernameOrOrgName, repoName := getUsernameAndRepoNameFromBitbucketRepositoryUrl(apps[i].Url)
				err := NewBitBucketService(c, nil, c.client).DeleteRepositoryWebhookById(usernameOrOrgName, repoName, apps[i].Webhook.ID, repo.Token)
				if err != nil {
					return err
				}
				apps[i].MetaData.IsWebhookEnabled = false
			}
		}
		err := c.repo.DeleteApplications(companyId, repositoryId, apps, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c companyService) GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository {
	return v1.Repository{
		Id:    "1",
		Token: "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",
	}
	//return c.repo.GetRepositoryByCompanyIdAndApplicationUrl(id, url)
}

func (c companyService) GetCompanyByApplicationUrl(url string) v1.Company {
	return c.repo.GetCompanyByApplicationUrl(url)
}

func (c companyService) Store(company v1.Company) error {
	option := v1.CompanyQueryOption{}
	if company.Id == "" {
		return errors.New("[ERROR]: No company id given")
	}
	if data, _ := c.GetByCompanyId(company.Id, option); data.Id == company.Id {
		return errors.New("[ERROR]: Company with id: " + company.Id + " already exists.")
	}
	for _, eachRepo := range company.Repositories {
		go c.CreateWebHookAndUpdateApplications(company.Id, eachRepo.Type, eachRepo.Id, eachRepo.Token, eachRepo.Applications)
	}
	return c.repo.Store(company)
}

func (c companyService) Delete(companyId string) error {
	err := c.repo.Delete(companyId)
	if err != nil {
		return err
	}
	return nil
}

func (c companyService) GetCompanies(option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Company {
	companies, _ := c.repo.GetCompanies(option, status)
	return companies
}

func (c companyService) GetByCompanyId(id string, option v1.CompanyQueryOption) (v1.Company, int64) {
	company:=v1.Company{MetaData: struct {
		Labels                    map[string]string `bson:"labels" json:"labels" yaml:"labels"`
		NumberOfConcurrentProcess int64             `bson:"number_of_concurrent_process" json:"number_of_concurrent_process" yaml:"number_of_concurrent_process"`
		TotalProcessPerDay        int64             `bson:"total_process_per_day" json:"total_process_per_day" yaml:"total_process_per_day"`
	}{Labels: nil, NumberOfConcurrentProcess: 10, TotalProcessPerDay: 10}}
	var total int64
	//company, total := c.repo.GetByCompanyId(id, option)
	return company, total
	//return c.repo.GetByCompanyId(id, option)
}

func (c companyService) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64) {
	return c.repo.GetRepositoriesByCompanyId(id, option)
}

func (c companyService) GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Application {
	return c.repo.GetApplicationsByCompanyIdAndRepositoryType(id, _type, option, status)
}

// NewCompanyService returns Company type service
func NewCompanyService(repo repository.CompanyRepository, client service.HttpClient) service.Company {
	return &companyService{
		repo:   repo,
		client: client,
	}
}
