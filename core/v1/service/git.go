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
	GetContent(repositoryName, username, token, path string) (v1.GitContent, error)
	CreateDirectoryContent(repositoryName, username, token, path string, content v1.DirectoryContentCreatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error)
	UpdateDirectoryContent(repositoryName, username, token, path string, content v1.DirectoryContentUpdatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error)
	CreateRepositoryWebhook(username, repositoryName, token string, companyId, appId string) (v1.GitWebhook, error)
	DeleteRepositoryWebhookById(username, repositoryName, webhookId, token string) error
	GetBranches(username, repositoryName, token string) (v1.GitBranches, error)
	GetCommitsByBranch(username, repositoryName, branch, token string, option v1.Pagination) ([]v1.GitCommit, int64, error)
}
