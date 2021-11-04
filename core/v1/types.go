package v1

import (
	"errors"
	"github.com/klovercloud-ci/enums"
	"reflect"
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
	if repository.Id == "" {
		return errors.New("Repository id is required!")
	}
	if repository.Token == "" {
		return errors.New("Repository token is required!")
	}
	for _, each := range repository.Applications {
		err := each.Validate()
		if err != nil {
			return err
		}
	}
	if repository.Type == enums.GITHUB || repository.Type == enums.BIT_BUCKET {
		return nil
	} else if repository.Type == "" {
		return errors.New("Repository type is required")
	} else {
		return errors.New("Repository type is invalid!")
	}
}

type Application struct {
	MetaData ApplicationMetadata  `bson:"_metadata" json:"_metadata"`
	Url      string               `bson:"url" json:"url"`
	Webhook  GithubWebhook        `bson:"webhook" json:"webhook"`
	Status   enums.COMPANY_STATUS `bson:"status" json:"status"`
}

func (application Application) Validate() error {
	if application.Url == "" {
		return errors.New("Application url is required!")
	}
	err := application.MetaData.Validate()
	if err != nil {
		return err
	}
	return nil
}

type ApplicationMetadata struct {
	Labels           map[string]string `bson:"labels" json:"labels"`
	Id               string            `bson:"id" json:"id"`
	Name             string            `bson:"name" json:"name"`
	Branches         []string          `bson:"branches" json:"branches"`
	IsWebhookEnabled bool              `bson:"is_webhook_enabled" json:"is_webhook_enabled"`
}

func (metadata ApplicationMetadata) Validate() error {
	keys := reflect.ValueOf(metadata.Labels).MapKeys()
	for i := 0; i < len(keys); i++ {
		if metadata.Labels[keys[i].String()] == "" {
			return errors.New("Application metadata label is missing!")
		}
	}
	if metadata.Id == "" {
		return errors.New("Application metadata id is required!")
	}
	if metadata.Name == "" {
		return errors.New("Application metadata name is required!")
	}
	return nil
}

type CompanyMetadata struct {
	Labels                    map[string]string `bson:"labels" json:"labels" yaml:"labels"`
	NumberOfConcurrentProcess int64             `bson:"number_of_concurrent_process" json:"number_of_concurrent_process" yaml:"number_of_concurrent_process"`
	TotalProcessPerDay        int64             `bson:"total_process_per_day" json:"total_process_per_day" yaml:"total_process_per_day"`
}

func (metadata CompanyMetadata) Validate() error {
	keys := reflect.ValueOf(metadata.Labels).MapKeys()
	for i := 0; i < len(keys); i++ {
		if metadata.Labels[keys[i].String()] == "" {
			return errors.New("Company metadata label is missing!")
		}
	}
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
type PipelineMetadata struct {
	CompanyId       string          `json:"company_id" yaml:"company_id"`
	CompanyMetadata CompanyMetadata `json:"company_metadata" yaml:"company_metadata"`
}
