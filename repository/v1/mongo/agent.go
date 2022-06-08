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

// AgentCollection collection name
var (
	AgentCollection = "agentCollection"
)

type agentRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (a agentRepository) GetByName(name string) (v1.Agent, error) {
	var agent v1.Agent
	query := bson.M{
		"$and": []bson.M{
			{"name": name},
		},
	}
	coll := a.manager.Db.Collection(AgentCollection)
	result, err := coll.Find(a.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
		return agent, err
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Agent)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		agent = *elemValue
	}
	return agent, nil
}

func (a agentRepository) Update(oldAgent v1.Agent) error {
	coll := a.manager.Db.Collection(AgentCollection)
	query := bson.M{
		"$and": []bson.M{
			{"name": oldAgent.Name},
		},
	}
	update := bson.M{
		"$set": oldAgent,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := coll.FindOneAndUpdate(a.manager.Ctx, query, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (a agentRepository) Store(agent v1.Agent) error {
	oldAgent, _ := a.GetByName(agent.Name)
	if oldAgent.Name == "" {
		coll := a.manager.Db.Collection(AgentCollection)
		_, err := coll.InsertOne(a.manager.Ctx, agent)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
		}
		return nil
	} else {
		err := a.Update(agent)
		if err != nil {
			return err
		}
		return nil
	}
}

// NewAgentsRepository returns AgentRepository type object
func NewAgentsRepository(timeout int) repository.AgentRepository {
	return &agentRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
