package dependency

import (
	"github.com/klovercloud-ci/config"
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/klovercloud-ci/repository/v1/inmemory"
	"github.com/klovercloud-ci/repository/v1/mongo"
)

// GetV1GithubService returns Git service
func GetV1GithubService() service.Git {
	var observers []service.Observer
	observers = append(observers, logic.NewCiCoreEventService(logic.NewHttpClientService()))
	observers = append(observers, logic.NewProcessInventoryEventObserverService(logic.NewHttpClientService()))
	return logic.NewGithubService(GetV1CompanyService(), nil, logic.NewHttpClientService())
}

// GetV1Observers returns Observer services
func GetV1Observers() []service.Observer {
	var observers []service.Observer
	observers = append(observers, logic.NewCiCoreEventService(logic.NewHttpClientService()))
	observers = append(observers, logic.NewProcessInventoryEventObserverService(logic.NewHttpClientService()))
	observers = append(observers, logic.NewLogEventService(logic.NewHttpClientService()))
	observers = append(observers, logic.NewProcessEventService(logic.NewHttpClientService()))
	return observers
}

// GetV1CompanyService returns Company service
func GetV1CompanyService() service.Company {
	var company service.Company
	if config.Database == enums.MONGO {
		company = logic.NewCompanyService(mongo.NewCompanyRepository(3000), logic.NewHttpClientService())
	} else {
		company = logic.NewCompanyService(inmemory.NewCompanyRepository(3000), logic.NewHttpClientService())
	}
	return company
}

// GetV1ProcessInventoryEventService returns ProcessInventoryEvent service
func GetV1ProcessInventoryEventService() service.ProcessInventoryEvent {
	return logic.NewProcessInventoryEventService(logic.NewHttpClientService())
}

// GetV1JwtService returns Jwt service
func GetV1JwtService() service.Jwt {
	return logic.NewJwtService()
}
