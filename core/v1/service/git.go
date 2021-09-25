package service

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Git interface {
	GetPipeline(repogitory_name,username,revision,token string)(*v1.Pipeline,error)
	GetDescriptors(repogitory_name,username,revision,token,path string) ([]unstructured.Unstructured,error)
	GetDirectoryContents(repogitory_name,username, revision,token, path string)([]v1.GithubDirectoryContent,error)
}
