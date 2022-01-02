package mongo

import (
	"context"
	"errors"
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

func (c companyRepository) GetApplicationByApplicationId(companyId string, applicationId string) v1.Application {
	var app v1.Application
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": companyId}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValues := new(v1.Company)
		err := result.Decode(elemValues)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, each := range elemValues.Repositories {
			for _, eachApp := range each.Applications {
				if eachApp.MetaData.Id == applicationId {
					app = eachApp
					break
				}
			}
		}
	}
	return app
}

func (c companyRepository) UpdateApplication(applicationId string, app v1.Application) error {
	//TODO implement me
	panic("implement me")
}

func (c companyRepository) GetRepositoryByRepositoryId(id string) v1.Repository {
	var repo v1.Repository
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"repositories.id": id}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValues := new(v1.Company)
		err := result.Decode(elemValues)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, each := range elemValues.Repositories {
			repo = each
		}
	}
	return repo
}

func (c companyRepository) GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application {
	var app v1.Application
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": companyId, "repositories.id": repositoryId, "repositories.applications.url": applicationUrl}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValues := new(v1.Company)
		err := result.Decode(elemValues)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, each := range elemValues.Repositories {
			for _, eachApp := range each.Applications {
				app = eachApp
			}
		}
	}
	return app
}

func (c companyRepository) AppendRepositories(companyId string, repos []v1.Repository) error {

	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	company, _ := c.GetByCompanyId(companyId, option)

	company.Repositories = append(company.Repositories, repos...)
	filter := bson.M{
		"$and": []bson.M{
			{"id": companyId},
		},
	}
	update := bson.M{
		"$set": company,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := c.manager.Db.Collection(CompanyCollection)
	err := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (c companyRepository) DeleteRepositories(companyId string, repos []v1.Repository, isSoftDelete bool) error {
	var count int64
	var repositories []v1.Repository
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	company, _ := c.GetByCompanyId(companyId, option)

	if isSoftDelete {
		for i, each := range company.Repositories {
			if each.Id == repos[i].Id {
				for j := range each.Applications {
					each.Applications[j].Status = enums.INACTIVE
				}
			}
		}
	} else {
		repositories = company.Repositories
		for i := range repos {
			for j, each := range company.Repositories {
				if repos[i].Id == each.Id {
					repositories = RemoveRepository(company.Repositories, j)
					count++
				}
			}
			company.Repositories = repositories
		}
	}

	if count < 1 {
		return errors.New("Repository Id is not matched!")
	}

	filter := bson.M{
		"$and": []bson.M{
			{"id": companyId},
		},
	}
	update := bson.M{
		"$set": company,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := c.manager.Db.Collection(CompanyCollection)
	err := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (c companyRepository) AppendApplications(companyId, repositoryId string, apps []v1.Application) error {
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	company, _ := c.GetByCompanyId(companyId, option)

	for i := range company.Repositories {
		if company.Repositories[i].Id == repositoryId {
			company.Repositories[i].Applications = append(company.Repositories[i].Applications, apps...)
		}
	}
	filter := bson.M{
		"$and": []bson.M{
			{"id": companyId, "repositories.id": repositoryId},
		},
	}
	update := bson.M{
		"$set": company,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := c.manager.Db.Collection(CompanyCollection)
	err := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (c companyRepository) DeleteApplications(companyId, repositoryId string, apps []v1.Application, isSoftDelete bool) error {
	var applications []v1.Application
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
	}
	company, _ := c.GetByCompanyId(companyId, option)
	if isSoftDelete {
		for _, each := range company.Repositories {
			for j, eachApp := range each.Applications {
				for k := range apps {
					if apps[k].MetaData.Id == eachApp.MetaData.Id {
						each.Applications[j].Status = enums.INACTIVE
					}
				}
			}
		}
	} else {
		var count int64 = 0
		for i, each := range company.Repositories {
			applications = each.Applications
			if company.Repositories[i].Id == repositoryId {
				for j := range apps {
					for k := range applications {
						if each.Applications[k].MetaData.Id == apps[j].MetaData.Id {
							app := RemoveApplication(applications, k)
							applications = app
							count++
						}
					}
				}
			}
			company.Repositories[i].Applications = applications
		}
		if count < 1 {
			return errors.New("Application Id is not matched!")
		}
	}

	filter := bson.M{
		"$and": []bson.M{
			{"id": companyId, "repositories.id": repositoryId},
		},
	}
	update := bson.M{
		"$set": company,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := c.manager.Db.Collection(CompanyCollection)
	err := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
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
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, each := range elemValue.Repositories {
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
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"repositories.applications.url": url}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
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
		results = *elemValue
	}
	return results
}

func (c companyRepository) GetCompanies(option v1.CompanyQueryOption) ([]v1.Company, int64) {
	var results []v1.Company
	coll := c.manager.Db.Collection(CompanyCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	result, err := coll.Find(c.manager.Ctx, bson.D{}, &options.FindOptions{
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
		if option.LoadRepositories {
			if option.LoadApplications {
				results = append(results, *elemValue)
			} else {
				results = append(results, elemValue.GetCompanyWithRepository())
			}
		} else {
			results = append(results, elemValue.GetCompanyWithoutRepository())
		}
	}
	return results, int64(len(results))
}

func (c companyRepository) GetByCompanyId(id string, option v1.CompanyQueryOption) (v1.Company, int64) {
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
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if option.LoadRepositories {
			if option.LoadApplications {
				results = *elemValue
			} else {
				results = elemValue.GetCompanyWithRepository()
			}
		} else {
			results = elemValue.GetCompanyWithoutRepository()
		}
	}
	count, err := coll.CountDocuments(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

func (c companyRepository) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64) {
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
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if option.LoadApplications == true {
			results = elemValue.Repositories
		} else {
			var rep v1.Repository
			for _, each := range elemValue.Repositories {
				rep.Type = each.Type
				rep.Id = each.Id
				rep.Token = each.Token
				rep.Applications = nil

				results = append(results, rep)
			}
		}
	}
	return results, int64(len(results))
}

func (c companyRepository) GetApplicationsByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Application, int64) {
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
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, eachRepo := range elemValue.Repositories {
			for _, eachApp := range eachRepo.Applications {
				results = append(results, eachApp)
			}
		}
	}
	return results, int64(len(results))
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
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		var app []v1.Application
		for _, each := range elemValue.Repositories {
			if _type == each.Type {
				app = each.Applications
			}
		}
		results = app
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

// RemoveRepository removes repository from a list by index
func RemoveRepository(s []v1.Repository, i int) []v1.Repository {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// RemoveApplication removes applications from a list by index
func RemoveApplication(s []v1.Application, i int) []v1.Application {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// NewCompanyRepository returns CompanyRepository type object
func NewCompanyRepository(timeout int) repository.CompanyRepository {
	return &companyRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}

}
