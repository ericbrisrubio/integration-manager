package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (a applicationRepository) GetById(companyId string, repoId string, applicationId string) v1.Application {
	query := bson.M{
		"$and": []bson.M{
			{"repositoryId": repoId},
			{"companyId": companyId},
			{"_metadata.id": applicationId},
		},
	}
	temp := new(v1.Application)
	coll := a.manager.Db.Collection(ApplicationCollection)
	result := coll.FindOne(a.manager.Ctx, query)
	err := result.Decode(&temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return *temp
}

func (a applicationRepository) GetAll(companyId string, option v1.CompanyQueryOption) ([]v1.Application, int64) {
	var results []v1.Application
	query := bson.M{
		"$and": []bson.M{
			{"companyId": companyId},
		},
	}
	coll := a.manager.Db.Collection(ApplicationCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOption := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
	}
	result, err := coll.Find(a.manager.Ctx, query, &findOption)
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
	count, err := coll.CountDocuments(a.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

func (a applicationRepository) GetByCompanyIdAndRepoId(companyId, repoId string, pagination bool, option v1.CompanyQueryOption, statusQuery bool, status v1.StatusQueryOption) ([]v1.Application, int64) {
	var results []v1.Application
	var query bson.M
	if statusQuery {
		query = bson.M{
			"$and": []bson.M{
				{"companyId": companyId},
				{"repositoryId": repoId},
				{"status": status.Option},
			},
		}
	} else {
		query = bson.M{
			"$and": []bson.M{
				{"companyId": companyId},
				{"repositoryId": repoId},
			},
		}
	}
	coll := a.manager.Db.Collection(ApplicationCollection)
	var findOption options.FindOptions
	if pagination {
		skip := option.Pagination.Page * option.Pagination.Limit
		findOption.Limit = &option.Pagination.Limit
		findOption.Skip = &skip
	}
	result, err := coll.Find(a.manager.Ctx, query, &findOption)
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
	count, err := coll.CountDocuments(a.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

func (a applicationRepository) GetByCompanyIdAndRepositoryIdAndUrl(companyId, repositoryId, applicationUrl string) v1.Application {
	query := bson.M{
		"$and": []bson.M{
			{"repositoryId": repositoryId},
			{"companyId": companyId},
			{"url": applicationUrl},
		},
	}
	temp := new(v1.Application)
	coll := a.manager.Db.Collection(ApplicationCollection)
	result := coll.FindOne(a.manager.Ctx, query)
	err := result.Decode(&temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return *temp
}

func (a applicationRepository) GetByCompanyIdAndUrl(companyId, applicationUrl string) v1.Application {
	query := bson.M{
		"$and": []bson.M{
			{"companyId": companyId},
			{"url": applicationUrl},
		},
	}
	temp := new(v1.Application)
	coll := a.manager.Db.Collection(ApplicationCollection)
	result := coll.FindOne(a.manager.Ctx, query)
	err := result.Decode(&temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return *temp
}

func (a applicationRepository) GetApplicationsByCompanyIdAndRepositoryType(companyId string, _type enums.REPOSITORY_TYPE, pagination bool, option v1.CompanyQueryOption, statusQuery bool, status v1.StatusQueryOption) ([]v1.Application, int64) {
	var results []v1.Application
	var query bson.M
	if statusQuery {
		query = bson.M{
			"$and": []bson.M{
				{"companyId": companyId},
				{"repository_type": _type},
				{"status": status.Option},
			},
		}
	} else {
		query = bson.M{
			"$and": []bson.M{
				{"companyId": companyId},
				{"repository_type": _type},
			},
		}
	}
	coll := a.manager.Db.Collection(ApplicationCollection)
	var findOption options.FindOptions
	if pagination {
		skip := option.Pagination.Page * option.Pagination.Limit
		findOption.Limit = &option.Pagination.Limit
		findOption.Skip = &skip
	}
	result, err := coll.Find(a.manager.Ctx, query, &findOption)
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
	count, err := coll.CountDocuments(a.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
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

func (a applicationRepository) Update(companyId, repositoryId string, application v1.Application) error {
	filter := bson.M{
		"$and": []bson.M{
			{"_metadata.id": application.MetaData.Id},
			{"repositoryId": repositoryId},
			{"companyId": companyId},
		},
	}
	update := bson.M{
		"$set": application,
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := a.manager.Db.Collection(ApplicationCollection)
	err := coll.FindOneAndUpdate(a.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (a applicationRepository) SoftDeleteApplication(application v1.Application) error {
	filter := bson.M{
		"$and": []bson.M{
			{"_metadata.id": application.MetaData.Id},
		},
	}
	update := bson.M{
		"$set": application,
	}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	coll := a.manager.Db.Collection(ApplicationCollection)
	err := coll.FindOneAndUpdate(a.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (a applicationRepository) DeleteApplication(companyId, repositoryId, applicationId string) error {
	filter := bson.M{
		"$and": []bson.M{
			{"_metadata.id": applicationId},
			{"repositoryId": repositoryId},
			{"companyId": companyId},
		},
	}
	coll := a.manager.Db.Collection(ApplicationCollection)
	_, err := coll.DeleteOne(a.manager.Ctx, filter, nil)
	if err != nil {
		log.Println("[ERROR]", err)
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
