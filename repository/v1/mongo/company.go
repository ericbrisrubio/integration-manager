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

// CompanyCollection collection name
var (
	CompanyCollection = "companyCollection"
)

type companyRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (c companyRepository) GetCompanies(option v1.CompanyQueryOption, status v1.StatusQueryOption) ([]v1.Company, int64) {
	var results []v1.Company
	query := bson.M{
		"$and": []bson.M{{"status": string(status.Option)}},
	}
	coll := c.manager.Db.Collection(CompanyCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	result, err := coll.Find(c.manager.Ctx, query, &options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results, int64(len(results))
}

func (c companyRepository) GetByCompanyId(id string) v1.Company {
	query := bson.M{
		"$and": []bson.M{{"id": id}},
	}
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if elemValue.Id != "" {
			return *elemValue
		}
	}
	return v1.Company{}
}

func (c companyRepository) GetByName(name string, status v1.StatusQueryOption) v1.Company {
	query := bson.M{
		"$and": []bson.M{{"name": name}, {"status": string(status.Option)}},
	}
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query, nil)
	if err != nil {
		return v1.Company{}
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		return *elemValue
	}
	return v1.Company{}
}

func GetPagination(size int64, page int64, limit int64) (int64, int64) {
	startIndex := page * limit
	if startIndex >= size {
		return 0, 0
	}
	if size <= startIndex+limit {
		return startIndex, size
	}
	return startIndex, startIndex + limit
}

func (c companyRepository) Store(company v1.Company) error {
	coll := c.manager.Db.Collection(CompanyCollection)
	_, err := coll.InsertOne(c.manager.Ctx, company)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func (c companyRepository) Delete(companyId string) error {
	coll := c.manager.Db.Collection(CompanyCollection)
	filter := bson.M{"id": companyId, "status": bson.M{"$in": enums.ACTIVE}}
	update := bson.M{"$set": bson.M{"status": enums.INACTIVE}}
	_, err := coll.UpdateOne(
		c.manager.Ctx,
		filter,
		update,
	)
	return err
}

// NewCompanyRepository returns CompanyRepository type object
func NewCompanyRepository(timeout int) repository.CompanyRepository {
	return &companyRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
