package main

import (
	"github.com/klovercloud-ci-cd/integration-manager/api"
	"github.com/klovercloud-ci-cd/integration-manager/config"
	_ "github.com/klovercloud-ci-cd/integration-manager/docs"
)

// @title integration-manager API
// @description integration-manager API
func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
