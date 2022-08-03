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

type bitbucketMockService struct {
	observerList []service.Observer
	client       service.HttpClient
}

func (b bitbucketMockService) GetCommitsByBranch(username, repositoryName, branch, token string, option v1.Pagination) ([]v1.GitCommit, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (b bitbucketMockService) GetContent(repositoryName, username, token, path string) (v1.GitContent, error) {
	//TODO implement me
	panic("implement me")
}

func (b bitbucketMockService) CreateDirectoryContent(repositoryName, username, token, path string, content v1.DirectoryContentCreatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b bitbucketMockService) UpdateDirectoryContent(repositoryName, username, token, path string, content v1.DirectoryContentUpdatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b bitbucketMockService) GetCommitByBranch(username, repositoryName, branch, token string) (v1.GitCommit, error) {
	//TODO implement me
	panic("implement me")
}

func (b bitbucketMockService) GetBranches(username, repositoryName, token string) (v1.GitBranches, error) {
	//TODO implement me
	panic("implement me")
}

func (b bitbucketMockService) GetPipeline(repositoryName, username, revision, token string) (*v1.Pipeline, error) {
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

func (b bitbucketMockService) GetDescriptors(repositoryName, username, revision, token, path, env string) ([]unstructured.Unstructured, error) {
	contents, err := b.GetDirectoryContents(repositoryName, username, revision, token, path)
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

func getDescriptorFile() []byte {
	data := `{
				  "apiVersion": "v1",
				  "kind": "ConfigMap",
				  "metadata": {
					"name": "game-demo3"
				  },
				  "data": {
					"player_initial_lives": "3",
					"ui_properties_file_name": "user-interface.properties",
					"game.properties": "enemy.types=aliens,monsters\nplayer.maximum-lives=5\n",
					"user-interface.properties": "color.good=purple\ncolor.bad=yellow\nallow.textmode=true\n"
				  }
			}`
	var descriptions []byte
	descriptions = append(descriptions, []byte(data)...)
	data1 := `{
				  "apiVersion": "v1",
				  "kind": "ConfigMap",
				  "metadata": {
					"name": "game-demo2"
				  },
				  "data": {
					"player_initial_lives": "3",
					"ui_properties_file_name": "user-interface.properties",
					"game.properties": "enemy.types=aliens,monsters\nplayer.maximum-lives=5\n",
					"user-interface.properties": "color.good=purple\ncolor.bad=yellow\nallow.textmode=true\n"
				  }
		}`
	descriptions = append(descriptions, []byte(data1)...)
	return descriptions
}

func getPipelineFile() []byte {
	data := `{
				  "name": "test",
				  "steps": [
					{
					  "name": "build",
					  "type": "BUILD",
					  "trigger": "AUTO",
					  "params": {
						"repository_type": "git",
						"revision": "master",
						"service_account": "test-sa",
						"images": "zeromsi2/test-dev,zeromsi2/test-pro",
						"args_from_configmaps": "tekton/cm-test",
						"args": "key3:value1,key4:value2"
					  },
					  "next": [
						"interstep"
					  ]
					},
					{
					  "name": "interstep",
					  "type": "INTERMEDIARY",
					  "trigger": "AUTO",
					  "params": {
						"revision": "latest",
						"service_account": "test-sa",
						"images": "ubuntu",
						"envs_from_configmaps": "tekton/cm-test",
						"envs_from_secrets": "tekton/cm-test",
						"envs": "key3:value1,key4:value2",
						"command": "echo",
						"command_args": "Hello World"
					  },
					  "next": [
						"deployDev"
					  ]
					},
					{
					  "name": "deployDev",
					  "type": "DEPLOY",
					  "trigger": "AUTO",
					  "params": {
						"agent": "local_agent",
						"name": "ubuntu",
						"namespace": "default",
						"type": "deployment",
						"env": "dev",
						"images": "zeromsi2/test-dev"
					  },
					  "next": [
						"jenkinsJob"
					  ]
					},
					{
					  "name": "jenkinsJob",
					  "type": "JENKINS_JOB",
					  "trigger": "AUTO",
					  "params": {
						"url": "http://localhost:8080",
						"job": "tekton",
						"secret": "jenkins-credentials",
						"params": "id:123,verbosity:high"
					  },
					  "next": null
					}
				  ]
			}`
	var pipelineFile []byte
	pipelineFile = append(pipelineFile, []byte(data)...)

	return pipelineFile
}

func (b bitbucketMockService) GetDirectoryContents(repositoryName, username, revision, token, path string) ([]v1.GitDirectoryContent, error) {
	data := `{
				  "pagelen": 100,
				  "values": [
					{
					  "path": "klovercloud/pipeline/configs",
					  "type": "commit_directory",
					  "links": {
						"self": {
						  "href": "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/configs/"
						},
						"meta": {
						  "href": "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/configs/?format=meta"
						}
					  },
					  "commit": {
						"type": "commit",
						"hash": "920d562bac358658e431b66d749dfc6ff74d35dc",
						"links": {
						  "self": {
							"href": "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/commit/920d562bac358658e431b66d749dfc6ff74d35dc"
						  },
						  "html": {
							"href": "https://bitbucket.org/shabrulislam2451/testapp/commits/920d562bac358658e431b66d749dfc6ff74d35dc"
						  }
						}
					  }
					},
					{
					  "mimetype": null,
					  "links": {
						"self": {
						  "href": "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml"
						},
						"meta": {
						  "href": "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/src/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml?format=meta"
						},
						"history": {
						  "href": "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/filehistory/920d562bac358658e431b66d749dfc6ff74d35dc/klovercloud/pipeline/pipeline.yaml"
						}
					  },
					  "escaped_path": "klovercloud/pipeline/pipeline.yaml",
					  "path": "klovercloud/pipeline/pipeline.yaml",
					  "commit": {
						"type": "commit",
						"hash": "920d562bac358658e431b66d749dfc6ff74d35dc",
						"links": {
						  "self": {
							"href": "https://api.bitbucket.org/2.0/repositories/shabrulislam2451/testapp/commit/920d562bac358658e431b66d749dfc6ff74d35dc"
						  },
						  "html": {
							"href": "https://bitbucket.org/shabrulislam2451/testapp/commits/920d562bac358658e431b66d749dfc6ff74d35dc"
						  }
						}
					  },
					  "attributes": [],
					  "type": "commit_file",
					  "size": 1114
					}
				  ],
				  "page": 1
			}`
	var bitBucketDirectoryContents v1.BitbucketDirectoryContent
	err := json.Unmarshal([]byte(data), &bitBucketDirectoryContents)
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

func (b bitbucketMockService) CreateRepositoryWebhook(username, repositoryName, token string, companyId, appId string) (v1.GitWebhook, error) {
	//TODO implement me
	panic("implement me")
}

func (b bitbucketMockService) DeleteRepositoryWebhookById(username, repositoryName, webhookId, token string) error {
	//TODO implement me
	panic("implement me")
}

// NewBitBucketMockService returns Git type service
func NewBitBucketMockService(observerList []service.Observer, client service.HttpClient) service.Git {
	return &bitbucketMockService{
		observerList: observerList,
		client:       client,
	}
}
