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

//func (c companyRepository) GetDashboardData(companyId string) v1.DashboardData {
//	var enabled int64
//	var disabled int64
//	query := bson.M{
//		"$and": []bson.M{
//			{"id": companyId},
//		},
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	result, err := coll.Find(c.manager.Ctx, query, nil)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	company := new(v1.Company)
//	for result.Next(context.TODO()) {
//		err := result.Decode(&company)
//		if err != nil {
//			log.Println("[ERROR]", err)
//			break
//		}
//	}
//	for _, eachRepo := range company.Repositories {
//		for _, eachApp := range eachRepo.Applications {
//			if eachApp.Webhook.Active {
//				enabled++
//			} else {
//				disabled++
//			}
//		}
//	}
//	return v1.DashboardData{
//		Repository: struct {
//			Count int64 `json:"count"`
//		}(struct{ Count int64 }{Count: int64(len(company.Repositories))}), Application: struct {
//			Webhook struct {
//				Enabled  int64 `json:"enabled"`
//				Disabled int64 `json:"disabled"`
//			} `json:"webhook"`
//		}(struct {
//			Webhook struct {
//				Enabled  int64 `json:"enabled"`
//				Disabled int64 `json:"disabled"`
//			}
//		}{Webhook: struct {
//			Enabled  int64 `json:"enabled"`
//			Disabled int64 `json:"disabled"`
//		}(struct {
//			Enabled  int64
//			Disabled int64
//		}{Enabled: enabled, Disabled: disabled})})}
//}

