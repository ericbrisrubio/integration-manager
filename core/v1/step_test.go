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

// Test function for testing step for validation data - name, type, trigger, params
func TestStep_GetStepForValidationFromStep(t *testing.T) {
	type TestCase struct {
		data     Step
		expected StepForValidation
		actual   StepForValidation
	}

	var testdata []TestCase
	name := []string{"interstep", "test2", "test3"}
	typeD := []enums.STEP_TYPE{"build", "DEPLOY", "BUILD"}
	trigger := []enums.TRIGGER{"AUTO", "MANUAL", "ssss"}
	params := []map[enums.PARAMS]string{{"repository_type": "git", "revision": "myBranch"}, {"repository_type": "anyType"}, {"service_account": "sa"}}
	next := [][]string{{"interstep", "deployDev"}, {"build", "stop"}, {"stop"}}
	stepNameMap := map[string]bool{"interstep": true, "deployDev": true, "build": true}
	expec := []StepForValidation{
		{
			Name: map[string]string{
				"accept":   "*",
				"name":     "name",
				"validate": "true",
				"value":    "interstep",
			},
			Type: map[string]string{
				"accept":   "BUILD/DEPLOY/INTERMEDIARY/JENKINS_JOB",
				"name":     "type",
				"validate": "false",
				"value":    "build",
			},
			Trigger: map[string]string{
				"accept":   "AUTO/MANUAL",
				"name":     "trigger",
				"validate": "true",
				"value":    "AUTO",
			},
			Params: []map[string]string{
				{
					"accept":   "git",
					"name":     "repository_type",
					"validate": "true",
					"value":    "git",
				},
				{
					"accept":   "*",
					"name":     "revision",
					"validate": "true",
					"value":    "myBranch",
				},
			},
			Next: []map[string]string{
				{
					"accept":   "interstep/deployDev/build",
					"name":     "next",
					"validate": "true",
					"value":    "interstep",
				},
				{
					"accept":   "interstep/deployDev/build",
					"name":     "next",
					"validate": "true",
					"value":    "deployDev",
				},
			},
		},
		{
			Name: map[string]string{
				"accept":   "*",
				"name":     "name",
				"validate": "true",
				"value":    "test2",
			},
			Type: map[string]string{
				"accept":   "BUILD/DEPLOY/INTERMEDIARY/JENKINS_JOB",
				"name":     "type",
				"validate": "true",
				"value":    "DEPLOY",
			},
			Trigger: map[string]string{
				"accept":   "AUTO/MANUAL",
				"name":     "trigger",
				"validate": "true",
				"value":    "MANUAL",
			},
			Params: []map[string]string{
				{
					"accept":   "git",
					"name":     "repository_type",
					"validate": "false",
					"value":    "anyType",
				},
			},
			Next: []map[string]string{
				{
					"accept":   "interstep/deployDev/build",
					"name":     "next",
					"validate": "true",
					"value":    "build",
				},
				{
					"accept":   "interstep/deployDev/build",
					"name":     "next",
					"validate": "false",
					"value":    "stop",
				},
			},
		},
		{
			Name: map[string]string{
				"accept":   "*",
				"name":     "name",
				"validate": "true",
				"value":    "test3",
			},
			Type: map[string]string{
				"accept":   "BUILD/DEPLOY/INTERMEDIARY/JENKINS_JOB",
				"name":     "type",
				"validate": "true",
				"value":    "BUILD",
			},
			Trigger: map[string]string{
				"accept":   "AUTO/MANUAL",
				"name":     "trigger",
				"validate": "false",
				"value":    "ssss",
			},
			Params: []map[string]string{
				{
					"accept":   "*",
					"name":     "service_account",
					"validate": "true",
					"value":    "sa",
				},
			},
			Next: []map[string]string{
				{
					"accept":   "interstep/deployDev/build",
					"name":     "next",
					"validate": "false",
					"value":    "stop",
				},
			},
		},
	}

	for i := 0; i < len(name); i++ {
		testcase := TestCase{
			data: Step{
				Name:    name[i],
				Type:    typeD[i],
				Trigger: trigger[i],
				Params:  params[i],
				Next:    next[i],
			},
			expected: expec[i],
		}
		testdata = append(testdata, testcase)
	}
	for i := 0; i < len(name); i++ {
		testdata[i].actual = testdata[i].data.GetStepForValidationFromStep(stepNameMap)
		if !reflect.DeepEqual(testdata[i].expected, testdata[i].actual) {
			fmt.Println(testdata[i].actual)
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
