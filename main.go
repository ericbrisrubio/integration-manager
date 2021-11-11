package main

import (
	"github.com/klovercloud-ci/api"
	"github.com/klovercloud-ci/config"
)

// @title Klovercloud-ci-integration-manager API
// @description Klovercloud-ci-integration-manager API
func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
