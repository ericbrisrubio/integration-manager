package v1

import (
	"errors"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"reflect"
	"strings"
)

// Step contains pipeline step info
type Step struct {
	Name        string                       `json:"name" yaml:"name"`
	Type        enums.STEP_TYPE              `json:"type" yaml:"type"`
	Trigger     enums.TRIGGER                `json:"trigger" yaml:"trigger"`
	Params      map[enums.PARAMS]string      `json:"params" yaml:"params"`
	Next        []string                     `json:"next" yaml:"next"`
	Descriptors *[]unstructured.Unstructured `json:"descriptors" yaml:"descriptors"`
}

// StepForValidation contains pipeline step info for validation
type StepForValidation struct {
	Name    map[string]string   `json:"name" yaml:"name"`
	Type    map[string]string   `json:"type" yaml:"type"`
	Trigger map[string]string   `json:"trigger" yaml:"trigger"`
	Params  []map[string]string `json:"params" yaml:"params"`
	Next    []map[string]string `json:"next" yaml:"next"`
}

// Validate validates pipeline step
func (step Step) Validate() error {
	if step.Name == "" {
		return errors.New("step name is required")
	} else if len(step.Name) > 16 {
		return errors.New("step name length cannot be more than 16 character")
	} else {
		for i := 0; i < len(step.Name); i++ {
			if (step.Name[i] < 97 || step.Name[i] > 122) && (step.Name[i] < 48 || step.Name[i] > 57) {
				return errors.New("step name can only contain lower case characters or digits")
			}
		}
	}
	keys := reflect.ValueOf(step.Params).MapKeys()
	for i := 0; i < len(keys); i++ {
		if step.Params[enums.PARAMS(keys[i].String())] == "" {
			return errors.New("step params is missing")
		}
	}
	if step.Type == enums.BUILD || step.Type == enums.DEPLOY {
		if step.Trigger == enums.AUTO || step.Trigger == enums.MANUAL {
			return nil
		} else if step.Trigger == "" {
			return errors.New("step trigger is required")
		} else {
			return errors.New("step trigger is invalid")
		}
	} else if step.Type == "" {
		return errors.New("step type is required")
	} else {
		return errors.New("step type is invalid")
	}
}

// GetStepForValidationFromStep gets StepForValidation object from Step object
func (step Step) GetStepForValidationFromStep(stepNameMap map[string]bool) StepForValidation {
	var stepForValidation StepForValidation
	stepForValidation.Name = step.GetNameWithValidation()
	stepForValidation.Type = step.GetTypeWithValidation()
	stepForValidation.Trigger = step.GetTriggerWithValidation()
	stepForValidation.Params = step.GetParamsWithValidation()
	stepForValidation.Next = step.GetNextWithValidation(stepNameMap)
	return stepForValidation
}

func (step Step) GetNameWithValidation() map[string]string {
	nameMap := make(map[string]string)
	nameMap["name"] = "name"
	nameMap["value"] = step.Name
	nameMap["accepts"] = "*"
	for i := 0; i < len(step.Name); i++ {
		if (step.Name[i] < 97 || step.Name[i] > 122) && (step.Name[i] < 48 || step.Name[i] > 57) {
			nameMap["valid"] = "false"
			nameMap["message"] = "step name can only contain lower case characters or digits"
		}
	}
	if step.Name == "" {
		nameMap["valid"] = "false"
		nameMap["message"] = "step name is missing"
	} else if len(step.Name) > 16 {
		nameMap["valid"] = "false"
		nameMap["message"] = "step name length cannot be more than 16 character"
	} else {
		nameMap["valid"] = "true"
		nameMap["message"] = ""
	}
	return nameMap
}

func (step Step) GetTypeWithValidation() map[string]string {
	typeMap := make(map[string]string)
	typeMap["name"] = "type"
	typeMap["value"] = string(step.Type)
	typeMap["accepts"] = string(enums.BUILD + ", " + enums.DEPLOY + ", " + enums.INTERMEDIARY + ", " + enums.JENKINS_JOB)
	if step.Type == "" {
		typeMap["valid"] = "false"
		typeMap["message"] = "step type is missing"
	} else if val, _ := typeMap["value"]; val == string(enums.BUILD) || val == string(enums.DEPLOY) || val == string(enums.INTERMEDIARY) || val == string(enums.JENKINS_JOB) {
		typeMap["valid"] = "true"
		typeMap["message"] = ""
	} else {
		typeMap["valid"] = "false"
		typeMap["message"] = "invalid step type"
	}
	return typeMap
}

func (step Step) GetTriggerWithValidation() map[string]string {
	triggerMap := make(map[string]string)
	triggerMap["name"] = "trigger"
	triggerMap["value"] = string(step.Trigger)
	triggerMap["accepts"] = string(enums.AUTO + ", " + enums.MANUAL)
	if step.Trigger == "" {
		triggerMap["valid"] = "false"
		triggerMap["message"] = "trigger is missing"
	} else if val, _ := triggerMap["value"]; val == string(enums.AUTO) || val == string(enums.MANUAL) {
		triggerMap["valid"] = "true"
		triggerMap["message"] = ""
	} else {
		triggerMap["valid"] = "false"
		triggerMap["message"] = "invalid trigger"
	}
	return triggerMap
}

func (step Step) GetParamsWithValidation() []map[string]string {
	var paramsMap []map[string]string
	for key, val := range step.Params {
		paramMap := make(map[string]string)
		paramMap["name"] = string(key)
		paramMap["value"] = val
		if key == enums.REPOSITORY_TYPE_PARAM {
			paramMap["accepts"] = "git"
		} else {
			paramMap["accepts"] = "*"
		}
		if val == "" {
			paramMap["valid"] = "false"
			paramMap["message"] = "param is missing"
		} else if acceptValue, _ := paramMap["accepts"]; acceptValue == "*" || val == acceptValue {
			paramMap["valid"] = "true"
			paramMap["message"] = ""
		} else {
			paramMap["valid"] = "false"
			paramMap["message"] = "invalid param"
		}
		paramsMap = append(paramsMap, paramMap)
	}
	return paramsMap
}

func (step Step) GetNextWithValidation(stepNameMap map[string]bool) []map[string]string {
	var nextMaps []map[string]string
	var accept string
	for key, _ := range stepNameMap {
		accept = accept + key + ", "
	}
	accept = strings.TrimSuffix(accept, ", ")
	for _, each := range step.Next {
		nextMap := make(map[string]string)
		nextMap["name"] = "next"
		nextMap["value"] = each
		nextMap["accepts"] = accept
		if _, ok := stepNameMap[each]; ok {
			nextMap["valid"] = "true"
			nextMap["message"] = ""
		} else {
			nextMap["valid"] = "false"
			nextMap["message"] = "invalid step"
		}
		nextMaps = append(nextMaps, nextMap)
	}
	return nextMaps
}
