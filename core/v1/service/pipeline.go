package service

import v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"

type Pipeline interface {
	GetPipelineForValidation(companyId, repositoryId, url, revision string) (v1.PipelineForValidation, error)
	Create(companyId, repositoryId, url string, payload v1.DirectoryContentCreatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error)
	Update(companyId, repositoryId, url string, payload v1.DirectoryContentUpdatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error)
}
