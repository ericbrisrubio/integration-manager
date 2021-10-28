package dependency

import (
	"github.com/klovercloud-ci/config"
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/klovercloud-ci/repository/v1/in_memory"
	"github.com/klovercloud-ci/repository/v1/mongo"
)

func GetV1GithubService() service.Git {
	var observers []service.Observer
	observers = append(observers, logic.NewCiCoreEventService(logic.NewHttpClientService()))
	observers = append(observers, logic.NewProcessInventoryEventObserverService(logic.NewHttpClientService()))
	return logic.NewGithubService(GetV1CompanyService(), nil, logic.NewHttpClientService())
}

func GetV1Observers() []service.Observer {
	var observers []service.Observer
	observers = append(observers, logic.NewCiCoreEventService(logic.NewHttpClientService()))
	observers = append(observers, logic.NewProcessInventoryEventObserverService(logic.NewHttpClientService()))
	observers = append(observers, logic.NewLogEventService(logic.NewHttpClientService()))
	observers = append(observers, logic.NewProcessEventService(logic.NewHttpClientService()))
	return observers
}

func GetV1CompanyService() service.Company {
	var company service.Company
	if config.Database == enums.Mongo {
		company = logic.NewCompanyService(mongo.NewCompanyRepository(3000),logic.NewHttpClientService())
	} else {
		company = logic.NewCompanyService(in_memory.NewCompanyRepository(3000),logic.NewHttpClientService())
	}
	return company
}
func GetV1ProcessInventoryEventService()service.ProcessInventoryEvent{
	return logic.NewProcessInventoryEventService(logic.NewHttpClientService())
}
func GetJwtService()service.JwtService{
	return logic.NewJwtService()
}