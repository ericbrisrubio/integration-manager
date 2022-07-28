package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

// SearchData contains repositories and applications info
type SearchData struct {
	Repositories []Repository  `json:"repositories" bson:"repositories"`
	Applications []Application `json:"applications" bson:"applications"`
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

// DashboardData contains company resources count info
type DashboardData struct {
	Repository struct {
		Count int64 `json:"count"`
	} `json:"repository"`
	Application struct {
		Webhook struct {
			Enabled  int64 `json:"enabled"`
			Disabled int64 `json:"disabled"`
		} `json:"webhook"`
	} `json:"application"`
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
