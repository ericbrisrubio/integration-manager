package v1

import (
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/repository/v1/mongo"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	GithubEventRouter(g.Group("/githubs"))

}


func GithubEventRouter(g *echo.Group) {
	var observers [] service.Observer
	observers= append(observers, logic.NewCiCoreEventService(logic.NewHttpClientService()))
	observers= append(observers, logic.NewProcessInventoryEventService(logic.NewHttpClientService()))
	githubService:=logic.NewGithubService(logic.NewCompanyService(mongo.NewCompanyRepository(3000)),nil,logic.NewHttpClientService())
	githubApi:=NewGithubApi(githubService,logic.NewCompanyService(mongo.NewCompanyRepository(3000)),observers)
	g.POST("", githubApi.ListenEvent)
}


