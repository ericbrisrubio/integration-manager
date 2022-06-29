package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

// Repository contains repository info
type Repository struct {
	Id        string                `bson:"id" json:"id"`
	CompanyId string                `bson:"companyId" json:"companyId"`
	Type      enums.REPOSITORY_TYPE `bson:"type" json:"type"`
	Token     string                `bson:"token" json:"token"`
}

// RepositoryDto contains repository info
type RepositoryDto struct {
	Id           string                `bson:"id" json:"id"`
	Type         enums.REPOSITORY_TYPE `bson:"type" json:"type"`
	Token        string                `bson:"token" json:"token"`
	Applications []ApplicationDto      `bson:"applications" json:"applications"`
}

// RepositoriesDto contains repository list
type RepositoriesDto struct {
	Repositories []RepositoryDto `bson:"repositories" json:"repositories"`
}

// Validate validates repository info
//func (repository Repository) Validate() error {
//	if repository.Id == "" {
//		return errors.New("Repository id is required!")
//	}
//	if repository.Token == "" {
//		return errors.New("Repository token is required!")
//	}
//	for _, each := range repository.Applications {
//		err := each.Validate()
//		if err != nil {
//			return err
//		}
//	}
//	if repository.Type == enums.GITHUB || repository.Type == enums.BIT_BUCKET {
//		return nil
//	} else if repository.Type == "" {
//		return errors.New("Repository type is required")
//	}
//	return errors.New("Repository type is invalid!")
//}

// RepositoryUpdateOption contains repository update options
type RepositoryUpdateOption struct {
	Option enums.REPOSITORY_UPDATE_OPTION `json:"option"`
}

// RepositoriesDto contains repository list
//type RepositoriesDto struct {
//	Repositories []Repository `bson:"repositories" json:"repositories"`
//}
