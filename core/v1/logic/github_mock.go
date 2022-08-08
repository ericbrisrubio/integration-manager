package logic

import (
	"encoding/json"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"log"
	"strings"
)

type githubMockService struct {
	observerList []service.Observer
	client       service.HttpClient
}

func (g githubMockService) GetCommitsByBranch(username, repositoryName, branch, token string, option v1.Pagination) ([]v1.GitCommit, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (g githubMockService) GetContent(repositoryName, username, token, path string) (v1.GitContent, error) {
	//TODO implement me
	panic("implement me")
}

func (g githubMockService) CreateDirectoryContent(repositoryName, username, token, path string, content v1.DirectoryContentCreatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g githubMockService) UpdateDirectoryContent(repositoryName, username, token, path string, content v1.DirectoryContentUpdatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g githubMockService) GetCommitByBranch(username, repositoryName, branch, token string) (v1.GitCommit, error) {
	//TODO implement me
	panic("implement me")
}

func (g githubMockService) GetBranches(username, repositoryName, token string) (v1.GitBranches, error) {
	//TODO implement me
	panic("implement me")
}

func (g githubMockService) GetPipeline(repositoryName, username, revision, token string) (*v1.Pipeline, error) {
	contents, err := g.GetDirectoryContents(repositoryName, username, revision, token, enums.PIPELINE_FILE_BASE_DIRECTORY)
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
	data := getPipelineFile()
	pipeline := v1.Pipeline{}
	if strings.HasSuffix(pipelneFile, "yaml") || strings.HasSuffix(pipelneFile, "yml") {
		err = yaml.Unmarshal([]byte(data), &pipeline)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(data), &pipeline)
		if err != nil {
			log.Println(err.Error())

			return nil, err
		}
	}

	return &pipeline, nil
}

func (g githubMockService) GetDescriptors(repositoryName, username, revision, token, path, env string) ([]unstructured.Unstructured, error) {
	contents, err := g.GetDirectoryContents(repositoryName, username, revision, token, path)
	if err != nil {
		return nil, err
	}
	var files []unstructured.Unstructured
	for _, each := range contents {
		if each.Type != "file" {
			continue
		}
		data := getDescriptorFile()
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

func (g githubMockService) GetDirectoryContents(repositoryName, username, revision, token, path string) ([]v1.GitDirectoryContent, error) {
	data := DirectoryContents()
	var githubDirectoryContents []v1.GithubDirectoryContent
	err := json.Unmarshal(data, &githubDirectoryContents)
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

func (g githubMockService) CreateRepositoryWebhook(username, repositoryName, token string, companyId, appId, appSecret string) (v1.GitWebhook, error) {
	//TODO implement me
	panic("implement me")
}

func (g githubMockService) DeleteRepositoryWebhookById(username, repositoryName, webhookId, token string) error {
	//TODO implement me
	panic("implement me")
}

func DirectoryContents() []byte {
	data := `[
				  {
					"name": "configs",
					"path": "klovercloud/pipeline/configs",
					"sha": "d5391ad07f83adfb1fbba0463527484d11d5b9f4",
					"size": 0,
					"url": "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/configs?ref=aed348fe49eb664af19d38b08036262bad65aba9",
					"html_url": "https://github.com/flameOfDimitry/TestApp/tree/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/configs",
					"git_url": "https://api.github.com/repos/flameOfDimitry/TestApp/git/trees/d5391ad07f83adfb1fbba0463527484d11d5b9f4",
					"download_url": null,
					"type": "dir",
					"_links": {
					  "self": "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/configs?ref=aed348fe49eb664af19d38b08036262bad65aba9",
					  "git": "https://api.github.com/repos/flameOfDimitry/TestApp/git/trees/d5391ad07f83adfb1fbba0463527484d11d5b9f4",
					  "html": "https://github.com/flameOfDimitry/TestApp/tree/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/configs"
					}
				  },
				  {
					"name": "pipeline.yaml",
					"path": "klovercloud/pipeline/pipeline.yaml",
					"sha": "4e9658a477df11cf6619305b5a1ac897e1b2e42f",
					"size": 1122,
					"url": "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/pipeline.yaml?ref=aed348fe49eb664af19d38b08036262bad65aba9",
					"html_url": "https://github.com/flameOfDimitry/TestApp/blob/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml",
					"git_url": "https://api.github.com/repos/flameOfDimitry/TestApp/git/blobs/4e9658a477df11cf6619305b5a1ac897e1b2e42f",
					"download_url": "https://raw.githubusercontent.com/flameOfDimitry/TestApp/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml",
					"type": "file",
					"_links": {
					  "self": "https://api.github.com/repos/flameOfDimitry/TestApp/contents/klovercloud/pipeline/pipeline.yaml?ref=aed348fe49eb664af19d38b08036262bad65aba9",
					  "git": "https://api.github.com/repos/flameOfDimitry/TestApp/git/blobs/4e9658a477df11cf6619305b5a1ac897e1b2e42f",
					  "html": "https://github.com/flameOfDimitry/TestApp/blob/aed348fe49eb664af19d38b08036262bad65aba9/klovercloud/pipeline/pipeline.yaml"
					}
				  }
			]`
	var directoryContent []byte
	directoryContent = append(directoryContent, []byte(data)...)

	return directoryContent
}

// NewGithubMockService returns Git type service
func NewGithubMockService(observerList []service.Observer, client service.HttpClient) service.Git {
	return &githubMockService{
		observerList: observerList,
		client:       client,
	}
}
