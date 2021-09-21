package logic

import (
	"encoding/json"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"log"
)

type githubService struct {
	companyService service.Company
	observerList []service.Observer
	client service.HttpClient
}

func (githubService githubService) GetPipeline(repogitory_name,username, revision,token string) *v1.Pipeline {
	url:=enums.GITHUB_RAW_CONTENT_BASE_URL+username+"/"+repogitory_name+"/"+revision+"/"+enums.PIPELINE_FILE_NAME
	header :=make(map[string]string)
	header["Authorization"]="token "+token
	header["Accept"]="application/vnd.github.v3.raw"
	header["X-Requested-With"]="Curl"
	err,data:=githubService.client.Get(url,header)
	if err!=nil{
		// send to observer
		return nil
	}
	pipeline:=v1.Pipeline{}
	err = json.Unmarshal(data, &pipeline)
	if err != nil {
		log.Println(err.Error())
		// send to observer
		return nil
	}
	return &pipeline
}

func (githubService githubService) GetDescriptors(repogitory_name,username, revision,token, path string) []interface{} {
	panic("implement me")
}

func (githubService githubService)notifyAll(listener v1.Subject){
	for _, observer := range githubService.observerList {
		go observer.Listen(listener)
	}
}
func NewGithubService(companyService service.Company,observerList []service.Observer,client service.HttpClient) service.Git {
	return &githubService{
		companyService: companyService,
		observerList:  observerList,
		client:         client,
	}
}
