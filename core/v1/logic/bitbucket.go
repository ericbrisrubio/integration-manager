package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type bitbucketService struct {
	companyService service.Company
	observerList   []service.Observer
	client         service.HttpClient
}

func (b bitbucketService) GetPipeline(repositoryName, username, revision, token string) (*v1.Pipeline, error) {
	panic("implement me")
}

func (b bitbucketService) GetDescriptors(repositoryName, username, revision, token, path, env string) ([]unstructured.Unstructured, error) {
	panic("implement me")
}

func (b bitbucketService) GetDirectoryContents(repositoryName, username, revision, token, path string) ([]v1.GitDirectoryContent, error) {
	panic("implement me")
}

func (b bitbucketService) CreateRepositoryWebhook(username, repositoryName, token string) (v1.GitWebhook, error) {
	panic("implement me")
}

func (b bitbucketService) DeleteRepositoryWebhookById(username, repositoryName, webhookId, token string) error {
	panic("implement me")
}

// NewBitBucketService returns Git type service
func NewBitBucketService(companyService service.Company, observerList []service.Observer, client service.HttpClient) service.Git {
	return &bitbucketService{
		companyService: companyService,
		observerList:   observerList,
		client:         client,
	}
}
