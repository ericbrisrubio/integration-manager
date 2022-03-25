package dependency

import (
	"github.com/klovercloud-ci-cd/integration-manager/config"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/logic"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/klovercloud-ci-cd/integration-manager/repository/v1/inmemory"
	"github.com/klovercloud-ci-cd/integration-manager/repository/v1/mongo"
)

// GetV1GithubService returns Git service
func GetV1GithubService() service.Git {
	return logic.NewGithubService(GetV1CompanyService(), nil, logic.NewHttpClientService())
}

// GetMockV1GithubService returns Git service
func GetMockV1GithubService() service.Git {
	return logic.NewGithubService(GetV1MockCompanyService(), nil, logic.NewHttpClientService())
}

// GetV1BitbucketService returns Git service
func GetV1BitbucketService() service.Git {
	return logic.NewBitBucketService(GetV1CompanyService(), nil, logic.NewHttpClientService())
}

// GetV1MockBitbucketService returns Git service
func GetV1MockBitbucketService() service.Git {
	return logic.NewBitBucketMockService(GetV1MockCompanyService(), nil, logic.NewHttpClientService())
}

// GetV1MockGithubService returns Git service
func GetV1MockGithubService() service.Git {
	return logic.NewGithubMockService(GetV1MockCompanyService(), nil, logic.NewHttpClientService())
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
		company = logic.NewCompanyService(mongo.NewCompanyRepository(3000), mongo.NewApplicationMetadataRepository(3000), logic.NewHttpClientService())
	} else {
		company = logic.NewCompanyService(inmemory.NewCompanyRepository(3000), nil, logic.NewHttpClientService())
	}
	return company
}

// GetV1MockCompanyService returns Company service
func GetV1MockCompanyService() service.Company {
	var company service.Company
	company = logic.NewMockCompanyService(inmemory.NewCompanyRepository(3000), logic.NewHttpClientService())
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

// GetV1SearchService returns Search service
func GetV1SearchService() service.Search {
	var search service.Search
	search = logic.NewSearchService(mongo.NewApplicationMetadataRepository(3000))
	return search
}
