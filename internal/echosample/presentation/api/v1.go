package api

import (
	"lazialize/go-echo-sample/internal/echosample/presentation"
	"lazialize/go-echo-sample/internal/echosample/presentation/controllers"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewV1Route() fx.Option {
	return fx.Module(
		"v1routes",
		fx.Provide(
			AsRoute(controllers.NewEchoHandler),
			fx.Annotate(
				RegisterRoutes,
				fx.ParamTags("", `group:"routes"`),
			),
		),
		fx.Invoke(func(g *echo.Group) {}),
	)
}

func RegisterRoutes(e *echo.Echo, routes []presentation.Route) *echo.Group {
	g := e.Group("v1")

	for _, route := range routes {
		g.Add(route.Method(), route.Pattern(), route.Handle)
	}

	return g
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(presentation.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
