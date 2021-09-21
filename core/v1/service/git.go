package service

import v1 "github.com/klovercloud-ci/core/v1"

type Git interface {
	GetPipeline(repogitory_name,username,revision,token string)*v1.Pipeline
	GetDescriptors(repogitory_name,username,revision,token,path string) []interface{}
}
