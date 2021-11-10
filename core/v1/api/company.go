package api

import (
	"github.com/labstack/echo/v4"
)

type Company interface {
	Save(context echo.Context) error
	GetById(context echo.Context) error
	GetCompanies(context echo.Context) error
	GetRepositoriesById(context echo.Context) error
	UpdateRepositories(context echo.Context) error
}
