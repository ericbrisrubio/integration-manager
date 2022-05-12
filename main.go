package main

import (
	"github.com/klovercloud-ci-cd/integration-manager/api"
	"github.com/klovercloud-ci-cd/integration-manager/config"
	_ "github.com/klovercloud-ci-cd/integration-manager/docs"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// @title integration-manager API
// @description integration-manager API
func main() {
	e := config.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
