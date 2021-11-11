package service

import "github.com/klovercloud-ci-cd/integration-manager/core/v1"

// Observer Observer related operations.
type Observer interface {
	Listen(subject v1.Subject)
}
