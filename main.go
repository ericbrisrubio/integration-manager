package main

import (
	"github.com/klovercloud-ci-cd/integration-manager/api"
	"github.com/klovercloud-ci-cd/integration-manager/config"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/dependency"
	_ "github.com/klovercloud-ci-cd/integration-manager/docs"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

// @title integration-manager API
// @description integration-manager API
func main() {
	e := config.New()
	initCompany()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

func initCompany() {
	companyService := dependency.GetV1CompanyService()
	companyService.Store(v1.Company{
		MetaData: v1.CompanyMetadata(struct {
			Labels                    map[string]string
			NumberOfConcurrentProcess int64
			TotalProcessPerDay        int64
		}{Labels: nil, NumberOfConcurrentProcess: 10, TotalProcessPerDay: 10}),
		Id:     config.CompanyId,
		Name:   config.CompanyName,
		Status: enums.ACTIVE,
	})
}
