package api

import "github.com/labstack/echo/v4"

type Github interface {
	ListenEvent(context echo.Context) error
}

