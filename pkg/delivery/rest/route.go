package rest

import (
	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) *echo.Echo {
	userRoute, userHandler := e.Group("/user"), newUserHandler()
	userRoute.POST("/login", userHandler.Login)
	userRoute.POST("/sign-up", userHandler.SignUp)
	return e
}