//func (c companyRepository) GetApplicationsByRepositoryId(repoId string, companyId string, option v1.CompanyQueryOption, status v1.StatusQueryOption) ([]v1.Application, int64) {
//	var results []v1.Application
//	var repository v1.Repository
//	query := bson.M{
//		"$and": []bson.M{{"id": companyId}},
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	result, err := coll.Find(c.manager.Ctx, query, nil)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for result.Next(context.TODO()) {
//		elemValue := new(v1.Company)
//		err := result.Decode(elemValue)
//		if err != nil {
//			log.Println("[ERROR]", err)
//			break
//		}
//		for _, eachRepo := range elemValue.Repositories {
//			if repoId == eachRepo.Id {
//				repository = eachRepo
//				break
//			}
//		}
//		startIndex, endIndex := GetPagination(int64(len(repository.Applications)), option.Pagination.Page, option.Pagination.Limit)
//		apps := repository.Applications[startIndex:endIndex]
//		for _, eachApp := range apps {
//			if eachApp.Status == status.Option {
//				results = append(results, eachApp)
//			}
//		}
//	}
//	return results, int64(len(results))
//}

//func (c companyRepository) GetApplicationByApplicationId(companyId string, repoId string, applicationId string) v1.Application {
//	var app v1.Application
//	var repo v1.Repository
//	query := bson.M{
//		"$and": []bson.M{{"id": companyId}},
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	result, err := coll.Find(c.manager.Ctx, query)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for result.Next(context.TODO()) {
//		elemValues := new(v1.Company)
//		err := result.Decode(elemValues)
//		if err != nil {
//			log.Println("[ERROR]", err)
//			break
//		}
//		for _, each := range elemValues.Repositories {
//			if repoId == each.Id {
//				repo = each
//			}
//		}
//		for _, each := range repo.Applications {
//			if applicationId == each.MetaData.Id {
//				app = each
//			}
//		}
//	}
//	return app
//}

//func (c companyRepository) UpdateApplication(companyId string, repositoryId string, applicationId string, app v1.Application) error {
//	option := v1.CompanyQueryOption{
//		Pagination:       v1.Pagination{},
//		LoadRepositories: true,
//		LoadApplications: true,
//	}
//	company, _ := c.GetByCompanyId(companyId, option)
//
//	for i, eachRepo := range company.Repositories {
//		if eachRepo.Id == repositoryId {
//			for j, eachApp := range eachRepo.Applications {
//				if eachApp.MetaData.Id == app.MetaData.Id {
//					company.Repositories[i].Applications[j] = app
//					break
//				}
//			}
//			break
//		}
//	}
//
//	filter := bson.M{
//		"$and": []bson.M{
//			{"id": companyId},
//		},
//	}
//	update := bson.M{
//		"$set": company,
//	}
//	upsert := true
//	after := options.After
//	opt := options.FindOneAndUpdateOptions{
//		ReturnDocument: &after,
//		Upsert:         &upsert,
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	res := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, &opt)
//	if res.Err() != nil {
//		log.Println("[ERROR]", res.Err())
//		return res.Err()
//	}
//	return nil
//}

//func (c companyRepository) GetRepositoryByRepositoryId(id, repositoryId string, option v1.CompanyQueryOption) v1.Repository {
//	var repo v1.Repository
//	query := bson.M{
//		"$and": []bson.M{{"id": id, "repositories.id": repositoryId}},
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	result, err := coll.Find(c.manager.Ctx, query)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for result.Next(context.TODO()) {
//		elemValues := new(v1.Company)
//		err := result.Decode(elemValues)
//		if err != nil {
//			log.Println("[ERROR]", err)
//			break
//		}
//		for _, each := range elemValues.Repositories {
//			if repositoryId == each.Id {
//				if option.LoadApplications == false && option.LoadToken == false {
//					repositoryWithOutToken := v1.Repository{
//						Id:           each.Id,
//						Type:         each.Type,
//						Token:        "",
//						Applications: nil,
//					}
//					repo = repositoryWithOutToken
//				} else if option.LoadApplications == true && option.LoadToken == false {
//					repositoryWithOutToken := v1.Repository{
//						Id:           each.Id,
//						Type:         each.Type,
//						Token:        "",
//						Applications: each.Applications,
//					}
//					repo = repositoryWithOutToken
//				} else {
//					repo = each
//				}
//			}
//		}
//	}
//	return repo
//}

//func (c companyRepository) GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application {
//	var app v1.Application
//	var repo v1.Repository
//	query := bson.M{
//		"$and": []bson.M{{"id": companyId, "repositories.id": repositoryId, "repositories.applications.url": applicationUrl}},
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	result, err := coll.Find(c.manager.Ctx, query)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for result.Next(context.TODO()) {
//		elemValues := new(v1.Company)
//		err := result.Decode(elemValues)
//		if err != nil {
//			log.Println("[ERROR]", err)
//			break
//		}
//		for _, each := range elemValues.Repositories {
//			if each.Id == repositoryId {
//				repo = each
//			}
//		}
//		for _, each := range repo.Applications {
//			if each.Url == applicationUrl {
//				app = each
//			}
//		}
//	}
//	return app
//}

//func (c companyRepository) AppendRepositories(companyId string, repos []v1.Repository) error {
//
//	option := v1.CompanyQueryOption{
//		Pagination:       v1.Pagination{},
//		LoadRepositories: true,
//		LoadApplications: true,
//	}
//	company, _ := c.GetByCompanyId(companyId, option)
//
//	company.Repositories = append(company.Repositories, repos...)
//	filter := bson.M{
//		"$and": []bson.M{
//			{"id": companyId},
//		},
//	}
//	update := bson.M{
//		"$set": company,
//	}
//	upsert := true
//	after := options.After
//	opt := options.FindOneAndUpdateOptions{
//		ReturnDocument: &after,
//		Upsert:         &upsert,
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	res := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, &opt)
//	if res.Err() != nil {
//		log.Println("[ERROR]", res.Err())
//		return res.Err()
//	}
//	return nil
//}

//func (c companyRepository) DeleteRepositories(companyId string, repos []v1.Repository) error {
//	option := v1.CompanyQueryOption{
//		Pagination:       v1.Pagination{},
//		LoadRepositories: true,
//		LoadApplications: true,
//	}
//	company, _ := c.GetByCompanyId(companyId, option)
//	company.Repositories = repos
//	filter := bson.M{
//		"$and": []bson.M{
//			{"id": companyId},
//		},
//	}
//	update := bson.M{
//		"$set": company,
//	}
//	//upsert := true
//	//after := options.After
//	//opt := options.FindOneAndUpdateOptions{
//	//	ReturnDocument: &after,
//	//	Upsert:         &upsert,
//	//}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	res := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, nil)
//	if res.Err() != nil {
//		log.Println("[ERROR]", res.Err())
//		return res.Err()
//	}
//	return nil
//}

//func (c companyRepository) AppendApplications(companyId, repositoryId string, apps []v1.Application) error {
//	option := v1.CompanyQueryOption{
//		Pagination:       v1.Pagination{},
//		LoadRepositories: true,
//		LoadApplications: true,
//	}
//	company, _ := c.GetByCompanyId(companyId, option)
//
//	for i := range company.Repositories {
//		if company.Repositories[i].Id == repositoryId {
//			company.Repositories[i].Applications = append(company.Repositories[i].Applications, apps...)
//		}
//	}
//	filter := bson.M{
//		"$and": []bson.M{
//			{"id": companyId, "repositories.id": repositoryId},
//		},
//	}
//	update := bson.M{
//		"$set": company,
//	}
//	upsert := true
//	after := options.After
//	opt := options.FindOneAndUpdateOptions{
//		ReturnDocument: &after,
//		Upsert:         &upsert,
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	res := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, &opt)
//	if res.Err() != nil {
//		log.Println("[ERROR]", res.Err())
//		return res.Err()
//	}
//	return nil
//}
//
//func (c companyRepository) DeleteApplications(companyId, repositoryId string, repos []v1.Repository) error {
//	option := v1.CompanyQueryOption{
//		Pagination:       v1.Pagination{},
//		LoadRepositories: true,
//		LoadApplications: true,
//	}
//	company, _ := c.GetByCompanyId(companyId, option)
//	company.Repositories = repos
//	filter := bson.M{
//		"$and": []bson.M{
//			{"id": companyId, "repositories.id": repositoryId},
//		},
//	}
//	update := bson.M{
//		"$set": company,
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	res := coll.FindOneAndUpdate(c.manager.Ctx, filter, update, nil)
//	if res.Err() != nil {
//		log.Println("[ERROR]", res.Err())
//		return res.Err()
//	}
//	return nil
//}

//func (c companyRepository) GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository {
//	var results v1.Repository
//	query := bson.M{
//		"$and": []bson.M{{"id": id}},
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	result, err := coll.Find(c.manager.Ctx, query)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for result.Next(context.TODO()) {
//		elemValue := new(v1.Company)
//		err := result.Decode(elemValue)
//		if err != nil {
//			log.Println("[ERROR]", err)
//			break
//		}
//		for _, each := range elemValue.Repositories {
//			for _, eachApp := range each.Applications {
//				if url == eachApp.Url {
//					r := v1.Repository{
//						Id:           each.Id,
//						Type:         each.Type,
//						Token:        each.Token,
//						Applications: each.Applications,
//					}
//					results = r
//				}
//			}
//		}
//	}
//	return results
//}

//func (c companyRepository) GetCompanyByApplicationUrl(url string) v1.Company {
//	var results v1.Company
//	query := bson.M{
//		"$and": []bson.M{{"repositories.applications.url": url}},
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	result, err := coll.Find(c.manager.Ctx, query)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for result.Next(context.TODO()) {
//		elemValue := new(v1.Company)
//		err := result.Decode(elemValue)
//		if err != nil {
//			log.Println("[ERROR]", err)
//			break
//		}
//		for i, each := range elemValue.Repositories {
//			r := v1.Repository{
//				Id:           each.Id,
//				Type:         each.Type,
//				Token:        "",
//				Applications: each.Applications,
//			}
//			elemValue.Repositories[i] = r
//		}
//		results = *elemValue
//	}
//	return results
//}

func (c companyRepository) GetCompanies(option v1.CompanyQueryOption, status v1.StatusQueryOption) ([]v1.Company, int64) {
	var results []v1.Company
	query := bson.M{
		"$and": []bson.M{{"status": status.Option}},
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
	res := new(v1.Company)
	err = result.Decode(res)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return *res
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

//func (c companyRepository) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64) {
//	var results []v1.Repository
//	query := bson.M{
//		"$and": []bson.M{{"id": id}},
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	result, err := coll.Find(c.manager.Ctx, query, nil)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for result.Next(context.TODO()) {
//		elemValue := new(v1.Company)
//		err := result.Decode(elemValue)
//		if err != nil {
//			log.Println("[ERROR]", err)
//			break
//		}
//		startIndex, endIndex := GetPagination(int64(len(elemValue.Repositories)), option.Pagination.Page, option.Pagination.Limit)
//		repo := elemValue.Repositories[startIndex:endIndex]
//		for _, each := range repo {
//			if option.LoadApplications == true && option.LoadToken == true {
//				results = elemValue.Repositories
//			} else if option.LoadApplications == true && option.LoadToken == false {
//				r := v1.Repository{
//					Id:           each.Id,
//					Type:         each.Type,
//					Token:        "",
//					Applications: each.Applications,
//				}
//				results = append(results, r)
//			} else {
//				r := v1.Repository{
//					Id:           each.Id,
//					Type:         each.Type,
//					Token:        "",
//					Applications: nil,
//				}
//				results = append(results, r)
//			}
//		}
//	}
//	return results, int64(len(results))
//}

//func (c companyRepository) GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Application {
//	var results []v1.Application
//	query := bson.M{
//		"$and": []bson.M{{"id": id, "repositories.type": _type, "repositories.applications.status": status.Option}},
//	}
//	coll := c.manager.Db.Collection(CompanyCollection)
//	result, err := coll.Find(c.manager.Ctx, query, nil)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for result.Next(context.TODO()) {
//		elemValue := new(v1.Company)
//		err := result.Decode(elemValue)
//		if err != nil {
//			log.Println("[ERROR]", err)
//			break
//		}
//		var app []v1.Application
//		for _, each := range elemValue.Repositories {
//			if _type == each.Type {
//				app = each.Applications
//			}
//		}
//		startIndex, endIndex := GetPagination(int64(len(app)), option.Pagination.Page, option.Pagination.Limit)
//		app = app[startIndex:endIndex]
//		for _, each := range app {
//			if each.Status == status.Option {
//				results = append(results, each)
//			}
//		}
//	}
//	return results
//}

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
