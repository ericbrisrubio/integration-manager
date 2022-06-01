package v1

import (
	"errors"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"reflect"
)

// SearchData contains repositories and applications info
type SearchData struct {
	Repositories []Repository                    `json:"repositories" bson:"repositories"`
	Applications []ApplicationMetadataCollection `json:"applications" bson:"applications"`
}

// Subject observers listen event with an object of this struct
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
		Branch       string
	}
}

// Repository contains repository info
type Repository struct {
	Id           string                `bson:"id" json:"id"`
	Type         enums.REPOSITORY_TYPE `bson:"type" json:"type"`
	Token        string                `bson:"token" json:"token"`
	Applications []Application         `bson:"applications" json:"applications"`
}

// RepositoryDto contains repository info
type RepositoryDto struct {
	Id           string                `bson:"id" json:"id"`
	Type         enums.REPOSITORY_TYPE `bson:"type" json:"type"`
	Applications []Application         `bson:"applications" json:"applications"`
}

// Validate validates repository info
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
	}
	return errors.New("Repository type is invalid!")

}

// Application contains application info
type Application struct {
	MetaData ApplicationMetadata  `bson:"_metadata" json:"_metadata"`
	Url      string               `bson:"url" json:"url"`
	Webhook  GitWebhook           `bson:"webhook" json:"webhook"`
	Status   enums.COMPANY_STATUS `bson:"status" json:"status"`
}

// Validate validates application info
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

// ApplicationMetadata contains application metadata info
type ApplicationMetadata struct {
	Labels           map[string]string `bson:"labels" json:"labels"`
	Id               string            `bson:"id" json:"id"`
	Name             string            `bson:"name" json:"name"`
	IsWebhookEnabled bool              `bson:"is_webhook_enabled" json:"is_webhook_enabled"`
}

// ApplicationMetadataCollection contains application metadata collection info
type ApplicationMetadataCollection struct {
	MetaData ApplicationMetadata  `bson:"_metadata" json:"_metadata"`
	Status   enums.COMPANY_STATUS `bson:"status" json:"status"`
}

// ApplicationMetadataCollectionsDto contains application metadata collections list
type ApplicationMetadataCollections struct {
	ApplicationMetadataCollection []ApplicationMetadataCollection `bson:"application_metadata_collection" json:"application_metadata_collection"`
}

// Validate validates application metadata
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

// CompanyMetadata contains company metadata info
type CompanyMetadata struct {
	Labels                    map[string]string `bson:"labels" json:"labels" yaml:"labels"`
	NumberOfConcurrentProcess int64             `bson:"number_of_concurrent_process" json:"number_of_concurrent_process" yaml:"number_of_concurrent_process"`
	TotalProcessPerDay        int64             `bson:"total_process_per_day" json:"total_process_per_day" yaml:"total_process_per_day"`
}

// Validate validates company metadata
func (metadata CompanyMetadata) Validate() error {
	keys := reflect.ValueOf(metadata.Labels).MapKeys()
	for i := 0; i < len(keys); i++ {
		if metadata.Labels[keys[i].String()] == "" {
			return errors.New("Company metadata label is missing!")
		}
	}
	return nil
}

// ApplicationUpdateOption contains applications update options
type ApplicationUpdateOption struct {
	Option enums.APPLICATION_UPDATE_OPTION `json:"option"`
}

// RepositoryUpdateOption contains repository update options
type RepositoryUpdateOption struct {
	Option enums.REPOSITORY_UPDATE_OPTION `json:"option"`
}

// StatusQueryOption contains company update options
type StatusQueryOption struct {
	Option enums.COMPANY_STATUS `json:"option"`
}

// CompanyQueryOption contains company query options
type CompanyQueryOption struct {
	Pagination       Pagination
	LoadRepositories bool
	LoadApplications bool
	LoadToken        bool
}

// Pagination contains pagination options
type Pagination struct {
	Page  int64
	Limit int64
}

// Process contains process inventory event options
type Process struct {
	ProcessId    string                 `bson:"process_id" json:"process_id"`
	CompanyId    string                 `bson:"company_id" json:"company_id"`
	AppId        string                 `bson:"app_id" json:"app_id"`
	RepositoryId string                 `bson:"repository_id" json:"repository_id"`
	CommitId     string                 `bson:"commit_id" json:"commit_id"`
	Data         map[string]interface{} `bson:"data" json:"data"`
	Branch       string                 `bson:"branch" json:"branch"`
}

// PipelineMetadata contains pipeline metadata event options
type PipelineMetadata struct {
	CompanyId       string          `json:"company_id" yaml:"company_id"`
	CompanyMetadata CompanyMetadata `json:"company_metadata" yaml:"company_metadata"`
	CommitId        string          `json:"commit_id" yaml:"commit_id"`
}

// OnlyCompanyDto contains only company info
type OnlyCompanyDto struct {
	MetaData CompanyMetadata      `bson:"_metadata" json:"_metadata"`
	Id       string               `bson:"id" json:"id"`
	Name     string               `bson:"name" json:"name"`
	Status   enums.COMPANY_STATUS `bson:"status" json:"status"`
}

// GetCompanyWithoutRepository returns company without repositories
func (dto Company) GetCompanyWithoutRepository() Company {
	company := Company{
		MetaData:     dto.MetaData,
		Id:           dto.Id,
		Name:         dto.Name,
		Repositories: nil,
		Status:       dto.Status,
	}
	return company
}

// CompanyWiseRepositoriesDto contains company wise repositories
type CompanyWiseRepositoriesDto struct {
	MetaData     CompanyMetadata      `bson:"_metadata" json:"_metadata"`
	Id           string               `bson:"id" json:"id"`
	Name         string               `bson:"name" json:"name"`
	Status       enums.COMPANY_STATUS `bson:"status" json:"status"`
	Repositories []struct {
		Id    string                `bson:"_Id" json:"_Id"`
		Type  enums.REPOSITORY_TYPE `bson:"type" json:"type"`
		Token string                `bson:"token" json:"token"`
	} `bson:"repositories" json:"repositories"`
}

// ApplicationsDto contains application list
type ApplicationsDto struct {
	Applications []Application `bson:"applications" json:"applications"`
}

// RepositoriesDto contains repository list
type RepositoriesDto struct {
	Repositories []Repository `bson:"repositories" json:"repositories"`
}

// GetCompanyWithoutApplications returns company with repositories
func (dto Company) GetCompanyWithoutApplications(option CompanyQueryOption) Company {
	company := Company{
		MetaData: dto.MetaData,
		Id:       dto.Id,
		Name:     dto.Name,
		Status:   dto.Status,
	}
	for _, each := range dto.Repositories {
		if option.LoadToken {
			r := Repository{
				Id:           each.Id,
				Type:         each.Type,
				Token:        each.Token,
				Applications: nil,
			}
			company.Repositories = append(company.Repositories, r)
		} else {
			company.Repositories = append(company.Repositories, Repository{
				Id:           each.Id,
				Type:         each.Type,
				Token:        "",
				Applications: nil,
			})
		}
	}
	return company
}

// OnlyRepository contains only repository info
type OnlyRepository struct {
	Id    string                `bson:"_Id" json:"_Id"`
	Type  enums.REPOSITORY_TYPE `bson:"type" json:"type"`
	Token string                `bson:"token" json:"token"`
}

// GetRepositoryWithoutApplication returns repository without applications
func (dto OnlyRepository) GetRepositoryWithoutApplication() Repository {
	repository := Repository{
		Id:           dto.Id,
		Type:         dto.Type,
		Token:        dto.Token,
		Applications: nil,
	}
	return repository
}
