package routes

import (
	"github.com/labstack/echo/v4"
	"medilane-api/config"
	"medilane-api/core/errorHandling"
	accRoute "medilane-api/packages/accounts/routes"
	cartRoute "medilane-api/packages/cart/routes"
	drugStoreRoute "medilane-api/packages/drugstores/routes"
	productRoute "medilane-api/packages/medicines/routes"
	"medilane-api/packages/notification/routes"
	orderRoute "medilane-api/packages/order/routes"
	promotionRoute "medilane-api/packages/promotion/routes"
	settingRoute "medilane-api/packages/settings/routes"
	s "medilane-api/server"
	"time"

	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *s.Server, config *config.Config) {
	// middleware
	server.Echo.Debug = config.DevMode

	server.Echo.Use(middlewareLogging)
	server.Echo.Use(middleware.Recover())
	server.Echo.Use(errorHandling.Recover)
	server.Echo.Use(middleware.CORS())
	server.Echo.Use(middleware.RemoveTrailingSlash())
	server.Echo.Use(middleware.Secure())
	//server.Echo.Use(middleware.Gzip())
	server.Echo.Use(middleware.RequestID())

	// Or can use EchoWrapHandler func with configurations.
	url := echoSwagger.URL(config.SwaggerDocUrl) //The url pointing to API definition
	server.Echo.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))
	appRoute := server.Echo.Group("/api/v1")

	accRoute.ConfigureAccountRoutes(appRoute, server)
	productRoute.ConfigureProductRoutes(appRoute, server)
	drugStoreRoute.ConfigureDrugStoreRoutes(appRoute, server)
	promotionRoute.ConfigureAccountRoutes(appRoute, server)
	cartRoute.ConfigureCartRoutes(appRoute, server)
	orderRoute.ConfigureOrderRoutes(appRoute, server)
	settingRoute.ConfigureSettingtRoutes(appRoute, server)
	routes.ConfigureNotificationRoutes(appRoute, server)
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

//func errorHandler(err error, c echo.Context) {
//	report, ok := err.(*echo.HTTPError)
//	if ok {
//		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
//	} else {
//		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
//	}
//
//	makeLogEntry(c).Error(report.Message)
//	c.HTML(report.Code, report.Message.(string))
//}
