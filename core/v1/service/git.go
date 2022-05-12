package service

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Git Git related operations.
type Git interface {
	GetPipeline(repositoryName, username, revision, token string) (*v1.Pipeline, error)
	GetDescriptors(repositoryName, username, revision, token, path, env string) ([]unstructured.Unstructured, error)
	GetDirectoryContents(repositoryName, username, revision, token, path string) ([]v1.GitDirectoryContent, error)
	CreateRepositoryWebhook(username, repositoryName, token string, companyId string) (v1.GitWebhook, error)
	DeleteRepositoryWebhookById(username, repositoryName, webhookId, token string) error
	GetBranches(username, repositoryName, token string) (v1.GitBranches, error)
	GetCommitByBranch(username, repositoryName, branch, token string) (v1.GitCommit, error)
}
