package api

import "github.com/labstack/echo/v4"

type Application interface {
	UpdateApplication(context echo.Context) error
}
