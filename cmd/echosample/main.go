package main

import (
	"context"
	"lazialize/go-echo-sample/internal/echosample/presentation/api"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	v1Module := api.NewV1Route()
	fx.New(
		fx.Provide(newEcho),
		v1Module,
		fx.Invoke(func(e *echo.Echo) {}),
	).Run()
}

func newEcho(lc fx.Lifecycle) *echo.Echo {
	e := echo.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go e.Start(":8080")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Close()
		},
	})

	return e
}
