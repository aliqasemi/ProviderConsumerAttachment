package routes

import (
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/handlers/http/controllers"
	middleware "github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/handlers/http/middlewares"
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) error {
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	authApi := e.Group("", middleware.AuthApi)
	authApi = authApi.Group("", middleware.ValidateToken())
	authApi.GET("/users", controllers.Index, middleware.Authorize("user"))
	return nil
}
