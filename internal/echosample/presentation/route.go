package presentation

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Route interface {
	Handle(ctx echo.Context) error
	Pattern() string
	Method() string
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
