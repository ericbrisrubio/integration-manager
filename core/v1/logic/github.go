package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/klovercloud-ci-cd/integration-manager/config"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"log"
	"strconv"
	"strings"
)

type githubService struct {
	observerList []service.Observer
	client       service.HttpClient
}

func (githubService githubService) GetContent(repositoryName, username, token, path string) (v1.GitContent, error) {
	url := enums.GITHUB_API_BASE_URL + "repos/" + username + "/" + repositoryName + "/contents" + path
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Accept"] = "application/vnd.github.v3+json"
	header["cache-control"] = "no-cache"
	response, err := githubService.client.Get(url, header)
	if err != nil {
		return v1.GitContent{}, err
	}
	var gitContent v1.GitContent
	err = json.Unmarshal(response, &gitContent)

	if err != nil {
		return v1.GitContent{}, err
	}
	return gitContent, nil
}

func (githubService githubService) CreateDirectoryContent(repositoryName, username, token, path string, content v1.DirectoryContentCreatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error) {
	url := enums.GITHUB_API_BASE_URL + "repos/" + username + "/" + repositoryName + "/contents/" + path
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Accept"] = "application/vnd.github.v3+json"
	header["cache-control"] = "no-cache"
	b, err := json.Marshal(content)
	if err != nil {
		log.Println(err.Error())
		return v1.DirectoryContentCreateAndUpdateResponse{}, err
	}
	data, err := githubService.client.Put(url, header, b)
	if err != nil {
		log.Println(err.Error())
		return v1.DirectoryContentCreateAndUpdateResponse{}, err
	}
	directoryContentCreateAndUpdateResponse := v1.DirectoryContentCreateAndUpdateResponse{}
	err = json.Unmarshal(data, &directoryContentCreateAndUpdateResponse)
	if err != nil {
		log.Println(err.Error())
		return v1.DirectoryContentCreateAndUpdateResponse{}, err
	}
	return directoryContentCreateAndUpdateResponse, nil

}

func (githubService githubService) UpdateDirectoryContent(repositoryName, username, token, path string, content v1.DirectoryContentUpdatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error) {
	url := enums.GITHUB_API_BASE_URL + "repos/" + username + "/" + repositoryName + "/contents/" + path
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Accept"] = "application/vnd.github.v3+json"
	header["cache-control"] = "no-cache"
	b, err := json.Marshal(content)
	if err != nil {
		log.Println(err.Error())
		return v1.DirectoryContentCreateAndUpdateResponse{}, err
	}
	data, err := githubService.client.Put(url, header, b)
	if err != nil {
		log.Println(err.Error())
		return v1.DirectoryContentCreateAndUpdateResponse{}, err
	}
	directoryContentCreateAndUpdateResponse := v1.DirectoryContentCreateAndUpdateResponse{}
	err = json.Unmarshal(data, &directoryContentCreateAndUpdateResponse)
	if err != nil {
		log.Println(err.Error())
		return v1.DirectoryContentCreateAndUpdateResponse{}, err
	}
	return directoryContentCreateAndUpdateResponse, nil
}

func (githubService githubService) GetCommitsByBranch(username, repositoryName, branch, token string, option v1.Pagination) ([]v1.GitCommit, int64, error) {
	url := enums.GITHUB_API_BASE_URL + "repos/" + username + "/" + repositoryName + "/commits?sha=" + branch + "&per_page=" + strconv.Itoa(int(option.Limit)) + "&page=" + strconv.Itoa(int(option.Page))
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Accept"] = "application/vnd.github.v3+json"
	header["cache-control"] = "no-cache"
	response, err := githubService.client.Get(url, header)
	if err != nil {
		return nil, 0, err
	}
	var gitCommits []v1.GitCommit
	err = json.Unmarshal(response, &gitCommits)

	if err != nil {
		return nil, 0, err
	}
	return gitCommits, 0, nil
}

