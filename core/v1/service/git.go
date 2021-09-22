package service

import v1 "github.com/klovercloud-ci/core/v1"

type Git interface {
	GetPipeline(repogitory_name,username,revision,token string)(*v1.Pipeline,error)
	GetDescriptors(repogitory_name,username,revision,token,path string) ([]interface{},error)
	GetDirectoryContents(repogitory_name,username, revision,token, path string)([]v1.GithubDirectoryContent,error)
}
