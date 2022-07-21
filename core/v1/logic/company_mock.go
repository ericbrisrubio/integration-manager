package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
)

type mockCompanyService struct {
	repo   repository.CompanyRepository
	client service.HttpClient
}

func (m mockCompanyService) Delete(companyId string) error {
	//TODO implement me
	panic("implement me")
}

func (m mockCompanyService) GetByName(name string, status v1.StatusQueryOption) v1.Company {
	//TODO implement me
	panic("implement me")
}

func (m mockCompanyService) GetByCompanyId(id string) v1.Company {
	company := v1.Company{MetaData: struct {
		Labels                    map[string]string `bson:"labels" json:"labels" yaml:"labels"`
		NumberOfConcurrentProcess int64             `bson:"number_of_concurrent_process" json:"number_of_concurrent_process" yaml:"number_of_concurrent_process"`
		TotalProcessPerDay        int64             `bson:"total_process_per_day" json:"total_process_per_day" yaml:"total_process_per_day"`
	}{Labels: nil, NumberOfConcurrentProcess: 10, TotalProcessPerDay: 10}}
	return company
}

func (m mockCompanyService) Store(company v1.Company) error {
	return m.repo.Store(company)
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
	}
	m.Store(company)
}

// NewMockCompanyService returns Company type service
func NewMockCompanyService(repo repository.CompanyRepository, client service.HttpClient) service.Company {
	return &mockCompanyService{
		repo:   repo,
		client: client,
	}
}
