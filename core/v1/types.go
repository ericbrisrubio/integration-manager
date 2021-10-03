package v1

import (
	"github.com/klovercloud-ci/enums"
)

type Subject struct {
	Step, Log             string
	CoreRequestQueryParam map[string]string
	StepType              enums.STEP_TYPE
	EventData             map[string]interface{}
	ProcessLabel          map[string]string
	Pipeline              Pipeline
	App                   struct {
		CompanyId    string
		AppId        string
		RepositoryId string
	}
}

type Repository struct {
	Id           string                `bson:"id" json:"id"`
	Type         enums.REPOSITORY_TYPE `bson:"type" json:"type"`
	Token        string                `bson:"token" json:"token"`
	Applications []Application         `bson:"applications" json:"applications"`
}

func (repository Repository) Validate() error {

	return nil
}

type Application struct {
	MetaData ApplicationMetadata `bson:"_metadata" json:"_metadata"`
	Url      string              `bson:"url" json:"url"`
}

func (application Application) Validate() error {

	return nil
}

type ApplicationMetadata struct {
	Labels map[string]string `bson:"labels" json:"labels"`
	Id     string            `bson:"id" json:"id"`
	Name   string            `bson:"name" json:"name"`
}

func (metadata ApplicationMetadata) Validate() error {

	return nil
}

type CompanyMetadata struct {
	Labels map[string]string `bson:"labels" json:"labels"`
}

func (metadata CompanyMetadata) Validate() error {

	return nil
}

type CompanyUpdateOption struct {
	Option enums.COMPANY_UPDATE_OPTION `json:"option"`
}

type CompanyQueryOption struct {
	Pagination       Pagination
	LoadRepositories bool
	LoadApplications bool
}

type Pagination struct {
	Page  int64
	Limit int64
}

type ProcessInventoryEvent struct {
	ProcessId    string                 `bson:"process_id" json:"process_id"`
	CompanyId    string                 `bson:"company_id" json:"company_id"`
	AppId        string                 `bson:"app_id" json:"app_id"`
	RepositoryId string                 `bson:"repository_id" json:"repository_id"`
	Data         map[string]interface{} `bson:"data" json:"data"`
}
