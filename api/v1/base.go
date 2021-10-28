package v1

import (
	"github.com/klovercloud-ci/dependency"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	GithubEventRouter(g.Group("/githubs"))
	CompanyRouter(g.Group("/companies"))
	RepositoryRouter(g.Group("/repositories"))
	ApplicationRouter(g.Group("/applications"))
}

func GithubEventRouter(g *echo.Group) {
	githubApi := NewV1GithubApi(dependency.GetV1GithubService(), dependency.GetV1CompanyService(),dependency.GetV1ProcessInventoryEventService(), dependency.GetV1Observers())
	g.POST("", githubApi.ListenEvent)
}
func CompanyRouter(g *echo.Group) {
	companyApi := NewCompanyApi(dependency.GetV1CompanyService(), nil)
	g.POST("", companyApi.Save,AuthenticationAndAuthorizationHandler)
	g.GET("", companyApi.GetCompanies,AuthenticationAndAuthorizationHandler)
	g.GET("/:id", companyApi.GetById,AuthenticationAndAuthorizationHandler)
	g.GET("/:id/repositories", companyApi.GetRepositoriesById,AuthenticationAndAuthorizationHandler)
}

func RepositoryRouter(g *echo.Group) {
	repositoryApi := NewRepositoryApi(dependency.GetV1CompanyService(), nil)
	g.POST("", repositoryApi.Save,AuthenticationAndAuthorizationHandler)
	g.GET("/:id", repositoryApi.GetById,AuthenticationAndAuthorizationHandler)
	g.GET("/:id/applications", repositoryApi.GetApplicationsById,AuthenticationAndAuthorizationHandler)
}

func ApplicationRouter(g *echo.Group) {
	repositoryApi := NewApplicationApi(dependency.GetV1CompanyService(), nil)
	//companyId, repositoryId via query param
	g.PUT("", repositoryApi.UpdateApplication,AuthenticationAndAuthorizationHandler)
}
