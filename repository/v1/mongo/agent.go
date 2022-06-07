package mongo

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
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

func (a agentRepository) Store(agent v1.Agent) error {
	coll := a.manager.Db.Collection(AgentCollection)
	_, err := coll.InsertOne(a.manager.Ctx, agent)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

// NewAgentsRepository returns AgentsRepository type object
func NewAgentsRepository(timeout int) repository.AgentRepository {
	return &agentRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
