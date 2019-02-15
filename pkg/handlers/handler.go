package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// EchoHandler is the main handler for the echoserver
func EchoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Echo!")
}

// PingHandler is the handler for the /ping endpoint
func PingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!")
}
