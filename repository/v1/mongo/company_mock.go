package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

var data []v1.Company

// InitCompanyData loads mock company data
func InitCompanyData() []v1.Company {
	var data []v1.Company

	data = []v1.Company{
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 2, TotalProcessPerDay: 3}),
			Id:   "01",
			Name: "test1",
			Repositories: []v1.Repository{
				{
					Id:    "01",
					Type:  enums.GITHUB,
					Token: "01112",
					Applications: []v1.Application{
						{
							MetaData: v1.ApplicationMetadata{
								Labels:           map[string]string{"compId": "123", "teamId": "123"},
								Id:               "01",
								Name:             "test1",
								IsWebhookEnabled: false,
							},
							Url: "https://github.com/klovercloud-ci-cd/integration-manager",
						},
					},
				},
			},
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 3, TotalProcessPerDay: 4}),
			Id:   "02",
			Name: "test2",
			Repositories: []v1.Repository{
				{
					Id:    "02",
					Type:  enums.GITHUB,
					Token: "0111",
					Applications: []v1.Application{
						{
							MetaData: v1.ApplicationMetadata{
								Labels:           map[string]string{"compId": "123", "teamId": "123"},
								Id:               "02",
								Name:             "test2",
								IsWebhookEnabled: false,
							},
							Url: "https://github.com/klovercloud-ci-cd/integration-manager",
						},
					},
				},
			},
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 2, TotalProcessPerDay: 3}),
			Id:   "03",
			Name: "test3",
			Repositories: []v1.Repository{
				{
					Id:    "03",
					Type:  enums.GITHUB,
					Token: "011123",
					Applications: []v1.Application{
						{
							MetaData: v1.ApplicationMetadata{
								Labels:           map[string]string{"compId": "123", "teamId": "123"},
								Id:               "03",
								Name:             "test3",
								IsWebhookEnabled: false,
							},
							Url: "https://github.com/klovercloud-ci-cd/integration-manager",
						},
					},
				},
			},
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 2, TotalProcessPerDay: 3}),
			Id:   "04",
			Name: "test4",
			Repositories: []v1.Repository{
				{
					Id:    "04",
					Type:  enums.GITHUB,
					Token: "011124",
					Applications: []v1.Application{
						{
							MetaData: v1.ApplicationMetadata{
								Labels:           map[string]string{"compId": "123", "teamId": "123"},
								Id:               "04",
								Name:             "test4",
								IsWebhookEnabled: false,
							},
							Url: "https://github.com/klovercloud-ci-cd/integration-manager",
						},
					},
				},
			},
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 2, TotalProcessPerDay: 3}),
			Id:   "05",
			Name: "test5",
			Repositories: []v1.Repository{
				{
					Id:    "05",
					Type:  enums.GITHUB,
					Token: "011125",
					Applications: []v1.Application{
						{
							MetaData: v1.ApplicationMetadata{
								Labels:           map[string]string{"compId": "123", "teamId": "123"},
								Id:               "05",
								Name:             "test5",
								IsWebhookEnabled: false,
							},
							Url: "https://github.com/klovercloud-ci-cd/integration-manager",
						},
					},
				},
			},
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 2, TotalProcessPerDay: 3}),
			Id:   "06",
			Name: "test6",
			Repositories: []v1.Repository{
				{
					Id:    "06",
					Type:  enums.GITHUB,
					Token: "011126",
					Applications: []v1.Application{
						{
							MetaData: v1.ApplicationMetadata{
								Labels:           map[string]string{"compId": "123", "teamId": "123"},
								Id:               "06",
								Name:             "test6",
								IsWebhookEnabled: false,
							},
							Url: "https://github.com/klovercloud-ci-cd/integration-manager",
						},
					},
				},
			},
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 0, TotalProcessPerDay: 0}),
			Id:   "07",
			Name: "test7",
			Repositories: []v1.Repository{
				{
					Id:    "07",
					Type:  enums.GITHUB,
					Token: "011127",
					Applications: []v1.Application{
						{
							MetaData: v1.ApplicationMetadata{
								Labels:           map[string]string{"compId": "123", "teamId": "123"},
								Id:               "07",
								Name:             "test7",
								IsWebhookEnabled: false,
							},
							Url: "https://github.com/klovercloud-ci-cd/integration-manager",
						},
					},
				},
			},
			Status: enums.ACTIVE,
		},
	}
	return data
}

// NewMockCompanyRepository returns CompanyRepository type object
func NewMockCompanyRepository() repository.CompanyRepository {
	manager := GetMockDmManager()
	err := manager.Db.Drop(context.Background())
	if err != nil {
		return nil
	}
	return &companyRepository{
		manager: GetMockDmManager(),
		timeout: 3000,
	}

}
