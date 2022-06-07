package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/repository"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
)

type agentService struct {
	repo repository.AgentRepository
}

func (a agentService) GetByName(name string) (v1.Agent, error) {
	agent, err := a.repo.GetByName(name)
	if err != nil {
		return v1.Agent{}, err
	}
	return agent, nil
}

func (a agentService) Store(agent v1.Agent) error {
	return a.repo.Store(agent)
}

// NewAgentsService returns Agent type service
func NewAgentsService(repo repository.AgentRepository) service.Agent {
	return &agentService{
		repo: repo,
	}
}
