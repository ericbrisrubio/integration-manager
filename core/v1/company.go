package v1

import (
	"errors"
	"github.com/klovercloud-ci/enums"
)

type Company struct {
	MetaData     CompanyMetadata      `bson:"_metadata" json:"_metadata"`
	Id           string               `bson:"id" json:"id"`
	Name         string               `bson:"name" json:"name"`
	Repositories []Repository         `bson:"repositories" json:"repositories"`
	Status       enums.COMPANY_STATUS `bson:"status" json:"status"`
}

func (dto Company) Validate() error {
	err := dto.MetaData.Validate()
	if err != nil {
		return err
	}
	if dto.Id == "" {
		return errors.New("Company id is required!")
	}
	if dto.Name == "" {
		return errors.New("Company name is required!")
	}
	for _, each := range dto.Repositories {
		err := each.Validate()
		if err != nil {
			return err
		}
	}
	if dto.Status == enums.ACTIVE || dto.Status == enums.INACTIVE {
		return nil
	} else if dto.Status == "" {
		return errors.New("Company status is required!")
	} else {
		return errors.New("Company status invalid!")
	}
}
