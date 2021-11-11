package service

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Git Git related operations.
type Git interface {
	GetPipeline(repositoryName, username, revision, token string) (*v1.Pipeline, error)
	GetDescriptors(repositoryName, username, revision, token, path, env string) ([]unstructured.Unstructured, error)
	GetDirectoryContents(repositoryName, username, revision, token, path string) ([]v1.GithubDirectoryContent, error)
	CreateRepositoryWebhook(username, repositoryName, token string) (v1.GithubWebhook, error)
	DeleteRepositoryWebhookById(username, repositoryName, webhookId, token string) error
}
