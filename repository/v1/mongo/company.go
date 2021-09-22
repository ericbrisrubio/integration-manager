package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/enums"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	CompanyCollection = "CompanyCollection"
)

type companyRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (c companyRepository) GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository {
	var results v1.Repository
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": id}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new([]v1.Repository)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, each := range *elemValue {
			for _, eachApp := range each.Applications {
				if url == eachApp.Url {
					results = each
				}
			}
		}
	}
	return results
}

func (c companyRepository) GetCompanyByApplicationUrl(url string) v1.Company {
	var results v1.Company
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, nil, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new([]v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, each := range *elemValue {
			for _, eachRepo := range each.Repositories {
				for _, eachApp := range eachRepo.Applications {
					if url == eachApp.Url {
						results = each
					}
				}
			}
		}
	}
	return results
}

func (c companyRepository) GetCompanies(option v1.CompanyQueryOption) ([]v1.Company, int64) {
	var results []v1.Company
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, nil, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new([]v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if option.LoadRepositories {
			if option.LoadApplications {
				results = *elemValue
			} else {
				for i, each := range *elemValue {
					results[i].Id = each.Id
					results[i].MetaData = each.MetaData
					results[i].Name = each.Name
					results[i].Status = each.Status
					for j, eachRepo := range each.Repositories {
						results[i].Repositories[j].Type = eachRepo.Type
						results[i].Repositories[j].Token = eachRepo.Token
						results[i].Repositories[j].Applications = nil
					}
				}
			}
		} else {
			for i, each := range *elemValue {
				results[i].Id = each.Id
				results[i].MetaData = each.MetaData
				results[i].Name = each.Name
				results[i].Status = each.Status
				results[i].Repositories = nil
			}

		}

	}
	return results, int64(len(results))
}

func (c companyRepository) GetByCompanyId(id string, option v1.CompanyQueryOption) v1.Company {
	var results v1.Company
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": id}}
	query["$and"] = and
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
		if option.LoadRepositories {
			if option.LoadApplications {
				elemValue := new(v1.Company)
				err := result.Decode(elemValue)
				if err != nil {
					log.Println("[ERROR]", err)
					break
				}
				results = *elemValue
			} else {
				elemValue := new(v1.CompanyWiseRepositoriesDto)
				err := result.Decode(elemValue)
				if err != nil {
					log.Println("[ERROR]", err)
					break
				}

				results.Id = elemValue.Id
				results.MetaData = elemValue.MetaData
				results.Name = elemValue.Name
				results.Status = elemValue.Status
				for i, each := range elemValue.Repositories {
					results.Repositories[i].Type = each.Type
					results.Repositories[i].Token = each.Token
					results.Repositories[i].Applications = nil
				}
			}
		} else {
			elemValue := new(v1.OnlyCompanyDto)
			err := result.Decode(elemValue)
			if err != nil {
				log.Println("[ERROR]", err)
				break
			}

			results.Id = elemValue.Id
			results.MetaData = elemValue.MetaData
			results.Name = elemValue.Name
			results.Status = elemValue.Status
			results.Repositories = nil
		}

	}
	return results
}

func (c companyRepository) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) []v1.Repository {
	var results []v1.Repository
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": id}}
	query["$and"] = and
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
		if option.LoadRepositories {
			if option.LoadApplications {
				elemValue := new([]v1.Repository)
				err := result.Decode(elemValue)
				if err != nil {
					log.Println("[ERROR]", err)
					break
				}
				results = *elemValue
			} else {
				elemValue := new([]v1.RepositoryWithOutApplication)
				err := result.Decode(elemValue)
				if err != nil {
					log.Println("[ERROR]", err)
					break
				}
				for i, each := range *elemValue {
					results[i].Type = each.Type
					results[i].Token = each.Token
					results[i].Applications = nil
				}
			}
		}
	}
	return results
}

func (c companyRepository) GetApplicationsByCompanyId(id string, option v1.CompanyQueryOption) []v1.Application {
	var results []v1.Application
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": id}}
	query["$and"] = and
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
		elemValue := new([]v1.Application)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = *elemValue
	}
	return results
}

func (c companyRepository) GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption) []v1.Application {
	var results []v1.Application
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": id}}
	query["$and"] = and
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
		elemValue := new([]v1.Repository)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, each := range *elemValue {
			if _type == each.Type {
				results = append(results, each.Applications...)
			}
		}
	}
	return results
}

func (c companyRepository) Store(company v1.Company) error {
	coll := c.manager.Db.Collection(CompanyCollection)
	_, err := coll.InsertOne(c.manager.Ctx, company)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func (c companyRepository) Update(company v1.Company, companyUpdateOption v1.CompanyUpdateOption) {
	panic("implement me")
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

func NewCompanyRepository(timeout int) repository.CompanyRepository {
	return &companyRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}

}
