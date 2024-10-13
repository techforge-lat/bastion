package server

import (
	"strings"
	"time"

	"github.com/techforge-lat/bastion/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	prettylogger "github.com/rdbell/echo-pretty-logger"
)

func newEcho(conf config.Root, errorHandler echo.HTTPErrorHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// to show status code colorized
		Output: e.Logger.Output(),
		// we just remove the `error` key from the default format
		// because it is logged by our custom logger in the echo.HTTPErrorHandler
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
	}))
	e.Use(prettylogger.Logger)
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rateLimiterMemoryStore)))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Minute,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(conf.AllowedDomains, ","),
		AllowMethods: strings.Split(conf.AllowedMethods, ","),
	}))

	e.HTTPErrorHandler = errorHandler

	if conf.Env != "prod" {
		e.Debug = true
	}

	return e
}
