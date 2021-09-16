package v1

import "github.com/klovercloud-ci/enums"

type Repository struct {
	Type enums.REPOSITORY_TYPE `bson:"type" json:"type"`
	Token string  `bson:"token" json:"token"`
	Applications [] Application `bson:"applications" json:"applications"`
}

func (repository Repository) Validate() error{

	return nil
}

type Application struct {
	MetaData ApplicationMetadata `bson:"_metadata" json:"_metadata"`
	Url string  `bson:"url" json:"url"`
}

func (application Application) Validate() error{

	return nil
}
type ApplicationMetadata struct {
	Labels map[string]string  `bson:"labels" json:"labels"`
}

func (metadata ApplicationMetadata) Validate() error{

	return nil
}
type CompanyMetadata struct {
	Labels map[string]string  `bson:"labels" json:"labels"`
}


func (metadata CompanyMetadata) Validate() error{

	return nil
}

type CompanyUpdateOption struct {
	Option enums.COMPANY_UPDATE_OPTION   `json:"option"`
}

type CompanyQueryOption struct {
	Pagination Pagination
	LoadRepositories bool
	LoadApplications bool
}

type Pagination struct {
	Page  int64
	Limit int64
}