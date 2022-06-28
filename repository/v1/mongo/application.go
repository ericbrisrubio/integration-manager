package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// ApplicationCollection collection name
var (
	ApplicationCollection = "applicationCollection"
)

type applicationRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (a applicationRepository) GetByCompanyIdAndRepoId(companyId, repoId string) []v1.Application {
	var results []v1.Application
	query := bson.M{
		"$and": []bson.M{{"companyId": companyId}, {"repositoryId": repoId}},
	}
	coll := a.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(a.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Application)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (a applicationRepository) Store(application v1.Application) error {
	coll := a.manager.Db.Collection(ApplicationCollection)
	_, err := coll.InsertOne(a.manager.Ctx, application)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func (a applicationRepository) StoreAll(applications []v1.Application) error {
	coll := a.manager.Db.Collection(ApplicationCollection)
	if len(applications) > 0 {
		var payload []interface{}
		payload = append(payload, applications)
		_, err := coll.InsertMany(a.manager.Ctx, payload)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewApplicationRepository returns ApplicationRepository type object
func NewApplicationRepository(timeout int) repository.ApplicationRepository {
	return &applicationRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
