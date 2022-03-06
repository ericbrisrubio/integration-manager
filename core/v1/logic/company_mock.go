package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

type mockCompanyService struct {
	repo   repository.CompanyRepository
	client service.HttpClient
}

func (m mockCompanyService) GetAllApplications(companyId string, option v1.CompanyQueryOption) ([]v1.Application, int64) {
	panic("implement me")
}

func (m mockCompanyService) GetRepositoryByRepositoryId(id string, companyId string, option v1.CompanyQueryOption) v1.Repository {
	//TODO implement me
	panic("implement me")
}

func (m mockCompanyService) GetApplicationsByRepositoryId(repositoryId string, companyId string, option v1.CompanyQueryOption, status v1.StatusQueryOption) ([]v1.Application, int64) {
	//TODO implement me
	panic("implement me")
}

func (m mockCompanyService) GetApplicationByApplicationId(companyId string, repoId string, applicationId string) v1.Application {
	//TODO implement me
	panic("implement me")
}

func (m mockCompanyService) Store(company v1.Company) error {
	return m.repo.Store(company)
}

func (m mockCompanyService) UpdateRepositories(companyId string, repositories []v1.Repository, companyUpdateOption v1.RepositoryUpdateOption) error {
	//TODO implement me
	panic("implement me")
}

func (m mockCompanyService) UpdateApplications(companyId string, repositoryId string, apps []v1.Application, companyUpdateOption v1.ApplicationUpdateOption) error {
	//TODO implement me
	panic("implement me")
}

func (m mockCompanyService) Delete(companyId string) error {
	//TODO implement me
	panic("implement me")
}

func (m mockCompanyService) GetCompanies(option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Company {
	companies, _ := m.repo.GetCompanies(option, status)
	return companies
}
func (m mockCompanyService) initMockCompanyData() {
	company := v1.Company{
		MetaData: struct {
			Labels                    map[string]string `bson:"labels" json:"labels" yaml:"labels"`
			NumberOfConcurrentProcess int64             `bson:"number_of_concurrent_process" json:"number_of_concurrent_process" yaml:"number_of_concurrent_process"`
			TotalProcessPerDay        int64             `bson:"total_process_per_day" json:"total_process_per_day" yaml:"total_process_per_day"`
		}{Labels: nil, NumberOfConcurrentProcess: 10, TotalProcessPerDay: 10},
		Id: "1",
		Repositories: []v1.Repository{
			{Id: "1", Token: "", Applications: []v1.Application{{MetaData: struct {
				Labels           map[string]string `bson:"labels" json:"labels"`
				Id               string            `bson:"id" json:"id"`
				Name             string            `bson:"name" json:"name"`
				IsWebhookEnabled bool              `bson:"is_webhook_enabled" json:"is_webhook_enabled"`
			}{Id: "1001", IsWebhookEnabled: true}}}},
		},
	}
	m.Store(company)
}
func (m mockCompanyService) GetByCompanyId(id string, option v1.CompanyQueryOption) (v1.Company, int64) {
	company := v1.Company{MetaData: struct {
		Labels                    map[string]string `bson:"labels" json:"labels" yaml:"labels"`
		NumberOfConcurrentProcess int64             `bson:"number_of_concurrent_process" json:"number_of_concurrent_process" yaml:"number_of_concurrent_process"`
		TotalProcessPerDay        int64             `bson:"total_process_per_day" json:"total_process_per_day" yaml:"total_process_per_day"`
	}{Labels: nil, NumberOfConcurrentProcess: 10, TotalProcessPerDay: 10}}
	var total int64
	return company, total
}

func (m mockCompanyService) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64) {
	return m.repo.GetRepositoriesByCompanyId(id, option)
}

func (m mockCompanyService) GetCompanyByApplicationUrl(url string) v1.Company {
	return m.repo.GetCompanyByApplicationUrl(url)
}

func (m mockCompanyService) GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Application {
	return m.repo.GetApplicationsByCompanyIdAndRepositoryType(id, _type, option, status)
}

func (m mockCompanyService) GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository {
	return v1.Repository{
		Id:    "flameOfDimitry",
		Token: "",
	}
}

func (m mockCompanyService) GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application {
	return v1.Application{
		MetaData: v1.ApplicationMetadata{
			Id:               "1001",
			IsWebhookEnabled: true,
		},
	}
}

// NewMockCompanyService returns Company type service
func NewMockCompanyService(repo repository.CompanyRepository, client service.HttpClient) service.Company {
	return &mockCompanyService{
		repo:   repo,
		client: client,
	}
}
