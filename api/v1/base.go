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
	githubApi:=NewGithubApi(dependency.GetGithubService(),dependency.GetCompanyService(),dependency.GetObservers())
	g.POST("", githubApi.ListenEvent)
}

func CompanyRouter(g *echo.Group){
	companyApi:=NewCompanyApi(dependency.GetCompanyService(),nil)
	g.POST("",companyApi.Save)
	g.GET("/:id",companyApi.GetById)
	g.GET("/:id/repositories",companyApi.GetRepositoriesById)
}

func RepositoryRouter(g *echo.Group){
	repositoryApi:=NewRepositoryApi(dependency.GetCompanyService(),nil)
	g.POST("",repositoryApi.Save)
	g.GET("/:id",repositoryApi.GetById)
	g.GET("/:id/applications",repositoryApi.GetApplicationsById)
}
