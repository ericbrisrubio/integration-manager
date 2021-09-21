package v1

import "github.com/klovercloud-ci/enums"

type Step struct {
	Name string `json:"name" yaml:"name"`
	Type enums.STEP_TYPE `json:"type" yaml:"type"`
	ServiceAccount string `json:"service_account" yaml:"service_account"`
	Input Resource `json:"input"  yaml:"input"`
	Outputs []Resource `json:"outputs"  yaml:"outputs"`
	Arg Variable  `json:"arg"  yaml:"arg"`
	Env Variable  `json:"env"  yaml:"env"`
}