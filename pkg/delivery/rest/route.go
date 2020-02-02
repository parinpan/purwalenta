package rest

import (
	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) *echo.Echo {
	userRoute, userHandler := e.Group("/user"), newUserHandler()
	userRoute.GET("/login", userHandler.Login)
	return e
}
