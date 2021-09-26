package in_memory

import (
	"fmt"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/enums"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"testing"
)

func TestCompanyRepository_Store(t *testing.T) {
	type Testdata struct {
		data     v1.Company
		expected int64
	}
	metadata := v1.CompanyMetadata{Labels: map[string]string{"env1": "value1"}}
	var applications []v1.Application
	applications = append(applications, v1.Application{
		MetaData: v1.ApplicationMetadata{Labels: map[string]string{"sys1": "value1"}},
		Url:      "www.test.com",
	})
	var repository []v1.Repository
	repository = append(repository, v1.Repository{
		Type:         enums.Inmemory,
		Token:        "123456",
		Applications: applications,
	})
	testdata := []Testdata{
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "01",
			Name:         "test1",
			Repositories: repository,
			Status:       "active",
		}, expected: 1},
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "02",
			Name:         "test2",
			Repositories: repository,
			Status:       "active",
		}, expected: 2},
	}

	repo := NewCompanyRepository(10)
	for _, each := range testdata {
		repo.Store(each.data)
		com, size := repo.GetCompanies(v1.CompanyQueryOption{
			Pagination: v1.Pagination{
				Page:  0,
				Limit: 10,
			},
			LoadRepositories: true,
			LoadApplications: true,
		})
		log.Println(size, " ", each.expected, size != each.expected, com)
		if size != each.expected {
			assert.ElementsMatch(t, size, each.expected)
		}
	}
}

func TestCompanyRepository_GetCompanies(t *testing.T) {
	type Testdata struct {
		data     v1.Company
		expected int64
	}
	metadata := v1.CompanyMetadata{Labels: map[string]string{"env1": "value1"}}
	var applications []v1.Application
	applications = append(applications, v1.Application{
		MetaData: v1.ApplicationMetadata{Labels: map[string]string{"sys1": "value1"}},
		Url:      "www.test.com",
	})
	var repository []v1.Repository
	repository = append(repository, v1.Repository{
		Type:         enums.Inmemory,
		Token:        "123456",
		Applications: applications,
	})
	testdata := []Testdata{
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "01",
			Name:         "test1",
			Repositories: repository,
			Status:       "active",
		}, expected: 2},
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "02",
			Name:         "test2",
			Repositories: repository,
			Status:       "active",
		}, expected: 2},
	}

	repo := NewCompanyRepository(10)
	for i, each := range testdata {
		flag := true
		if i%2 == 0 {
			flag = false
		}
		com, size := repo.GetCompanies(v1.CompanyQueryOption{
			Pagination: v1.Pagination{
				Page:  0,
				Limit: 10,
			},
			LoadRepositories: flag,
			LoadApplications: flag,
		})
		log.Println(size, " ", each.expected, size != each.expected, com)
		if size != each.expected {
			assert.ElementsMatch(t, size, each.expected)
		}
	}
}

func TestCompanyRepository_GetByCompanyId(t *testing.T) {
	type Testdata struct {
		data     v1.Company
		expected v1.Company
		actual   v1.Company
	}
	metadata := v1.CompanyMetadata{Labels: map[string]string{"env1": "value1"}}
	var applications []v1.Application
	applications = append(applications, v1.Application{
		MetaData: v1.ApplicationMetadata{Labels: map[string]string{"sys1": "value1"}},
		Url:      "www.test.com",
	})
	var repository []v1.Repository
	repository = append(repository, v1.Repository{
		Type:         enums.Inmemory,
		Token:        "123456",
		Applications: applications,
	})
	testdata := []Testdata{
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "01",
			Name:         "test1",
			Repositories: repository,
			Status:       "active",
		}, expected: v1.Company{MetaData: metadata, Id: "01", Name: "test1", Repositories: repository, Status: "active"}},
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "02",
			Name:         "test2",
			Repositories: repository,
			Status:       "active",
		}, expected: v1.Company{MetaData: metadata, Id: "01", Name: "test1", Repositories: repository, Status: "active"}},
	}

	repo := NewCompanyRepository(10)
	for _, each := range testdata {
		each.actual, _ = repo.GetByCompanyId("01", v1.CompanyQueryOption{
			Pagination: v1.Pagination{
				Page:  0,
				Limit: 10,
			},
			LoadRepositories: true,
			LoadApplications: true,
		})
		if !reflect.DeepEqual(each.expected, each.actual) {
			assert.ElementsMatch(t, each.expected, each.actual)
		}
	}
}

