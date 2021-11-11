package v1

import (
	"github.com/klovercloud-ci/dependency"
	"github.com/labstack/echo/v4"
)

// Router api/v1 base router
func Router(g *echo.Group) {
	GithubEventRouter(g.Group("/githubs"))
	CompanyRouter(g.Group("/companies"))
	RepositoryRouter(g.Group("/repositories"))
	ApplicationRouter(g.Group("/applications"))
}

// GithubEventRouter api/v1/githubs/* router
func GithubEventRouter(g *echo.Group) {
	githubApi := NewGithubApi(dependency.GetV1GithubService(), dependency.GetV1CompanyService(), dependency.GetV1ProcessInventoryEventService(), dependency.GetV1Observers())
	g.POST("", githubApi.ListenEvent)
}

// CompanyRouter api/v1/companies/* router
func CompanyRouter(g *echo.Group) {
	companyApi := NewCompanyApi(dependency.GetV1CompanyService(), nil)
	g.POST("", companyApi.Save, AuthenticationAndAuthorizationHandler)
	g.GET("", companyApi.Get, AuthenticationAndAuthorizationHandler)
	g.GET("/:id", companyApi.GetById, AuthenticationAndAuthorizationHandler)
	g.GET("/:id/repositories", companyApi.GetRepositoriesById, AuthenticationAndAuthorizationHandler)
	g.PUT("/:id/repositories", companyApi.UpdateRepositories, AuthenticationAndAuthorizationHandler)
}

// RepositoryRouter api/v1/repositories/* router
func RepositoryRouter(g *echo.Group) {
	repositoryApi := NewRepositoryApi(dependency.GetV1CompanyService(), nil)
	g.GET("/:id", repositoryApi.GetById, AuthenticationAndAuthorizationHandler)
	g.GET("/:id/applications", repositoryApi.GetApplicationsById, AuthenticationAndAuthorizationHandler)
}

// ApplicationRouter api/v1/applications/* router
func ApplicationRouter(g *echo.Group) {
	repositoryApi := NewApplicationApi(dependency.GetV1CompanyService(), nil)
	//companyId, repositoryId via query param
	g.POST("", repositoryApi.Update, AuthenticationAndAuthorizationHandler)
}
