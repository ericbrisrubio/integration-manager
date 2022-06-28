package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// RepositoryCollection collection name
var (
	RepositoryCollection = "repositoryCollection"
)

type repositoryRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (r repositoryRepository) GetByCompanyId(companyId string) []v1.Repository {
	var results []v1.Repository
	query := bson.M{
		"$and": []bson.M{{"companyId": companyId}},
	}
	coll := r.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(r.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Repository)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (r repositoryRepository) Store(repositories []v1.Repository) error {
	coll := r.manager.Db.Collection(RepositoryCollection)
	if len(repositories) > 0 {
		var payload []interface{}
		payload = append(payload, repositories)
		_, err := coll.InsertMany(r.manager.Ctx, payload)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewRepositoryRepository returns RepositoryRepository type object
func NewRepositoryRepository(timeout int) repository.RepositoryRepository {
	return &repositoryRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
