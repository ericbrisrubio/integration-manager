package service

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
)

// Agent related operations.
type Agent interface {
	Store(agents v1.Agent) error
}
