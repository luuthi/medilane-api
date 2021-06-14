package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	log2 "github.com/labstack/gommon/log"
	accRoute "medilane-api/packages/accounts/routes"
	drugStoreRoute "medilane-api/packages/drugstores/routes"
	promotionRoute "medilane-api/packages/promotion/routes"
	s "medilane-api/server"
	"net/http"
	"time"

	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *s.Server) {
	// middleware

	server.Echo.Use(middlewareLogging)
	server.Echo.HTTPErrorHandler = errorHandler
	server.Echo.Logger.SetLevel(log2.DEBUG)

	server.Echo.Use(middleware.CORS())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
	appRoute := server.Echo.Group("/api/v1")

	accRoute.ConfigureAccountRoutes(appRoute, server)
	//mediRoute.ConfigureAccountRoutes(appRoute, server)
	drugStoreRoute.ConfigureDrugStoreRoutes(appRoute, server)

	promotionRoute.ConfigureAccountRoutes(appRoute, server)
}

func makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}
func middlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func errorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}
