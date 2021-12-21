package v1

import (
	"errors"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"reflect"
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
