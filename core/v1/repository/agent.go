package repository

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
)

// AgentRepository  related operations
type AgentRepository interface {
	Store(agent v1.Agent) error
	GetByName(name string) (v1.Agent, error)
	Update(oldAgent v1.Agent) error
}
