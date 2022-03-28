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
	repo                          repository.CompanyRepository
	applicationMetadataRepository repository.ApplicationMetadataRepository
	client                        service.HttpClient
}

func (c companyService) GetAllApplications(companyId string, option v1.CompanyQueryOption) ([]v1.Application, int64) {
	repositories, _ := c.GetRepositoriesByCompanyId(companyId, option)
	var applications []v1.Application
	for _, eachRepo := range repositories {
		apps := eachRepo.Applications
		applications = append(applications, apps...)
	}
	return applications, int64(len(applications))
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
		log.Println("ERROR while creating webhook for application: ", err.Error())
	}
	app.Webhook = gitWebhook
	app.MetaData.IsWebhookEnabled = gitWebhook.Active
	app.Status = enums.ACTIVE
	applicationMetadataCollection := v1.ApplicationMetadataCollection{
		MetaData: app.MetaData,
		Status:   app.Status,
	}
	err = c.applicationMetadataRepository.Update(companyId, applicationMetadataCollection)
	if err != nil {
		return
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
		log.Println("ERROR while creating webhook for application: ", err.Error())
	}
	app.Webhook = gitWebhook
	app.MetaData.IsWebhookEnabled = gitWebhook.Active
	app.Status = enums.ACTIVE
	applicationMetadataCollection := v1.ApplicationMetadataCollection{
		MetaData: app.MetaData,
		Status:   app.Status,
	}
	err = c.applicationMetadataRepository.Update(companyId, applicationMetadataCollection)
	if err != nil {
		return
	}
	err = c.repo.UpdateApplication(companyId, repoId, app.MetaData.Id, app)
	if err != nil {
		return
	}
}

func (c companyService) GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application {
	return c.repo.GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl)
}
func (c companyService) UpdateRepositories(companyId string, repositories []v1.Repository, companyUpdateOption v1.RepositoryUpdateOption) error {
	option := v1.CompanyQueryOption{LoadRepositories: true, LoadApplications: true, LoadToken: true}
	company, _ := c.repo.GetByCompanyId(companyId, option)
	if company.Id == "" {
		return errors.New("[ERROR] Company does not exist")
	}
	if companyUpdateOption.Option == enums.APPEND_REPOSITORY {
		return c.AppendRepositories(companyId, repositories)
	} else if companyUpdateOption.Option == enums.SOFT_DELETE_REPOSITORY {
		return c.SoftDeleteRepositories(companyId, company, repositories)
	} else if companyUpdateOption.Option == enums.DELETE_REPOSITORY {
		return c.DeleteRepositories(companyId, company, repositories)
	} else {
		return errors.New("invalid repository update option")
	}
}

func (c companyService) AppendRepositories(companyId string, repositories []v1.Repository) error {
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
		return err
	}
	return nil
}

func (c companyService) SoftDeleteRepositories(companyId string, company v1.Company, repositories []v1.Repository) error {
	var count int64
	for i, _ := range company.Repositories {
		for j, _ := range repositories {
			if company.Repositories[i].Id == repositories[j].Id {
				for k := range company.Repositories[i].Applications {
					company.Repositories[i].Applications[k].Status = enums.INACTIVE
					applicationMetadataCollection := v1.ApplicationMetadataCollection{
						MetaData: company.Repositories[i].Applications[k].MetaData,
						Status:   company.Repositories[i].Applications[k].Status,
					}
					err := c.applicationMetadataRepository.Update(companyId, applicationMetadataCollection)
					if err != nil {
						return err
					}
				}
				count++
			}
		}
	}
	if count < 1 {
		return errors.New("repository id does not match")
	}
	err := c.repo.DeleteRepositories(companyId, company.Repositories)
	if err != nil {
		return err
	}
	return nil
}

