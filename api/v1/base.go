package v1

import (
	"github.com/klovercloud-ci/dependency"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	GithubEventRouter(g.Group("/githubs"))
	CompanyRouter(g.Group("/companies"))
	RepositoryRouter(g.Group("/repositories"))
}

func GithubEventRouter(g *echo.Group) {
	githubApi := NewV1GithubApi(dependency.GetV1GithubService(), dependency.GetV1CompanyService(), dependency.GetV1Observers())
	g.POST("", githubApi.ListenEvent)
}
func CompanyRouter(g *echo.Group) {
	companyApi := NewCompanyApi(dependency.GetV1CompanyService(), nil)
	g.POST("", companyApi.Save)
	g.GET("", companyApi.GetCompanies)
	g.GET("/:id", companyApi.GetById)
	g.GET("/:id/repositories", companyApi.GetRepositoriesById)
}

func RepositoryRouter(g *echo.Group) {
	repositoryApi := NewRepositoryApi(dependency.GetV1CompanyService(), nil)
	g.POST("", repositoryApi.Save)
	g.GET("/:id", repositoryApi.GetById)
	g.GET("/:id/applications", repositoryApi.GetApplicationsById)
}
