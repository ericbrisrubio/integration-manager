package api

import "github.com/labstack/echo/v4"

// Application application api operations
type Application interface {
	Update(context echo.Context) error
	GetById(context echo.Context) error
}
