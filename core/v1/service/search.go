package service

import v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"

type Search interface {
	SearchReposAndAppsByCompanyIdAndName(companyId, name string, reposOption, appsOption bool) v1.SearchData
}