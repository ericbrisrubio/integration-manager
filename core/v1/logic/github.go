package logic

import (
	"encoding/json"
	"fmt"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"log"
	"strings"
)

type githubService struct {
	companyService service.Company
	observerList []service.Observer
	client service.HttpClient
}

func (githubService githubService) GetPipeline(repogitory_name,username, revision,token string) (*v1.Pipeline,error) {
	url:=enums.GITHUB_RAW_CONTENT_BASE_URL+username+"/"+repogitory_name+"/"+revision+"/"+enums.PIPELINE_FILE_NAME
	header :=make(map[string]string)
	header["Authorization"]="token "+token
	header["Accept"]="application/vnd.github.v3.raw"
	header["X-Requested-With"]="Curl"
	err,data:=githubService.client.Get(url,header)
	if err!=nil{
		// send to observer
		return nil,err
	}
	pipeline:=v1.Pipeline{}
	err = json.Unmarshal(data, &pipeline)
	if err != nil {
		log.Println(err.Error())
		// send to observer
		return nil,err
	}
	return &pipeline,nil
}

func (githubService githubService) GetDescriptors(repogitory_name,username, revision,token, path string) ([]interface{},error) {
	contents,err:=githubService.GetDirectoryContents(repogitory_name,username,revision,token,path)
	if err!=nil{
		return nil, err
	}
	var files []interface{}

	for _,each:=range contents{
		if each.Type!="file" {
			continue
		}
		if !strings.HasSuffix(each.Name,".yaml") && !strings.HasSuffix(each.Name,".yml"){
			continue
		}
		url:=fmt.Sprint(each.DownloadURL)
		header :=make(map[string]string)
		header["Authorization"]="token "+token
		header["Accept"]="application/vnd.github.v3.raw"
		header["X-Requested-With"]="Curl"
		err,data:=githubService.client.Get(url,header)
		if err!=nil{
			// send to observer
			return nil,err
		}
		files= append(files, data)
	}
	return files,nil
}

func(githubService githubService) GetDirectoryContents(repogitory_name,username, revision,token, path string)([]v1.GithubDirectoryContent,error){
	if strings.HasPrefix(path,"/"){
		path=strings.TrimPrefix(path,"/")
	}
	url:=enums.GITHUB_API_BASE_URL+"repos/"+username+"/"+repogitory_name+"/contents/"+path+"?ref="+revision
	header :=make(map[string]string)
	header["Authorization"]="token "+token
	header["Accept"]="application/vnd.github.v3+json"
	err,data:=githubService.client.Get(url,header)
	if err!=nil{
		// send to observer
		return nil,err
	}
	contents:=[]v1.GithubDirectoryContent{}
	err = json.Unmarshal(data, &contents)
	if err != nil {
		log.Println(err.Error())
		// send to observer
		return nil,err
	}

	return contents,nil
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
