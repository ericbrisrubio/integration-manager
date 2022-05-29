package v1

import (
	"errors"
)

// Pipeline contains pipeline data
type Pipeline struct {
	MetaData   PipelineMetadata `json:"_metadata" yaml:"_metadata"`
	ApiVersion string           `json:"api_version" yaml:"api_version"`
	Name       string           `json:"name"  yaml:"name"`
	ProcessId  string           `json:"process_id" yaml:"process_id"`
	Steps      []Step           `json:"steps" yaml:"steps"`
}

// Pipeline contains pipeline data for validation
type PipelineForValidation struct {
	Name  string              `json:"name"  yaml:"name"`
	Steps []StepForValidation `json:"steps" yaml:"steps"`
}

// Validate validates pipeline data
func (pipeline Pipeline) Validate() error {
	if pipeline.ApiVersion == "" {
		return errors.New("Api version is required!")
	}
	if pipeline.Name == "" {
		return errors.New("Pipeline name is required!")
	}
	if pipeline.ProcessId == "" {
		return errors.New("Pipeline process id is required!")
	}
	for _, each := range pipeline.Steps {
		err := each.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (pipeline Pipeline) GetStepNameMap() map[string]bool {
	stepNameMap := make(map[string]bool)
	for _, each := range pipeline.Steps {
		stepNameMap[each.Name] = true
	}
	return stepNameMap
}

func (pipeline Pipeline) GetPipelineForValidationFromPipeline() PipelineForValidation {
	var pipelineForValidation PipelineForValidation
	pipelineForValidation.Name = pipeline.Name
	var stepsForValidations []StepForValidation
	stepNameMap := pipeline.GetStepNameMap()
	for _, each := range pipeline.Steps {
		stepForValidation := each.GetStepForValidationFromStep(stepNameMap)
		stepsForValidations = append(stepsForValidations, stepForValidation)
	}
	pipelineForValidation.Steps = stepsForValidations
	return pipelineForValidation
}
