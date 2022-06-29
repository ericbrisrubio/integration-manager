package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r repositoryRepository) GetById(companyId, repositoryId string) v1.Repository {
	query := bson.M{
		"$and": []bson.M{
			{"id": repositoryId},
			{"companyId": companyId},
		},
	}
	temp := new(v1.Repository)
	coll := r.manager.Db.Collection(RepositoryCollection)
	result := coll.FindOne(r.manager.Ctx, query)
	err := result.Decode(&temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return *temp
}

func (r repositoryRepository) GetByCompanyId(companyId string, pagination bool, option v1.CompanyQueryOption) ([]v1.Repository, int64) {
	var results []v1.Repository
	query := bson.M{
		"$and": []bson.M{{"companyId": companyId}},
	}
	coll := r.manager.Db.Collection(RepositoryCollection)
	var findOption options.FindOptions
	if pagination {
		skip := option.Pagination.Page * option.Pagination.Limit
		findOption.Limit = &option.Pagination.Limit
		findOption.Skip = &skip
	}
	result, err := coll.Find(r.manager.Ctx, query, &findOption)
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
	count, err := coll.CountDocuments(r.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
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

func (r repositoryRepository) DeleteRepository(companyId, repositoryId string) error {
	filter := bson.M{
		"$and": []bson.M{
			{"id": repositoryId},
			{"companyId": companyId},
		},
	}
	coll := r.manager.Db.Collection(RepositoryCollection)
	_, err := coll.DeleteOne(r.manager.Ctx, filter, nil)
	if err != nil {
		log.Println("[ERROR]", err)
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
