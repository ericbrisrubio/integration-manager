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
				"accepts": "*",
				"name":    "name",
				"valid":   "true",
				"value":   "interstep",
				"message": "",
			},
			Type: map[string]string{
				"accepts": "BUILD, DEPLOY, INTERMEDIARY, JENKINS_JOB",
				"name":    "type",
				"valid":   "false",
				"value":   "build",
				"message": "invalid step type is given",
			},
			Trigger: map[string]string{
				"accepts": "AUTO, MANUAL",
				"name":    "trigger",
				"valid":   "true",
				"value":   "AUTO",
				"message": "",
			},
			Params: []map[string]string{
				{
					"accepts": "git",
					"name":    "repository_type",
					"valid":   "true",
					"value":   "git",
					"message": "",
				},
				{
					"accepts": "*",
					"name":    "revision",
					"valid":   "true",
					"value":   "myBranch",
					"message": "",
				},
			},
			Next: []map[string]string{
				{
					"accepts": "interstep, deployDev, build",
					"name":    "next",
					"valid":   "true",
					"value":   "interstep",
					"message": "",
				},
				{
					"accepts": "interstep, deployDev, build",
					"name":    "next",
					"valid":   "true",
					"value":   "deployDev",
					"message": "",
				},
			},
		},
		{
			Name: map[string]string{
				"accepts": "*",
				"name":    "name",
				"valid":   "true",
				"value":   "test2",
				"message": "",
			},
			Type: map[string]string{
				"accepts": "BUILD, DEPLOY, INTERMEDIARY, JENKINS_JOB",
				"name":    "type",
				"valid":   "true",
				"value":   "DEPLOY",
				"message": "",
			},
			Trigger: map[string]string{
				"accepts": "AUTO, MANUAL",
				"name":    "trigger",
				"valid":   "true",
				"value":   "MANUAL",
				"message": "",
			},
			Params: []map[string]string{
				{
					"accepts": "git",
					"name":    "repository_type",
					"valid":   "false",
					"value":   "anyType",
					"message": "invalid step param is given",
				},
			},
			Next: []map[string]string{
				{
					"accepts": "interstep, deployDev, build",
					"name":    "next",
					"valid":   "true",
					"value":   "build",
					"message": "",
				},
				{
					"accepts": "interstep, deployDev, build",
					"name":    "next",
					"valid":   "false",
					"value":   "stop",
					"message": "invalid step next is given",
				},
			},
		},
		{
			Name: map[string]string{
				"accepts": "*",
				"name":    "name",
				"valid":   "true",
				"value":   "test3",
				"message": "",
			},
			Type: map[string]string{
				"accepts": "BUILD, DEPLOY, INTERMEDIARY, JENKINS_JOB",
				"name":    "type",
				"valid":   "true",
				"value":   "BUILD",
				"message": "",
			},
			Trigger: map[string]string{
				"accepts": "AUTO, MANUAL",
				"name":    "trigger",
				"valid":   "false",
				"value":   "ssss",
				"message": "invalid step trigger is given",
			},
			Params: []map[string]string{
				{
					"accepts": "*",
					"name":    "service_account",
					"valid":   "true",
					"value":   "sa",
					"message": "",
				},
			},
			Next: []map[string]string{
				{
					"accepts": "interstep, deployDev, build",
					"name":    "next",
					"valid":   "false",
					"value":   "stop",
					"message": "invalid step next is given",
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

		err := true
		for key1, val1 := range testdata[i].actual.Name {
			err = true
			for key2, val2 := range testdata[i].expected.Name {
				if key1 == key2 && val1 == val2 {
					err = false
					break
				}
			}
			if err {
				assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
			}
		}

		err = true
		for key1, val1 := range testdata[i].actual.Type {
			err = true
			for key2, val2 := range testdata[i].expected.Type {
				if key1 == key2 && val1 == val2 {
					err = false
					break
				}
			}
			if err {
				assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
			}
		}
		if err {
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}

		err = true
		for key1, val1 := range testdata[i].actual.Trigger {
			err = true
			for key2, val2 := range testdata[i].expected.Trigger {
				if key1 == key2 && val1 == val2 {
					err = false
					break
				}
			}
			if err {
				assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
			}
		}
		if err {
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}

		err = true
		for _, eachActualMap := range testdata[i].actual.Params {
			err = true
			for key1, val1 := range eachActualMap {
				err = true
				for _, eachExpectedMap := range testdata[i].expected.Params {
					for key2, val2 := range eachExpectedMap {
						if key1 == key2 && val1 == val2 {
							err = false
							break
						}
					}
				}
				if err {
					assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
				}
			}
			if err {
				assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
			}
		}
		if err {
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}

		err = true
		for _, eachActualMap := range testdata[i].actual.Next {
			err = true
			for key1, val1 := range eachActualMap {
				err = true
				for _, eachExpectedMap := range testdata[i].expected.Next {
					for key2, val2 := range eachExpectedMap {
						if key1 == key2 && val1 == val2 {
							err = false
							break
						}
					}
				}
				if err {
					assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
				}
			}
			if err {
				assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
			}
		}
		if err {
			assert.ElementsMatch(t, testdata[i].expected, testdata[i].actual)
		}
	}
}
