package v1

import (

	"github.com/labstack/echo/v4"
	"log"
)

func Router(g *echo.Group) {
	GithubEventRouter(g.Group("/githubs"))

}

func GithubEventRouter(g *echo.Group) {
	g.POST("", GetEvents)
}


func  GetEvents(context echo.Context) error {

	log.Println(context.Request())
	return nil
}