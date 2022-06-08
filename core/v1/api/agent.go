package api

import (
	"github.com/labstack/echo/v4"
)

// Agent api operations
type Agent interface {
	Store(context echo.Context) error
	GetByName(context echo.Context) error
}
