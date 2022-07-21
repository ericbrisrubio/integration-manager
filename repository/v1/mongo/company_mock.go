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
			Id:     "01",
			Name:   "test1",
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 3, TotalProcessPerDay: 4}),
			Id:     "02",
			Name:   "test2",
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 2, TotalProcessPerDay: 3}),
			Id:     "03",
			Name:   "test3",
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 2, TotalProcessPerDay: 3}),
			Id:     "04",
			Name:   "test4",
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 2, TotalProcessPerDay: 3}),
			Id:     "05",
			Name:   "test5",
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 2, TotalProcessPerDay: 3}),
			Id:     "06",
			Name:   "test6",
			Status: enums.ACTIVE,
		},
		{
			MetaData: v1.CompanyMetadata(struct {
				Labels                    map[string]string
				NumberOfConcurrentProcess int64
				TotalProcessPerDay        int64
			}{Labels: map[string]string{"compId": "123", "teamId": "123"}, NumberOfConcurrentProcess: 0, TotalProcessPerDay: 0}),
			Id:     "07",
			Name:   "test7",
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
