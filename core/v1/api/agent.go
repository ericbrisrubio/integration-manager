package api

import (
	"github.com/labstack/echo/v4"
)

// Agents api operations
type Agent interface {
	Store(context echo.Context) error
}
