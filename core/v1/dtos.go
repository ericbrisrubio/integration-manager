package v1

import (
	"github.com/klovercloud-ci/enums"
	_ "github.com/klovercloud-ci/enums"
)

type OnlyCompanyDto struct {
	MetaData CompanyMetadata      `bson:"_metadata" json:"_metadata"`
	Id       string               `bson:"id" json:"id"`
	Name     string               `bson:"name" json:"name"`
	Status   enums.COMPANY_STATUS `bson:"status" json:"status"`
}

type CompanyWiseRepositoriesDto struct {
	MetaData     CompanyMetadata      `bson:"_metadata" json:"_metadata"`
	Id           string               `bson:"id" json:"id"`
	Name         string               `bson:"name" json:"name"`
	Status       enums.COMPANY_STATUS `bson:"status" json:"status"`
	Repositories []struct {
		Type  enums.REPOSITORY_TYPE `bson:"type" json:"type"`
		Token string                `bson:"token" json:"token"`
	} `bson:"repositories" json:"repositories"`
}
type RepositoryWithOutApplication struct {
	Type  enums.REPOSITORY_TYPE `bson:"type" json:"type"`
	Token string                `bson:"token" json:"token"`
}