func (c companyService) DeleteRepositories(companyId string, company v1.Company, repositories []v1.Repository) error {
	var count int64
	for i := range repositories {
		for j, eachRepo := range company.Repositories {
			if repositories[i].Id == eachRepo.Id {
				for _, eachApp := range eachRepo.Applications {
					err := c.applicationMetadataRepository.Delete(eachApp.MetaData.Id, companyId)
					if err != nil {
						return err
					}
				}
				company.Repositories = v1.RemoveRepository(company.Repositories, j)
				count++
				break
			}
		}
	}
	if count < 1 {
		return errors.New("repository id does not match")
	}
	err := c.repo.DeleteRepositories(companyId, company.Repositories)
	if err != nil {
		return err
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

func (c companyService) webHookForGithub(apps []v1.Application, companyId string, token string) {
	for i := range apps {
		usernameOrorgName, repoName := getUsernameAndRepoNameFromGithubRepositoryUrl(apps[i].Url)
		gitWebhook, err := NewGithubService(c, nil, c.client).CreateRepositoryWebhook(usernameOrorgName, repoName, token, companyId)
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
		err = c.applicationMetadataRepository.Store(applicationMetadataCollection)
		if err != nil {
			return
		}
	}
}

func (c companyService) webHookForBitbucket(apps []v1.Application, companyId string, token string) {
	for i := range apps {
		usernameOrOrgName, repoName := getUsernameAndRepoNameFromGithubRepositoryUrl(apps[i].Url)
		b, err := c.client.Get(enums.BITBUCKET_API_BASE_URL+"repositories/"+usernameOrOrgName+"/"+repoName, nil)
		var repositoryDetails v1.BitbucketRepository
		err = yaml.Unmarshal(b, &repositoryDetails)
		if err != nil {
			log.Println(err.Error())
		}
		bitbucketWebhook, err := NewBitBucketService(c, nil, c.client).CreateRepositoryWebhook(repositoryDetails.Workspace.Slug, repositoryDetails.Slug, token, companyId)
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
		err = c.applicationMetadataRepository.Store(applicationMetadataCollection)
		if err != nil {
			return
		}
	}
}

func (c companyService) UpdateApplications(companyId string, repositoryId string, apps []v1.Application, companyUpdateOption v1.ApplicationUpdateOption) error {
	option := v1.CompanyQueryOption{LoadRepositories: true, LoadApplications: true, LoadToken: true}
	company, _ := c.repo.GetByCompanyId(companyId, option)
	if company.Id == "" {
		return errors.New("[ERROR] Company does not exist")
	}
	if companyUpdateOption.Option == enums.APPEND_APPLICATION {
		return c.AppendApplications(companyId, repositoryId, apps, option)
	} else if companyUpdateOption.Option == enums.SOFT_DELETE_APPLICATION {
		return c.SoftDeleteApplications(companyId, repositoryId, company, apps)
	} else if companyUpdateOption.Option == enums.DELETE_APPLICATION {
		return c.DeleteApplications(companyId, repositoryId, company, apps, option)
	} else {
		return errors.New("invalid application update option")
	}
}
func (c companyService) AppendApplications(companyId, repositoryId string, apps []v1.Application, option v1.CompanyQueryOption) error {
	for i := range apps {
		apps[i].MetaData.Id = uuid.New().String()
	}
	repo := c.GetRepositoryByRepositoryId(repositoryId, companyId, option)
	if repo.Id == "" {
		return errors.New("repository not found")
	}
	if repo.Type == enums.GITHUB {
		c.webHookForGithub(apps, companyId, repo.Token)
	} else if repo.Type == enums.BIT_BUCKET {
		c.webHookForBitbucket(apps, companyId, repo.Token)
	}
	err := c.repo.AppendApplications(companyId, repositoryId, apps)
	if err != nil {
		return err
	}
	return nil
}

func (c companyService) SoftDeleteApplications(companyId, repositoryId string, company v1.Company, apps []v1.Application) error {
	for i, each := range company.Repositories {
		if repositoryId == each.Id {
			for j, eachApp := range each.Applications {
				for k := range apps {
					if apps[k].MetaData.Id == eachApp.MetaData.Id {
						company.Repositories[i].Applications[j].Status = enums.INACTIVE
						applicationMetadataCollection := v1.ApplicationMetadataCollection{
							MetaData: company.Repositories[i].Applications[j].MetaData,
							Status:   company.Repositories[i].Applications[j].Status,
						}
						err := c.applicationMetadataRepository.Update(companyId, applicationMetadataCollection)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	err := c.repo.DeleteApplications(companyId, repositoryId, company.Repositories)
	if err != nil {
		return err
	}
	return nil
}

func (c companyService) DeleteApplications(companyId, repositoryId string, company v1.Company, apps []v1.Application, option v1.CompanyQueryOption) error {
	repo := c.GetRepositoryByRepositoryId(repositoryId, companyId, option)
	if repo.Type == enums.GITHUB {
		for i := range apps {
			usernameOrorgName, repoName := getUsernameAndRepoNameFromGithubRepositoryUrl(apps[i].Url)
			err := NewGithubService(c, nil, c.client).DeleteRepositoryWebhookById(usernameOrorgName, repoName, apps[i].Webhook.ID, repo.Token)
			if err != nil {
				return err
			}
			apps[i].MetaData.IsWebhookEnabled = false
			err = c.applicationMetadataRepository.Delete(apps[i].MetaData.Id, companyId)
			if err != nil {
				return err
			}
		}
	} else if repo.Type == enums.BIT_BUCKET {
		for i := range apps {
			usernameOrOrgName, repoName := getUsernameAndRepoNameFromGithubRepositoryUrl(apps[i].Url)
			err := NewBitBucketService(c, nil, c.client).DeleteRepositoryWebhookById(usernameOrOrgName, repoName, apps[i].Webhook.ID, repo.Token)
			if err != nil {
				return err
			}
			apps[i].MetaData.IsWebhookEnabled = false
			err = c.applicationMetadataRepository.Delete(apps[i].MetaData.Id, companyId)
			if err != nil {
				return err
			}
		}
	}
	var count int64 = 0
	var applications []v1.Application
	for i, each := range company.Repositories {
		applications = each.Applications
		if company.Repositories[i].Id == repositoryId {
			for j := range apps {
				for k := range applications {
					if each.Applications[k].MetaData.Id == apps[j].MetaData.Id {
						applications = v1.RemoveApplication(applications, k)
						count++
						break
					}
				}
			}
			company.Repositories[i].Applications = applications
			break
		}
	}
	if count < 1 {
		return errors.New("application id does not match")
	}
	err := c.repo.DeleteApplications(companyId, repositoryId, company.Repositories)
	if err != nil {
		return err
	}
	return nil
}

func (c companyService) GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository {
	return c.repo.GetRepositoryByCompanyIdAndApplicationUrl(id, url)
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
	return c.repo.GetByCompanyId(id, option)
}

func (c companyService) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64) {
	return c.repo.GetRepositoriesByCompanyId(id, option)
}

func (c companyService) GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Application {
	return c.repo.GetApplicationsByCompanyIdAndRepositoryType(id, _type, option, status)
}

// NewCompanyService returns Company type service
func NewCompanyService(repo repository.CompanyRepository, applicationMetadataRepository repository.ApplicationMetadataRepository, client service.HttpClient) service.Company {
	return &companyService{
		repo:                          repo,
		applicationMetadataRepository: applicationMetadataRepository,
		client:                        client,
	}
}
