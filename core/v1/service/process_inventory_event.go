package service

import "github.com/klovercloud-ci/core/v1"

type ProcessInventoryEvent interface {
	Listen(subject v1.Subject)
}
