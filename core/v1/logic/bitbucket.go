package logic

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/klovercloud-ci-cd/integration-manager/config"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"log"
	"strings"
)

type bitbucketService struct {
	companyService service.Company
	observerList   []service.Observer
	client         service.HttpClient
}

func (b bitbucketService) GetPipeline(repositoryName, username, revision, token string) (*v1.Pipeline, error) {
	contents, err := b.GetDirectoryContents(repositoryName, username, revision, token, enums.PIPELINE_FILE_BASE_DIRECTORY)
	if err != nil {
		return nil, err
	}
	var pipelneFile string
	for _, each := range contents {
		split := strings.Split(each.Path, "/")
		if split[len(split)-1] == enums.PIPELINE_FILE_NAME+".yaml" || split[len(split)-1] == enums.PIPELINE_FILE_NAME+".yml" || split[len(split)-1] == enums.PIPELINE_FILE_NAME+".json" {
			pipelneFile = split[len(split)-1]
			break
		}
	}
	//raw_file_content:-https://api.bitbucket.org/2.0/repositories/shahidul34/abc/src/0e6724ff42018ae42ce0ae3b85f131bf7b10196e/README.md
	url := enums.BITBUCKET_API_BASE_URL + "repositories/" + username + "/" + repositoryName + "/src/" + revision + "/" + "klovercloud/pipeline/" + pipelneFile
	base64ConvertedToken := base64.StdEncoding.EncodeToString([]byte(username + ":" + token))
	header := make(map[string]string)
	header["Authorization"] = "Basic " + base64ConvertedToken
	header["Content-Type"] = "application/json"
	data, err := b.client.Get(url, header)
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

func (b bitbucketService) GetDescriptors(repositoryName, username, revision, token, path, env string) ([]unstructured.Unstructured, error) {
	contents, err := b.GetDirectoryContents(repositoryName, username, revision, token, path)
	if err != nil {
		return nil, err
	}
	var files []unstructured.Unstructured
	for _, each := range contents {
		if each.Type != "file" {
			continue
		}
		split := strings.Split(each.Path, "/")
		if split[len(split)-1] != env+".yaml" && split[len(split)-1] != env+".yml" && split[len(split)-1] != env+".json" {
			continue
		}
		url := fmt.Sprint(each.DownloadURL)
		base64ConvertedToken := base64.StdEncoding.EncodeToString([]byte(username + ":" + token))
		header := make(map[string]string)
		header["Authorization"] = "Basic " + base64ConvertedToken
		header["Content-Type"] = "application/json"
		data, err := b.client.Get(url, header)
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

func (b bitbucketService) GetDirectoryContents(repositoryName, username, revision, token, path string) ([]v1.GitDirectoryContent, error) {
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	base64ConvertedToken := base64.StdEncoding.EncodeToString([]byte(username + ":" + token))
	url := enums.BITBUCKET_API_BASE_URL + "repositories/" + username + "/" + repositoryName + "/src/" + revision + "/" + path + "?pagelen=10"
	log.Println(url)
	header := make(map[string]string)
	header["Authorization"] = "Basic " + base64ConvertedToken
	header["Content-Type"] = "application/json"
	data, err := b.client.Get(url, header)
	if err != nil {
		// send to observer
		return nil, err
	}
	var bitBucketDirectoryContents v1.BitbucketDirectoryContent
	err = json.Unmarshal(data, &bitBucketDirectoryContents)
	if err != nil {
		log.Println(err.Error())
		// send to observer
		return nil, err
	}
	var gitDirectoryContents []v1.GitDirectoryContent
	for _, each := range bitBucketDirectoryContents.Values {
		gitDirectoryContent := v1.BitbucketDirectoryContent{}
		gitDirectoryContent.Values = append(gitDirectoryContent.Values, each)
		gitDirectoryContents = append(gitDirectoryContents, gitDirectoryContent.GetGitDirectoryContent())
	}
	return gitDirectoryContents, nil
}

func (b bitbucketService) CreateRepositoryWebhook(username, repositoryName, token string, companyId string) (v1.GitWebhook, error) {
	base64ConvertedToken := base64.StdEncoding.EncodeToString([]byte(username + ":" + token))
	url := enums.BITBUCKET_API_BASE_URL + "repositories/" + username + "/" + repositoryName + "/hooks"
	header := make(map[string]string)
	header["Authorization"] = "Basic " + base64ConvertedToken
	header["Content-Type"] = "application/json"

	body := v1.BitbucketCreateWebhookRequest{
		Description: "Bitbucket webhook",
		URL:         config.BitbucketWebhookConsumingUrl + "?companyId=" + companyId,
		Active:      true,
		Events:      []string{"repo:" + string(enums.BITBUCKET_PUSH_EVENT)},
	}
	data, err := json.Marshal(body)
	if err != nil {
		log.Println(err.Error())
		return v1.GitWebhook{}, err
	}
	data, err = b.client.Post(url, header, data)
	if err != nil {
		log.Println(err.Error())
		return v1.GitWebhook{}, err
	}
	webhook := v1.BitbucketWebhook{}
	err = json.Unmarshal(data, &webhook)
	if err != nil {
		log.Println(err.Error())
		return v1.GitWebhook{}, err
	}
	return webhook.GetGitWebhook(), nil
}

func (b bitbucketService) DeleteRepositoryWebhookById(username, repositoryName, webhookId, token string) error {
	base64ConvertedToken := base64.StdEncoding.EncodeToString([]byte(username + ":" + token))
	url := enums.BITBUCKET_API_BASE_URL + "repositories/" + username + "/" + repositoryName + "/hooks/" + webhookId
	header := make(map[string]string)
	header["Authorization"] = "Basic " + base64ConvertedToken
	header["Content-Type"] = "application/json"
	err := b.client.Delete(url, header)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (b bitbucketService) notifyAll(listener v1.Subject) {
	for _, observer := range b.observerList {
		go observer.Listen(listener)
	}
}

// NewBitBucketService returns Git type service
func NewBitBucketService(companyService service.Company, observerList []service.Observer, client service.HttpClient) service.Git {
	return &bitbucketService{
		companyService: companyService,
		observerList:   observerList,
		client:         client,
	}
}
