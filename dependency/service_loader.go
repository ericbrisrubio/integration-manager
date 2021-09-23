package dependency

import (
	"github.com/klovercloud-ci/config"
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/klovercloud-ci/repository/v1/in_memory"
	"github.com/klovercloud-ci/repository/v1/mongo"
)

func GetGithubService()service.Git{
	var observers [] service.Observer
	observers= append(observers, logic.NewCiCoreEventService(logic.NewHttpClientService()))
	observers= append(observers, logic.NewProcessInventoryEventService(logic.NewHttpClientService()))
	return logic.NewGithubService(GetCompanyService(),nil,logic.NewHttpClientService())
}

func GetObservers()[]service.Observer{
	var observers [] service.Observer
	observers= append(observers, logic.NewCiCoreEventService(logic.NewHttpClientService()))
	observers= append(observers, logic.NewProcessInventoryEventService(logic.NewHttpClientService()))
	return observers
}

func GetCompanyService()service.Company{
	var company service.Company
	if config.Database==enums.Mongo{
		company=logic.NewCompanyService(mongo.NewCompanyRepository(3000))
	}else{
		company=logic.NewCompanyService(in_memory.NewCompanyRepository(3000))
	}
	return company
}