func (githubService githubService) GetBranches(username, repositoryName, token string) (v1.GitBranches, error) {
	githubUrl := enums.GITHUB_API_BASE_URL + "repos/" + username + "/" + repositoryName + "/branches"
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Accept"] = "application/vnd.github.v3+json"
	header["cache-control"] = "no-cache"

	response, err := githubService.client.Get(githubUrl, header)
	if err != nil {
		return nil, err
	}
	var gitBranches v1.GitBranches
	err = json.Unmarshal(response, &gitBranches)

	if err != nil {
		return nil, err
	}

	return gitBranches, nil
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

func (githubService githubService) CreateRepositoryWebhook(username, repositoryName, token string, companyId, appId, appSecret string) (v1.GitWebhook, error) {
	url := enums.GITHUB_API_BASE_URL + "repos/" + username + "/" + repositoryName + "/hooks"
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Content-Type"] = "application/json"
	header["cache-control"] = "no-cache"
	body := v1.GithubCreateWebhookRequest{
		Config: struct {
			URL         string `json:"url"`
			ContentType string `json:"content_type"`
		}{
			URL:         config.GithubWebhookConsumingUrl + "?companyId=" + companyId + "&appId=" + appId + "&appSecret=" + appSecret,
			ContentType: "json",
		},
		Events: []enums.GITHUB_EVENT{enums.GITHUB_PUSH_EVENT, enums.GITHUB_DELETE_EVENT, enums.GITHUB_RELEASE_EVENT},
	}
	b, err := json.Marshal(body)
	if err != nil {
		log.Println(err.Error())
		return v1.GitWebhook{}, err
	}
	data, err := githubService.client.Post(url, header, b)
	if err != nil {
		log.Println(err.Error())
		return v1.GitWebhook{}, err
	}
	webhook := v1.GithubWebhook{}
	webHookErrorRes := v1.GitWebHookErrorResponse{}
	err = json.Unmarshal(data, &webhook)
	if err != nil {
		log.Println(err.Error())
		return v1.GitWebhook{}, err
	}
	if webhook.URL == "" {
		err = json.Unmarshal(data, &webHookErrorRes)
		if len(webHookErrorRes.Errors) > 0 {
			return v1.GitWebhook{}, errors.New(webHookErrorRes.Errors[len(webHookErrorRes.Errors)-1].Message)
		}
		return v1.GitWebhook{}, errors.New(webHookErrorRes.Message)
	}
	return webhook.GetGitWebhook(), nil
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
	data, err := githubService.client.Get(url, header)
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
		data, err := githubService.client.Get(url, header)
		if err != nil {
			// send to observer
			return nil, err
		}

		fileAsString := string(data)[:]
		sepFiles := strings.Split(fileAsString, "---")
		for _, each := range sepFiles {
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

func (githubService githubService) GetDirectoryContents(repositoryName, username, revision, token, path string) ([]v1.GitDirectoryContent, error) {
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	url := enums.GITHUB_API_BASE_URL + "repos/" + username + "/" + repositoryName + "/contents/" + path + "?ref=" + revision
	header := make(map[string]string)
	header["Authorization"] = "token " + token
	header["Accept"] = "application/vnd.github.v3+json"
	data, err := githubService.client.Get(url, header)
	if err != nil {
		// send to observer
		return nil, err
	}
	var githubDirectoryContents []v1.GithubDirectoryContent
	err = json.Unmarshal(data, &githubDirectoryContents)
	if err != nil {
		log.Println(err.Error())
		// send to observer
		return nil, err
	}
	var gitDirectoryContents []v1.GitDirectoryContent
	for _, each := range githubDirectoryContents {
		gitDirectoryContents = append(gitDirectoryContents, each.GetGitDirectoryContent())
	}
	return gitDirectoryContents, nil
}

func (githubService githubService) notifyAll(listener v1.Subject) {
	for _, observer := range githubService.observerList {
		go observer.Listen(listener)
	}
}

// NewGithubService returns Git type service
func NewGithubService(observerList []service.Observer, client service.HttpClient) service.Git {
	return &githubService{
		observerList: observerList,
		client:       client,
	}
}
