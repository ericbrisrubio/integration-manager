package v1

import (
	"errors"
	"fmt"
	"github.com/klovercloud-ci/enums"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCompany_Validate(t *testing.T) {
	type TestCase struct {
		data     Company
		expected error
		actual   error
	}
	name := []string{"sss", "", "kkk", "joom"}
	id := []string{"", "555", "111", "7774"}
	status := []enums.COMPANY_STATUS{"ACTIVE", "ACTIVE", "sss", ""}
	exp := []error{errors.New("Company id is required!"), errors.New("Company name is required!"), errors.New("Company status invalid!"), errors.New("Company status is required!")}

	var testdata []TestCase

	for i := 0; i < 4; i++ {
		testcase := TestCase{
			data: Company{
				MetaData:     CompanyMetadata{Labels: map[string]string{"aaa": "aaa", "sss": "ss"}},
				Id:           id[i],
				Name:         name[i],
				Repositories: nil,
				Status:       status[i],
			},
			expected: exp[i],
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
