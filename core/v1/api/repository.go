package api

import "github.com/labstack/echo/v4"

type Repository interface {
	Save(context echo.Context) error
	GetById(context echo.Context) error
	GetApplicationsById(context echo.Context) error
}
