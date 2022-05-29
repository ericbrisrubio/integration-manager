package service

import v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"

type Pipeline interface {
	GetPipelineForValidation(companyId, repositoryId, url, revision string) (v1.PipelineForValidation, error)
}
