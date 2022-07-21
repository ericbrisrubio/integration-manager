package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// SearchCollection collection name
var (
	ApplicationMetadataCollection = "applicationMetadataCollection"
)

type applicationMetadataRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (a applicationMetadataRepository) Store(applicationMetadataCollection v1.ApplicationMetadataCollection) error {
	coll := a.manager.Db.Collection(ApplicationMetadataCollection)
	_, err := coll.InsertOne(a.manager.Ctx, applicationMetadataCollection)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func (a applicationMetadataRepository) SearchAppsByCompanyIdAndName(companyId, name string) []v1.ApplicationMetadataCollection {
	var results []v1.ApplicationMetadataCollection
	query := bson.M{
		"$and": []bson.M{
			{"_metadata.labels.companyId": companyId},
			{"_metadata.name": primitive.Regex{Pattern: name}},
		},
	}
	coll := a.manager.Db.Collection(ApplicationMetadataCollection)
	result, err := coll.Find(a.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.ApplicationMetadataCollection)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (a applicationMetadataRepository) GetById(id, companyId string) v1.ApplicationMetadataCollection {
	var data v1.ApplicationMetadataCollection
	query := bson.M{
		"$and": []bson.M{
			{"_metadata.labels.companyId": companyId},
			{"_metadata.id": id},
		},
	}
	coll := a.manager.Db.Collection(ApplicationMetadataCollection)
	result, err := coll.Find(a.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		err := result.Decode(data)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
	}
	return data
}

func (a applicationMetadataRepository) Update(companyId string, data v1.ApplicationMetadataCollection) error {
	filter := bson.M{
		"$and": []bson.M{
			{"_metadata.id": data.MetaData.Id},
			{"_metadata.labels.companyId": companyId},
		},
	}
	update := bson.M{
		"$set": data,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := a.manager.Db.Collection(ApplicationMetadataCollection)
	err := coll.FindOneAndUpdate(a.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (a applicationMetadataRepository) Delete(id, companyId string) error {
	filter := bson.M{
		"$and": []bson.M{
			{"_metadata.id": id},
			{"_metadata.labels.companyId": companyId},
		},
	}
	coll := a.manager.Db.Collection(ApplicationMetadataCollection)
	_, err := coll.DeleteOne(a.manager.Ctx, filter, nil)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return nil
}

// NewApplicationMetadataRepository returns ApplicationMetadataRepository type object
func NewApplicationMetadataRepository(timeout int) repository.ApplicationMetadataRepository {
	return &applicationMetadataRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
