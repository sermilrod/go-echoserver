package echoserver

import (
	"os"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sermilrod/go-echoserver/pkg/handlers"
)

// Start and configure the webserver
func Start() {
	log.SetLevel(log.INFO)
	log.SetOutput(os.Stdout)

	// Echo instance
	api := echo.New()
	api.HideBanner = true
	api.Server.Addr = ":19019"

	// Middleware
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	// Routes
	api.GET("/", handlers.EchoHandler)
	api.GET("/ping", handlers.PingHandler)
	api.GET("/healthz", handlers.HealthzHandler)

	// Start server and configure gracehttp to gracefully shut it down
	api.Logger.SetLevel(log.INFO)
	api.Logger.Fatal(gracehttp.Serve(api.Server))
}
