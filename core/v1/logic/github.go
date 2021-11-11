package logic

import (
	"encoding/json"
	"fmt"
	"github.com/klovercloud-ci/config"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"log"
	"strings"
)

type githubService struct {
	companyService service.Company
	observerList   []service.Observer
	client         service.HttpClient
}

func (githubService githubService) DeleteRepositoryWebhookById(username, repositoryName, webhookId, token string) error {
	url := enums.GITHUB_API_BASE_URL + "repos/" + username + "/" + repositoryName + "/hooks/" + webhookId
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Accept"] = "application/vnd.github.v3+json"
	header["cache-control"] = "no-cache"
	err := githubService.client.Delete(url, header)
	if err != nil {
		return err
	}
	return nil
}

func (githubService githubService) CreateRepositoryWebhook(username, repositoryName, token string) (v1.GithubWebhook, error) {
	url := enums.GITHUB_API_BASE_URL + "repos/" + username + "/" + repositoryName + "/hooks"
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Content-Type"] = "application/json"
	header["cache-control"] = "no-cache"
	body := v1.GithubCreateWebhookRequest{
		Config: struct {
			URL string `json:"url"`
		}{
			URL: config.GithubWebhookConsumingUrl,
		},
		Events: []enums.GIT_EVENT{enums.PUSH, enums.DELETE, enums.RELEASE},
	}
	b, err := json.Marshal(body)
	if err != nil {
		log.Println(err.Error())
		return v1.GithubWebhook{}, err
	}
	err, data := githubService.client.Post(url, header, b)
	if err != nil {
		log.Println(err.Error())
		return v1.GithubWebhook{}, err
	}
	webhook := v1.GithubWebhook{}
	err = json.Unmarshal(data, &webhook)
	if err != nil {
		log.Println(err.Error())
		return v1.GithubWebhook{}, err
	}
	return webhook, nil
}

func (githubService githubService) GetPipeline(repositoryName, username, revision, token string) (*v1.Pipeline, error) {
	contents, err := githubService.GetDirectoryContents(repositoryName, username, revision, token, enums.PIPELINE_FILE_BASE_DIRECTORY)
	if err != nil {
		return nil, err
	}
	var pipelneFile string

	for _, each := range contents {
		if each.Name == enums.PIPELINE_FILE_NAME+".yaml" || each.Name == enums.PIPELINE_FILE_NAME+".yml" || each.Name == enums.PIPELINE_FILE_NAME+".json" {
			pipelneFile = each.Name
			break
		}
	}
	url := enums.GITHUB_RAW_CONTENT_BASE_URL + username + "/" + repositoryName + "/" + revision + "/klovercloud/pipeline/" + pipelneFile
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Accept"] = "application/vnd.github.v3.raw"
	header["X-Requested-With"] = "Curl"
	err, data := githubService.client.Get(url, header)
	if err != nil {
		// send to observer
		return nil, err
	}
	pipeline := v1.Pipeline{}
	if strings.HasSuffix(pipelneFile, "yaml") || strings.HasSuffix(pipelneFile, "yml") {
		err = yaml.Unmarshal(data, &pipeline)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	} else {
		err = json.Unmarshal(data, &pipeline)
		if err != nil {
			log.Println(err.Error())

			return nil, err
		}
	}

	return &pipeline, nil
}

func (githubService githubService) GetDescriptors(repositoryName, username, revision, token, path, env string) ([]unstructured.Unstructured, error) {
	contents, err := githubService.GetDirectoryContents(repositoryName, username, revision, token, path)
	if err != nil {
		return nil, err
	}
	var files []unstructured.Unstructured

	for _, each := range contents {
		if each.Type != "file" {
			continue
		}
		if each.Name != env+".yaml" && each.Name != env+".yml" && each.Name != env+".json" {
			continue
		}
		url := fmt.Sprint(each.DownloadURL)
		header := make(map[string]string)
		header["Authorization"] = "token " + token
		header["Accept"] = "application/vnd.github.v3.raw"
		header["X-Requested-With"] = "Curl"
		err, data := githubService.client.Get(url, header)
		if err != nil {
			// send to observer
			return nil, err
		}

		fileAsString := string(data)[:]
		sepfiles := strings.Split(fileAsString, "---")
		for _, each := range sepfiles {
			obj := &unstructured.Unstructured{
				Object: map[string]interface{}{},
			}
			if err := yaml.Unmarshal([]byte(each), &obj.Object); err != nil {
				log.Println(err.Error())
				if err := json.Unmarshal([]byte(each), &obj.Object); err != nil {
					log.Println(err.Error())
					return nil, err
				}
			}
			files = append(files, *obj)
		}
	}
	return files, nil
}

func (githubService githubService) GetDirectoryContents(repositoryName, username, revision, token, path string) ([]v1.GithubDirectoryContent, error) {
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	url := enums.GITHUB_API_BASE_URL + "repos/" + username + "/" + repositoryName + "/contents/" + path + "?ref=" + revision
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Accept"] = "application/vnd.github.v3+json"
	err, data := githubService.client.Get(url, header)
	if err != nil {
		// send to observer
		return nil, err
	}
	contents := []v1.GithubDirectoryContent{}
	err = json.Unmarshal(data, &contents)
	if err != nil {
		log.Println(err.Error())
		// send to observer
		return nil, err
	}
	return contents, nil
}

func (githubService githubService) notifyAll(listener v1.Subject) {
	for _, observer := range githubService.observerList {
		go observer.Listen(listener)
	}
}

// NewGithubService returns Git type service
func NewGithubService(companyService service.Company, observerList []service.Observer, client service.HttpClient) service.Git {
	return &githubService{
		companyService: companyService,
		observerList:   observerList,
		client:         client,
	}
}
