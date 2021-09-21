package v1

import (
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/repository/v1/mongo"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	GithubEventRouter(g.Group("/githubs"))

}


func GithubEventRouter(g *echo.Group) {
	githubService:=NewGithubApi(logic.NewGithubService(logic.NewCompanyService(mongo.NewCompanyRepository(3000)),nil,logic.NewHttpClientService()),logic.NewCompanyService(mongo.NewCompanyRepository(3000)))
	g.POST("", githubService.ListenEvent)
}


