package v1

import (
	"errors"
	"fmt"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCompanyMetadata_Validate(t *testing.T) {
	type TestCase struct {
		data     CompanyMetadata
		expected error
		actual   error
	}
	var testcase []TestCase

	labelDemo1 := []map[string]string{{"key1": "value1", "key2": ""}, {"key1": "value1", "key2": "value2"}}
	expectedDemo := []error{errors.New("Company metadata label is missing!"), nil}

	for i := 0; i < 2; i++ {
		testCase := TestCase{
			data:     CompanyMetadata{Labels: labelDemo1[i]},
			expected: expectedDemo[i],
		}
		testcase = append(testcase, testCase)
	}

	for i := 0; i < 2; i++ {
		testcase[i].actual = testcase[i].data.Validate()
		if !reflect.DeepEqual(testcase[i].expected, testcase[i].actual) {
			fmt.Println(testcase[i].actual)
			assert.ElementsMatch(t, testcase[i].expected, testcase[i].actual)
		}
	}
}

func TestApplicationMetadata_Validate(t *testing.T) {
	type TestCase struct {
		data     ApplicationMetadata
		expected error
		actual   error
	}
	var Testdata []TestCase
	labelDemo1 := []map[string]string{{"key1": "value1", "key2": ""}, {"key1": "value1", "key2": "value2"}, {"key1": "value1", "key2": "value2"}}
	IdDemo := []string{"011163003", "", "011163003"}
	nameDemo := []string{"test1", "test2", ""}
	expectedDemo := []error{errors.New("Application metadata label is missing!"), errors.New("Application metadata id is required!"), errors.New("Application metadata name is required!")}

	for i := 0; i < 3; i++ {
		testcase := TestCase{
			data: ApplicationMetadata{
				Labels: labelDemo1[i],
				Id:     IdDemo[i],
				Name:   nameDemo[i],
			},
			expected: expectedDemo[i],
		}
		Testdata = append(Testdata, testcase)
	}
	for i := 0; i < 3; i++ {
		Testdata[i].actual = Testdata[i].data.Validate()
		if !reflect.DeepEqual(Testdata[i].expected, Testdata[i].actual) {
			fmt.Println(Testdata[i].actual)
			assert.ElementsMatch(t, Testdata[i].expected, Testdata[i].actual)
		}
	}
}

func TestApplication_Validate(t *testing.T) {
	type TestCase struct {
		data     Application
		expected error
		actual   error
	}
	var testdata []TestCase
	url := []string{"www.example.com", ""}
	exp := []error{nil, errors.New("Application url is required!")}

	for i := 0; i < 2; i++ {
		testcase := TestCase{
			data: Application{
				MetaData: ApplicationMetadata{
					Labels: nil,
					Id:     "011",
					Name:   "111",
				},
				Url: url[i],
			},
			expected: exp[i],
		}
		testdata = append(testdata, testcase)
	}
	for i := 0; i < 2; i++ {
		testdata[i].actual = testdata[i].data.Validate()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual, i)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}

func TestRepository_Validate(t *testing.T) {
	type TestCase struct {
		data     Repository
		expected error
		actual   error
	}
	var testdata []TestCase

	id := []string{"", "011", "0564", "5478"}
	token := []string{"jwt123", "", "asd555", "asd222"}
	typeDemo := []string{"GITHUB", "BIT_BUCKET", "", "test"}
	expecDemo := []error{errors.New("Repository id is required!"), errors.New("Repository token is required!"), errors.New("Repository type is required"), errors.New("Repository type is invalid!")}

	for i := 0; i < 4; i++ {
		testcase := TestCase{
			data: Repository{
				Id:    id[i],
				Type:  enums.REPOSITORY_TYPE(typeDemo[i]),
				Token: token[i],
				Applications: []Application{{
					MetaData: ApplicationMetadata{
						Labels: nil,
						Id:     "011",
						Name:   "111",
					},
					Url: "www.example.com",
				}},
			},
			expected: expecDemo[i],
		}
		testdata = append(testdata, testcase)
	}
	for i := 0; i < 4; i++ {
		testdata[i].actual = testdata[i].data.Validate()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual, i)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
