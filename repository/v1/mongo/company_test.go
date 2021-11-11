package mongo

import (
	"github.com/joho/godotenv"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path"
	"reflect"
	"testing"
)

func loadEnv(t *testing.T) error {
	dirname, err := os.Getwd()
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	dir, err := os.Open(path.Join(dirname, "../../../"))
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	err = godotenv.Load(os.ExpandEnv(dir.Name() + "/.env.mongo.test"))
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	return err
}

func TestCompanyRepository_Store(t *testing.T) {
	type TestData struct {
		expected int64
		actual   int64
	}

	testCase := TestData{
		expected: 7,
	}

	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}
	_, testCase.actual = repo.GetCompanies(option)
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}

func TestCompanyRepository_AppendRepositories(t *testing.T) {
	type TestData struct {
		companyId    string
		repositories []v1.Repository
		expected     int64
		actual       int64
	}
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	testCase := TestData{
		companyId: "07",
		repositories: []v1.Repository{
			{
				Id:    "08",
				Type:  enums.GITHUB,
				Token: "0111278",
				Applications: []v1.Application{
					{
						MetaData: v1.ApplicationMetadata{
							Labels:           map[string]string{"compId": "1238", "teamId": "1238"},
							Id:               "078",
							Name:             "test78",
							IsWebhookEnabled: false,
						},
						Url: "https://github.com/klovercloud-ci-cd/integration-manager",
					},
				},
			},
			{
				Id:    "09",
				Type:  enums.GITHUB,
				Token: "0111279",
				Applications: []v1.Application{
					{
						MetaData: v1.ApplicationMetadata{
							Labels:           map[string]string{"compId": "1239", "teamId": "1239"},
							Id:               "079",
							Name:             "test79",
							IsWebhookEnabled: false,
						},
						Url: "https://github.com/klovercloud-ci-cd/integration-manager",
					},
				},
			},
		},
		expected: 3,
	}

	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}
	err = repo.AppendRepositories(testCase.companyId, testCase.repositories)
	if err != nil {
		return
	}
	_, testCase.actual = repo.GetRepositoriesByCompanyId("07", option)
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}

func TestCompanyRepository_DeleteRepositories(t *testing.T) {
	type TestData struct {
		companyId    string
		repositories []v1.Repository
		expected     int64
		actual       int64
	}
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	testCase := TestData{
		companyId: "07",
		repositories: []v1.Repository{
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
		expected: 0,
	}

	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}
	err = repo.DeleteRepositories(testCase.companyId, testCase.repositories, false)
	if err != nil {
		return
	}
	_, testCase.actual = repo.GetRepositoriesByCompanyId("07", option)
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}

func TestCompanyRepository_AppendApplications(t *testing.T) {
	type TestData struct {
		companyId    string
		repositoryId string
		applications []v1.Application
		expected     int64
		actual       int64
	}
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	testCase := TestData{
		companyId:    "07",
		repositoryId: "07",
		applications: []v1.Application{
			{
				MetaData: v1.ApplicationMetadata{
					Labels:           map[string]string{"compId": "123122", "teamId": "123122"},
					Id:               "071",
					Name:             "test71",
					IsWebhookEnabled: false,
				},
				Url: "https://github.com/klovercloud-ci-cd/integration-manager",
			},
			{
				MetaData: v1.ApplicationMetadata{
					Labels:           map[string]string{"compId": "123122", "teamId": "123122"},
					Id:               "072",
					Name:             "test72",
					IsWebhookEnabled: false,
				},
				Url: "https://github.com/klovercloud-ci-cd/integration-manager",
			},
		},
		expected: 3,
	}

	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}
	err = repo.AppendApplications(testCase.companyId, testCase.repositoryId, testCase.applications)
	if err != nil {
		return
	}
	_, testCase.actual = repo.GetApplicationsByCompanyId("07", option)
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}

func TestCompanyRepository_DeleteApplications(t *testing.T) {
	type TestData struct {
		companyId    string
		repositoryId string
		applications []v1.Application
		expected     int64
		actual       int64
	}
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	testCase := TestData{
		companyId:    "07",
		repositoryId: "07",
		applications: []v1.Application{
			{
				MetaData: v1.ApplicationMetadata{
					Labels:           map[string]string{"compId": "123122", "teamId": "123122"},
					Id:               "071",
					Name:             "test71",
					IsWebhookEnabled: false,
				},
				Url: "https://github.com/klovercloud-ci-cd/integration-manager",
			},
			{
				MetaData: v1.ApplicationMetadata{
					Labels:           map[string]string{"compId": "123122", "teamId": "123122"},
					Id:               "072",
					Name:             "test72",
					IsWebhookEnabled: false,
				},
				Url: "https://github.com/klovercloud-ci-cd/integration-manager",
			},
		},
		expected: 1,
	}

	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}
	err = repo.DeleteApplications(testCase.companyId, testCase.repositoryId, testCase.applications, false)
	if err != nil {
		return
	}
	_, testCase.actual = repo.GetApplicationsByCompanyId("07", option)
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}

