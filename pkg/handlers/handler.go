package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// EchoHandler is the main and only handler for the echoserver
func EchoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Echo!")
}
