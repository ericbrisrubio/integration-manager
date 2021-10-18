package v1

import "github.com/klovercloud-ci/enums"

type Company struct {
	MetaData     CompanyMetadata      `bson:"_metadata" json:"_metadata"`
	Id           string               `bson:"id" json:"id"`
	Name         string               `bson:"name" json:"name"`
	Repositories []Repository         `bson:"repositories" json:"repositories"`
	Status       enums.COMPANY_STATUS `bson:"status" json:"status"`
}

func (dto Company) Validate() error {
	return nil
}
