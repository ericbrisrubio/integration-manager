package api

import "github.com/labstack/echo/v4"

type Git interface {
	ListenEvent(context echo.Context) error
}
