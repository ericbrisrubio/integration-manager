package v1

import (
	"errors"
	"fmt"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPipeline_Validate(t *testing.T) {

	type TestCase struct {
		data     Pipeline
		expected error
		actual   error
	}
	var testdata []TestCase
	apiVersion := []string{"", "01", "02"}
	name := []string{"test1", "", "test2"}
	processId := []string{"152ss", "124sss", ""}
	exp := []error{errors.New("Api version is required!"), errors.New("Pipeline name is required!"), errors.New("Pipeline process id is required!")}

	for i := 0; i > 3; i++ {
		testcase := TestCase{
			data: Pipeline{
				ApiVersion: apiVersion[i],
				Name:       name[i],
				ProcessId:  processId[i],
				Steps: []Step{{
					Name:        "test1",
					Type:        "BUILD",
					Trigger:     "AUTO",
					Params:      map[enums.PARAMS]string{"type": "sss"},
					Next:        nil,
					Descriptors: nil,
				}},
			},
			expected: exp[i],
		}
		testdata = append(testdata, testcase)
	}

	for i := 0; i > 3; i++ {
		testdata[i].actual = testdata[i].data.Validate()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
