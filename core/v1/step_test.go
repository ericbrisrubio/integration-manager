package v1

import (
	"errors"
	"fmt"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStep_Validate(t *testing.T) {
	type TestCase struct {
		data     Step
		expected error
		actual   error
	}

	var testdata []TestCase
	name := []string{"", "test1", "test2", "test3", "test4", "Test5"}
	typeD := []enums.STEP_TYPE{"DEPLOY", "DEPLOY", "BUILD", "", "sss", "BUILD"}
	trigger := []enums.TRIGGER{"AUTO", "", "ssss", "MANUAL", "AUTO", "AUTO"}
	params := []map[enums.PARAMS]string{{"type": "sss", "env": "12223"}, {"type": "sss", "env": "12223"}, {"type": "sss", "env": "12223"}, {"type": "sss", "env": "12223"}, {"type": "sss", "env": "12223"}, {"type": "", "env": "12223"}}
	expec := []error{errors.New("step name is required!"), errors.New("step trigger is required"), errors.New("step trigger is invalid"), errors.New("step type is required"), errors.New("step type is invalid"), errors.New("step params is missing!")}

	for i := 0; i < 6; i++ {
		testcase := TestCase{
			data: Step{
				Name:        name[i],
				Type:        typeD[i],
				Trigger:     trigger[i],
				Params:      params[i],
				Next:        nil,
				Descriptors: nil,
			},
			expected: expec[i],
		}
		testdata = append(testdata, testcase)
	}
	for i := 0; i < 6; i++ {
		testdata[i].actual = testdata[i].data.Validate()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
