package v1

import (
	"errors"
	"fmt"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Test function for testing step data - name, type, trigger, params
func TestStep_Validate(t *testing.T) {
	type TestCase struct {
		data     Step
		expected error
		actual   error
	}

	var testdata []TestCase
	name := []string{"", "test1", "test2", "test3", "test4", "test5", "aaaaaaaaaaaaaaaaaaaaaa", "123ABCabc", "$%&#@"}
	typeD := []enums.STEP_TYPE{"DEPLOY", "DEPLOY", "BUILD", "", "sss", "BUILD", "BUILD", "BUILD", "BUILD"}
	trigger := []enums.TRIGGER{"AUTO", "", "ssss", "MANUAL", "AUTO", "AUTO", "AUTO", "AUTO", "AUTO"}
	params := []map[enums.PARAMS]string{{"type": "sss", "env": "12223"}, {"type": "sss", "env": "12223"}, {"type": "sss", "env": "12223"}, {"type": "sss", "env": "12223"}, {"type": "sss", "env": "12223"}, {"type": "", "env": "12223"}, {"type": "abc", "env": "12223"}, {"type": "abc", "env": "12223"}, {"type": "abc", "env": "12223"}}
	expec := []error{errors.New("step name is required"), errors.New("step trigger is required"), errors.New("step trigger is invalid"), errors.New("step type is required"), errors.New("step type is invalid"), errors.New("step params is missing"), errors.New("step name length cannot be more than 16 character"), errors.New("step name can only contain lower case characters or digits"), errors.New("step name can only contain lower case characters or digits")}

	for i := 0; i < 9; i++ {
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
	for i := 0; i < 9; i++ {
		testdata[i].actual = testdata[i].data.Validate()
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