func TestCompanyRepository_GetCompanyByApplicationUrl(t *testing.T) {
	type TestData struct {
		url      string
		expected v1.Company
		actual   v1.Company
	}
	testCase := TestData{
		url: "https://github.com/klovercloud-ci-cd/integration-manager",
		expected: v1.Company{
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
	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}
	testCase.actual = repo.GetCompanyByApplicationUrl(testCase.url)
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}

func TestCompanyRepository_GetRepositoryByRepositoryId(t *testing.T) {
	type TestData struct {
		repositoryId string
		expected     v1.Repository
		actual       v1.Repository
	}
	testCase := TestData{
		repositoryId: "07",
		expected: v1.Repository{
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
	}
	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}
	testCase.actual = repo.GetRepositoryByRepositoryId(testCase.repositoryId)
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}
func TestCompanyRepository_GetApplicationsByCompanyIdAndRepositoryType(t *testing.T) {
	type TestData struct {
		companyId string
		repoType  string
		expected  []v1.Application
		actual    []v1.Application
	}
	testCase := TestData{
		companyId: "07",
		repoType:  "GITHUB",
		expected: []v1.Application{
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
	}
	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	testCase.actual = repo.GetApplicationsByCompanyIdAndRepositoryType(testCase.companyId, enums.REPOSITORY_TYPE(testCase.repoType), option)
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}

func TestCompanyRepository_GetRepositoryByCompanyIdAndApplicationUrl(t *testing.T) {
	type TestData struct {
		companyId string
		url       string
		expected  v1.Repository
		actual    v1.Repository
	}
	testCase := TestData{
		companyId: "07",
		url:       "https://github.com/klovercloud-ci-cd/integration-manager",
		expected: v1.Repository{
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
	}
	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}

	testCase.actual = repo.GetRepositoryByCompanyIdAndApplicationUrl(testCase.companyId, testCase.url)
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}

func TestCompanyRepository_GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(t *testing.T) {
	type TestData struct {
		companyId string
		repoId    string
		url       string
		expected  v1.Application
		actual    v1.Application
	}
	testCase := TestData{
		companyId: "07",
		repoId:    "07",
		url:       "https://github.com/klovercloud-ci-cd/integration-manager",
		expected: v1.Application{
			MetaData: v1.ApplicationMetadata{
				Labels:           map[string]string{"compId": "123", "teamId": "123"},
				Id:               "07",
				Name:             "test7",
				IsWebhookEnabled: false,
			},
			Url: "https://github.com/klovercloud-ci-cd/integration-manager",
		},
	}
	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}

	testCase.actual = repo.GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(testCase.companyId, testCase.repoId, testCase.url)
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}

func TestCompanyRepository_DeleteApplicationsSoftDelete(t *testing.T) {
	type TestData struct {
		companyId    string
		repositoryId string
		applications []v1.Application
		expected     enums.COMPANY_STATUS
		actual       enums.COMPANY_STATUS
	}
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	testCase := TestData{
		companyId:    "07",
		repositoryId: "07",
		applications: []v1.Application{
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
		expected: enums.INACTIVE,
	}

	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}
	err = repo.DeleteApplications(testCase.companyId, testCase.repositoryId, testCase.applications, true)
	if err != nil {
		return
	}
	apps, _ := repo.GetApplicationsByCompanyId("07", option)
	for _, each := range apps {
		if each.MetaData.Id == "07" {
			testCase.actual = each.Status
			break
		}
	}
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}

func TestCompanyRepository_DeleteRepositoriesSoftDelete(t *testing.T) {
	type TestData struct {
		companyId    string
		repositories []v1.Repository
		expected     enums.COMPANY_STATUS
		actual       enums.COMPANY_STATUS
	}
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	testCase := TestData{
		companyId: "07",
		repositories: []v1.Repository{
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
		expected: "INACTIVE",
	}
	err := loadEnv(t)
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	repo := NewMockCompanyRepository()
	data := InitCompanyData()
	for _, each := range data {
		err := repo.Store(each)
		if err != nil {
			log.Println("ERROR:", err.Error())
			t.Fail()
		}
	}
	err = repo.DeleteRepositories(testCase.companyId, testCase.repositories, true)
	if err != nil {
		return
	}
	repos, _ := repo.GetRepositoriesByCompanyId("07", option)
	for _, each := range repos {
		if each.Id == "07" {
			for _, each := range each.Applications {
				testCase.actual = each.Status
				break
			}
		}
	}
	if !reflect.DeepEqual(testCase.expected, testCase.actual) {
		log.Println("ERROR:", "expected:", testCase.expected, "actual:", testCase.actual)
		assert.ElementsMatch(t, testCase.expected, testCase.actual)
	}
}
