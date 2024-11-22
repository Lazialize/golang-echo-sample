package api

import (
	"lazialize/go-echo-sample/internal/echosample/presentation"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewV2Route() fx.Option {
	return fx.Module(
		"v2route",
		fx.Provide(
			fx.Private,
			fx.Annotate(
				registerV2Routes,
				fx.ParamTags("", `group:"routes"`),
			),
		),
		fx.Invoke(func(g *echo.Group) {}),
	)
}

func registerV2Routes(e *echo.Echo, routes []presentation.Route) *echo.Group {
	g := e.Group("v2")

	for _, route := range routes {
		g.Add(route.Method(), route.Pattern(), route.Handle)
	}

	return g
}
