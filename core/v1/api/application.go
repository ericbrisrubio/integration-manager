package api

import "github.com/labstack/echo/v4"

// Application application api operations
type Application interface {
	GetById(context echo.Context) error
	Get(context echo.Context) error
}
