package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type echoHandler struct{}

func NewEchoHandler() echoHandler {
	return echoHandler{}
}

func (h echoHandler) Method() string {
	return "GET"
}

func (h echoHandler) Pattern() string {
	return "/echo"
}

func (h echoHandler) Handle(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"echo": "Hello, World!",
	})
}
