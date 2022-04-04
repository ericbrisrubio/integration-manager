package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/config"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/dependency"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/labstack/echo/v4"
)

// Router api/v1 base router
func Router(g *echo.Group) {
	GithubEventRouter(g.Group("/githubs"))
	CompanyRouter(g.Group("/companies"))
	RepositoryRouter(g.Group("/repositories"))
	ApplicationRouter(g.Group("/applications"))
	BitbucketEventRouter(g.Group("/bitbuckets"))
	SearchRouter(g.Group("/search"))
}

// BitbucketEventRouter api/v1/bitbuckets event router
func BitbucketEventRouter(g *echo.Group) {
	var bitbucketApi api.Git

	if config.Environment == string(enums.PRODUCTION) {
		bitbucketApi = NewBitbucketApi(dependency.GetV1BitbucketService(), dependency.GetV1CompanyService(), dependency.GetV1ProcessInventoryEventService(), dependency.GetV1Observers())
	} else {
		bitbucketApi = NewBitbucketApi(dependency.GetV1MockBitbucketService(), dependency.GetV1MockCompanyService(), dependency.GetV1ProcessInventoryEventService(), dependency.GetV1Observers())
	}
	g.POST("", bitbucketApi.ListenEvent)
	g.GET("", bitbucketApi.GetBranches)
}

// GithubEventRouter api/v1/githubs/* router
func GithubEventRouter(g *echo.Group) {
	var githubApi api.Git
	if config.Environment == string(enums.PRODUCTION) {
		githubApi = NewGithubApi(dependency.GetV1GithubService(), dependency.GetV1CompanyService(), dependency.GetV1ProcessInventoryEventService(), dependency.GetV1Observers())
	} else {
		githubApi = NewGithubApi(dependency.GetV1MockGithubService(), dependency.GetV1MockCompanyService(), dependency.GetV1ProcessInventoryEventService(), dependency.GetV1Observers())
	}
	g.POST("", githubApi.ListenEvent)
	g.GET("", githubApi.GetBranches)
}

// CompanyRouter api/v1/companies/* router
func CompanyRouter(g *echo.Group) {
	companyApi := NewCompanyApi(dependency.GetV1CompanyService(), nil)
	g.POST("", companyApi.Save, AuthenticationAndAuthorizationHandler)
	g.GET("", companyApi.Get, AuthenticationAndAuthorizationHandler)
	g.GET("/:id", companyApi.GetById, AuthenticationAndAuthorizationHandler)
	g.GET("/:id/repositories", companyApi.GetRepositoriesById, AuthenticationAndAuthorizationHandler)
	g.PUT("/:id/repositories", companyApi.UpdateRepositories, AuthenticationAndAuthorizationHandler)
	g.GET("/:id/applications", companyApi.GetApplicationsByCompanyIdAndRepositoryType, AuthenticationAndAuthorizationHandler)
}

// RepositoryRouter api/v1/repositories/* router
func RepositoryRouter(g *echo.Group) {
	repositoryApi := NewRepositoryApi(dependency.GetV1CompanyService(), nil)
	g.GET("/:id", repositoryApi.GetById, AuthenticationAndAuthorizationHandler)
	g.GET("/:id/applications", repositoryApi.GetApplicationsById, AuthenticationAndAuthorizationHandler)
}

// ApplicationRouter api/v1/applications/* router
func ApplicationRouter(g *echo.Group) {
	applicationApi := NewApplicationApi(dependency.GetV1CompanyService(), nil)
	//companyId, repositoryId via query param
	g.POST("", applicationApi.Update, AuthenticationAndAuthorizationHandler)
	g.GET("/:id", applicationApi.GetById, AuthenticationAndAuthorizationHandler)
	g.GET("", applicationApi.GetAll, AuthenticationAndAuthorizationHandler)
}

// SearchRouter api/v1/search/* router
func SearchRouter(g *echo.Group) {
	searchApi := NewSearchApi(dependency.GetV1SearchService())
	g.GET("", searchApi.SearchReposAndAppsByCompanyIdAndName, AuthenticationAndAuthorizationHandler)
}
