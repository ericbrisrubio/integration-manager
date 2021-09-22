package v1

import "github.com/klovercloud-ci/enums"

type Resource struct {
	Type               enums.PIPELINE_RESOURCE_TYPE `json:"type" yaml:"type"`
	Url                string                       `json:"url"  yaml:"url"`
	Revision           string                       `json:"revision"  yaml:"revision"`
	DeploymentResource *DeploymentResource          `json:"deployment_resource"  yaml:"deployment_resource"`
}

type DeploymentResource struct {
	MountPath   *string                      `json:"mount_path" yaml:"mount_path"`
	Descriptors *[]interface{}               `json:"descriptors" yaml:"descriptors"`
	ProcessId   string                       `json:"process_id" yaml:"process_id"`
	Agent       string                       `json:"agent" yaml:"agent"`
	Type        enums.PIPELINE_RESOURCE_TYPE `json:"type"`
	Step        string                       `json:"step" yaml:"step"`
	Name        string                       `json:"name" yaml:"name"`
	Namespace   string                       `json:"namespace" yaml:"namespace"`
	Replica     int32                        `json:"replica" yaml:"replica"`
	Images      []struct {
		ImageIndex int    `json:"image_index" yaml:"image_index"`
		Image      string `json:"image" yaml:"image"`
	} `json:"images" yaml:"images"`
}
type Variable struct {
	Secrets []struct {
		Name      string `json:"name" yaml:"name"`
		Namespace string `json:"namespace" yaml:"namespace"`
	} `json:"secrets" yaml:"secrets"`
	ConfigMaps []struct {
		Name      string `json:"name" yaml:"name"`
		Namespace string `json:"namespace" yaml:"namespace"`
	} `json:"configMaps"  yaml:"configMaps"`
	Data map[string]string `json:"data"  yaml:"data"`
}

type Subject struct {
	Step,Log string
	CoreRequestQueryParam map[string]string
	StepType enums.STEP_TYPE
	EventData map[string]interface{}
	ProcessLabel map[string]string
	Pipeline     Pipeline
}
type Repository struct {
	Id           string                `bson:"_Id" json:"_Id"`
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
