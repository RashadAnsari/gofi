package router

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"github.com/RashadAnsari/gofi/internal/app/gofi/config"

	"github.com/labstack/echo-contrib/prometheus"
)

func New(cfg config.Config) *echo.Echo {
	e := echo.New()

	debug := logrus.IsLevelEnabled(logrus.DebugLevel)

	e.Debug = debug

	e.HidePort = true
	e.HideBanner = true

	e.Server.ReadTimeout = cfg.Server.ReadTimeout
	e.Server.WriteTimeout = cfg.Server.WriteTimeout

	recoverConfig := middleware.DefaultRecoverConfig
	recoverConfig.DisablePrintStack = !debug
	e.Use(middleware.RecoverWithConfig(recoverConfig))

	prometheus.NewPrometheus("gofi", middleware.DefaultSkipper).Use(e)

	if debug {
		e.GET("/debug/pprof/*", echo.WrapHandler(http.DefaultServeMux))
		e.Use(middleware.CORS())
	}

	e.Server.SetKeepAlivesEnabled(false)

	return e
}