func TestCompanyRepository_GetRepositoriesByCompanyId(t *testing.T) {
	type Testdata struct {
		data     v1.Company
		expected []v1.Repository
		actual   []v1.Repository
	}
	metadata := v1.CompanyMetadata{Labels: map[string]string{"env1": "value1"}}
	var applications []v1.Application
	applications = append(applications, v1.Application{
		MetaData: v1.ApplicationMetadata{Labels: map[string]string{"sys1": "value1"}},
		Url:      "www.test.com",
	})
	var repository []v1.Repository
	repository = append(repository, v1.Repository{
		Type:         enums.Inmemory,
		Token:        "123456",
		Applications: applications,
	})
	testdata := []Testdata{
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "01",
			Name:         "test1",
			Repositories: repository,
			Status:       "active",
		}, expected: []v1.Repository{
			{
				Type:         enums.Inmemory,
				Token:        "123456",
				Applications: applications,
			},
		}},
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "02",
			Name:         "test2",
			Repositories: nil,
			Status:       "active",
		}, expected: nil},
	}

	repo := NewCompanyRepository(10)
	for _, each := range testdata {
		repo.Store(each.data)
		each.actual, _ = repo.GetRepositoriesByCompanyId("02", v1.CompanyQueryOption{
			Pagination: v1.Pagination{
				Page:  0,
				Limit: 10,
			},
			LoadRepositories: true,
			LoadApplications: true,
		})
		fmt.Println(each.actual)
		if !reflect.DeepEqual(each.expected, each.actual) {
			assert.ElementsMatch(t, each.expected, each.actual)
		}
	}
}

func TestCompanyRepository_GetApplicationsByCompanyId(t *testing.T) {
	type Testdata struct {
		data     v1.Company
		expected []v1.Application
		actual   []v1.Application
	}
	metadata := v1.CompanyMetadata{Labels: map[string]string{"env1": "value1"}}
	var applications []v1.Application
	applications = append(applications, v1.Application{
		MetaData: v1.ApplicationMetadata{Labels: map[string]string{"sys1": "value1"}},
		Url:      "www.test.com",
	})
	var repository []v1.Repository
	repository = append(repository, v1.Repository{
		Type:         enums.Inmemory,
		Token:        "123456",
		Applications: applications,
	})
	testdata := []Testdata{
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "01",
			Name:         "test1",
			Repositories: repository,
			Status:       "active",
		}, expected: []v1.Application{
			{
				MetaData: v1.ApplicationMetadata{Labels: map[string]string{"sys1": "value1"}},
				Url:      "www.test.com",
			},
		}},
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "02",
			Name:         "test2",
			Repositories: nil,
			Status:       "active",
		}, expected: nil},
	}

	repo := NewCompanyRepository(10)
	for _, each := range testdata {
		repo.Store(each.data)
		each.actual, _ = repo.GetApplicationsByCompanyId(each.data.Id, v1.CompanyQueryOption{
			Pagination: v1.Pagination{
				Page:  0,
				Limit: 10,
			},
			LoadRepositories: true,
			LoadApplications: true,
		})
		fmt.Println(each.actual)
		if !reflect.DeepEqual(each.expected, each.actual) {
			assert.ElementsMatch(t, each.expected, each.actual)
		}
	}
}
func TestCompanyRepository_GetApplicationsByCompanyIdAndRepositoryType(t *testing.T) {
	type Testdata struct {
		data     v1.Company
		expected []v1.Application
		actual   []v1.Application
	}
	metadata := v1.CompanyMetadata{Labels: map[string]string{"env1": "value1"}}
	var applications1 []v1.Application
	applications1 = append(applications1, v1.Application{
		MetaData: v1.ApplicationMetadata{Labels: map[string]string{"sys1": "value1"}},
		Url:      "www.test.com",
	})
	var applications2 []v1.Application
	applications2 = append(applications2, v1.Application{
		MetaData: v1.ApplicationMetadata{Labels: map[string]string{"sys2": "value2"}},
		Url:      "www.test1.com",
	})
	var repository []v1.Repository
	repository = append(repository, v1.Repository{
		Type:         enums.GITHUB,
		Token:        "123456",
		Applications: applications1,
	})
	var repository1 []v1.Repository
	repository1 = append(repository1, v1.Repository{
		Type:         enums.BIT_BUCKET,
		Token:        "123456",
		Applications: applications2,
	})
	testdata := []Testdata{
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "01",
			Name:         "test1",
			Repositories: repository,
			Status:       "active",
		}, expected: []v1.Application{
			{
				MetaData: v1.ApplicationMetadata{Labels: map[string]string{"sys1": "value1"}},
				Url:      "www.test.com",
			},
		}},
		{data: v1.Company{
			MetaData:     metadata,
			Id:           "02",
			Name:         "test2",
			Repositories: repository1,
			Status:       "active",
		}, expected: []v1.Application{
			{
				MetaData: v1.ApplicationMetadata{Labels: map[string]string{"sys2": "value2"}},
				Url:      "www.test1.com",
			},
		}},
	}

	repo := NewCompanyRepository(10)
	for _, each := range testdata {
		repo.Store(each.data)
		for _, eachRepository := range each.data.Repositories {
			each.actual = repo.GetApplicationsByCompanyIdAndRepositoryType(each.data.Id, eachRepository.Type, v1.CompanyQueryOption{
				Pagination: v1.Pagination{
					Page:  0,
					Limit: 10,
				},
				LoadRepositories: true,
				LoadApplications: true,
			})
			fmt.Println(each.actual)
			if !reflect.DeepEqual(each.expected, each.actual) {
				assert.ElementsMatch(t, each.expected, each.actual)
			}
		}
	}
}
