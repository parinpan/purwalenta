package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/parinpan/purwalenta/pkg/config"
	"github.com/parinpan/purwalenta/pkg/delivery/rest/handler"
)

func GetRoutes(e *echo.Echo) *echo.Echo {
	authCfg := config.GetConfig().UserAuthentication
	authMiddleware := middleware.JWT([]byte(authCfg.SecretToken))

	userRoute, userHandler := e.Group("/user"), handler.NewUserHandler()
	userRoute.POST("/login", userHandler.Login)
	userRoute.POST("/sign-up", userHandler.SignUp)
	userRoute.POST("/verify", userHandler.Verify)
	userRoute.POST("/change-password", userHandler.ChangePassword)
	userRoute.POST("/forgot-password", userHandler.ForgotPassword)

	oauthRoute, oauthHandler := e.Group("/oauth"), handler.NewOauthHandler()
	oauthRoute.POST("/exchange", oauthHandler.Exchange)

	assessmentRoute, assessmentHandler := e.Group("/assessment", authMiddleware), handler.NewAssessmentHandler()
	assessmentRoute.GET("/personality", assessmentHandler.FindPersonalityQuestions)

	return e
}
