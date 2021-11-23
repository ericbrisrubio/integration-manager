package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type gitlabService struct {
	companyService service.Company
	observerList   []service.Observer
	client         service.HttpClient
}

func (g gitlabService) GetPipeline(repositoryName, username, revision, token string) (*v1.Pipeline, error) {
	panic("implement me")
}

func (g gitlabService) GetDescriptors(repositoryName, username, revision, token, path, env string) ([]unstructured.Unstructured, error) {
	panic("implement me")
}

func (g gitlabService) GetDirectoryContents(repositoryName, username, revision, token, path string) ([]v1.GitDirectoryContent, error) {
	panic("implement me")
}

func (g gitlabService) CreateRepositoryWebhook(username, repositoryName, token string) (v1.GitWebhook, error) {
	panic("implement me")
}

func (g gitlabService) DeleteRepositoryWebhookById(username, repositoryName, webhookId, token string) error {
	panic("implement me")
}

// NewGitlabService returns Git type service
func NewGitlabService(companyService service.Company, observerList []service.Observer, client service.HttpClient) service.Git {
	return &gitlabService{
		companyService: companyService,
		observerList:   observerList,
		client:         client,
	}
}
