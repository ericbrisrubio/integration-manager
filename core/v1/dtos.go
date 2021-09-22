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

func (dto OnlyCompanyDto) GetCompanyWithoutRepository() Company {
	company := Company{
		MetaData:     dto.MetaData,
		Id:           dto.Id,
		Name:         dto.Name,
		Repositories: nil,
		Status:       dto.Status,
	}
	return company
}

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

func (dto CompanyWiseRepositoriesDto) GetCompanyWithRepository() Company {
	company := Company{
		MetaData:     dto.MetaData,
		Id:           dto.Id,
		Name:         dto.Name,
		Repositories: nil,
		Status:       dto.Status,
	}
	for _, each := range dto.Repositories {
		company.Repositories = append(company.Repositories, Repository{
			Id:           each.Id,
			Type:         each.Type,
			Token:        each.Token,
			Applications: nil,
		})
	}
	return company
}

type RepositoryWithOutApplication struct {
	Id    string                `bson:"_Id" json:"_Id"`
	Type  enums.REPOSITORY_TYPE `bson:"type" json:"type"`
	Token string                `bson:"token" json:"token"`
}

func (dto RepositoryWithOutApplication) GetRepositoryWithoutApplication() Repository {
	repository := Repository{
		Id:           dto.Id,
		Type:         dto.Type,
		Token:        dto.Token,
		Applications: nil,
	}
	return repository
}
