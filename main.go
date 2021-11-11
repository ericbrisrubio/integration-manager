package main

import (
	"github.com/klovercloud-ci-cd/integration-manager/api"
	"github.com/klovercloud-ci-cd/integration-manager/config"
)

// @title Klovercloud-ci-integration-manager API
// @description Klovercloud-ci-integration-manager API
func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
