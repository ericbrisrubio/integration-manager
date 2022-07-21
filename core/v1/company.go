package v1

import (
	"errors"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"reflect"
)

// CompanyDto contains company data
type CompanyDto struct {
	MetaData     CompanyMetadata      `bson:"_metadata" json:"_metadata"`
	Id           string               `bson:"id" json:"id"`
	Name         string               `bson:"name" json:"name"`
	Repositories []RepositoryDto      `bson:"repositories" json:"repositories"`
	Status       enums.COMPANY_STATUS `bson:"status" json:"status"`
}

// Company contains company data
type Company struct {
	MetaData CompanyMetadata      `bson:"_metadata" json:"_metadata"`
	Id       string               `bson:"id" json:"id"`
	Name     string               `bson:"name" json:"name"`
	Status   enums.COMPANY_STATUS `bson:"status" json:"status"`
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

// Validate validates company data
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
	//for _, each := range dto.Repositories {
	//	err := each.Validate()
	//	if err != nil {
	//		return err
	//	}
	//}
	if dto.Status == enums.ACTIVE || dto.Status == enums.INACTIVE {
		return nil
	} else if dto.Status == "" {
		return errors.New("Company status is required!")
	}
	return errors.New("Company status invalid!")
}

// GetCompanyWithoutRepository returns company without repositories
//func (dto Company) GetCompanyWithoutRepository() Company {
//	company := Company{
//		MetaData:     dto.MetaData,
//		Id:           dto.Id,
//		Name:         dto.Name,
//		Repositories: nil,
//		Status:       dto.Status,
//	}
//	return company
//}

// GetCompanyWithoutApplications returns company with repositories
//func (dto Company) GetCompanyWithoutApplications(option CompanyQueryOption) Company {
//	company := Company{
//		MetaData: dto.MetaData,
//		Id:       dto.Id,
//		Name:     dto.Name,
//		Status:   dto.Status,
//	}
//	for _, each := range dto.Repositories {
//		if option.LoadToken {
//			r := Repository{
//				Id:           each.Id,
//				Type:         each.Type,
//				Token:        each.Token,
//				Applications: nil,
//			}
//			company.Repositories = append(company.Repositories, r)
//		} else {
//			company.Repositories = append(company.Repositories, Repository{
//				Id:           each.Id,
//				Type:         each.Type,
//				Token:        "",
//				Applications: nil,
//			})
//		}
//	}
//	return company
//}
