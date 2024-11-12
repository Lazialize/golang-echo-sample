package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(newEcho),
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
