package presentation

import "github.com/labstack/echo/v4"

type Route interface {
	Handle(ctx echo.Context) error
	Pattern() string
	Method() string
}
