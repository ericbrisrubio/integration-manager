package logic

import (
	"fmt"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
)

type companyService struct {
	repo repository.CompanyRepository
}

func (c companyService) GetRepositoryByRepositoryId(id string) v1.Repository {
	return c.repo.GetRepositoryByRepositoryId(id)
}

func (c companyService) GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application {
	//return c.repo.GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl)
	return v1.Application{
		MetaData: v1.ApplicationMetadata{
			Id: "1001",
		},
	}
}
func (c companyService) UpdateRepositories(company v1.Company, companyUpdateOption v1.CompanyUpdateOption) error {
	if companyUpdateOption.Option == enums.APPEND_REPOSITORY {
		err := c.repo.AppendRepositories(company.Id, company.Repositories)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	if companyUpdateOption.Option == enums.SOFT_DELETE_REPOSITORY {
		err := c.repo.DeleteRepositories(company.Id, company.Repositories, true)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	if companyUpdateOption.Option == enums.DELETE_REPOSITORY {
		err := c.repo.DeleteRepositories(company.Id, company.Repositories, false)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (c companyService) UpdateApplications(companyId string, repositoryId string, apps []v1.Application, companyUpdateOption v1.CompanyUpdateOption) {
	if companyUpdateOption.Option == enums.APPEND_APPLICATION {
		err := c.repo.AppendApplications(companyId, repositoryId, apps)
		if err != nil {
			fmt.Println(err)
		}
	}
	if companyUpdateOption.Option == enums.SOFT_DELETE_APPLICATION {
		err := c.repo.DeleteApplications(companyId, repositoryId, apps, true)
		if err != nil {
			fmt.Println(err)
		}
	}
	if companyUpdateOption.Option == enums.DELETE_APPLICATION {
		err := c.repo.DeleteApplications(companyId, repositoryId, apps, false)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c companyService) GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository {
	//return c.repo.GetRepositoryByCompanyIdAndApplicationUrl(id, url)
	return v1.Repository{
		Id:    "1",
		Token: "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",
	}
}

func (c companyService) GetCompanyByApplicationUrl(url string) v1.Company {
	return c.repo.GetCompanyByApplicationUrl(url)
}

func (c companyService) Store(company v1.Company) error {
	return c.repo.Store(company)
}

func (c companyService) Delete(companyId string) error {
	err := c.repo.Delete(companyId)
	if err != nil {
		return err
	}
	return nil
}

func (c companyService) GetCompanies(option v1.CompanyQueryOption) []v1.Company {
	companies, _ := c.repo.GetCompanies(option)

	return companies
}

func (c companyService) GetByCompanyId(id string, option v1.CompanyQueryOption) (v1.Company, int64) {
	company, total := c.repo.GetByCompanyId(id, option)

	return company, total
}

func (c companyService) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64) {
	repositories, total := c.repo.GetRepositoriesByCompanyId(id, option)

	return repositories, total
}

func (c companyService) GetApplicationsByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Application, int64) {
	applications, total := c.repo.GetApplicationsByCompanyId(id, option)

	return applications, total
}

func (c companyService) GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption) []v1.Application {
	applications := c.repo.GetApplicationsByCompanyIdAndRepositoryType(id, _type, option)

	return applications
}

func NewCompanyService(repo repository.CompanyRepository) service.Company {
	return &companyService{
		repo: repo,
	}
}
