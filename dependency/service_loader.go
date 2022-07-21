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
	return logic.NewGithubService(nil, logic.NewHttpClientService())
}

// GetMockV1GithubService returns Git service
func GetMockV1GithubService() service.Git {
	return logic.NewGithubService(nil, logic.NewHttpClientService())
}

// GetV1BitbucketService returns Git service
func GetV1BitbucketService() service.Git {
	return logic.NewBitBucketService(nil, logic.NewHttpClientService())
}

// GetV1MockBitbucketService returns Git service
func GetV1MockBitbucketService() service.Git {
	return logic.NewBitBucketMockService(nil, logic.NewHttpClientService())
}

// GetV1MockGithubService returns Git service
func GetV1MockGithubService() service.Git {
	return logic.NewGithubMockService(nil, logic.NewHttpClientService())
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

// GetV1RepositoryService returns Repository service
func GetV1RepositoryService() service.Repository {
	var repository service.Repository
	repository = logic.NewRepositoryService(mongo.NewRepositoryRepository(3000), GetV1ApplicationService())
	return repository
}

// GetV1ApplicationMetadataService returns ApplicationMetadata service
func GetV1ApplicationMetadataService() service.ApplicationMetadataService {
	var applicationMetadata service.ApplicationMetadataService
	applicationMetadata = logic.NewApplicationMetadataService(mongo.NewApplicationMetadataRepository(3000))
	return applicationMetadata
}

// GetV1ApplicationService returns Application service
func GetV1ApplicationService() service.Application {
	var application service.Application
	application = logic.NewApplicationService(mongo.NewApplicationRepository(3000), GetV1ApplicationMetadataService(), logic.NewHttpClientService())
	return application
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

// GetV1PipelineService returns Pipeline service
func GetV1PipelineService() service.Pipeline {
	pipeline := logic.NewPipelineService(GetV1GithubService(), GetV1BitbucketService(), GetV1RepositoryService())
	return pipeline
}

// GetV1AgentService returns Search service
func GetV1AgentService() service.Agent {
	agent := logic.NewAgentsService(mongo.NewAgentsRepository(3000))
	return agent
}
