package routes

import (
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/handlers/http/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) error {
	e.GET("/users", controllers.Index)
	e.POST("/register", controllers.Register)
	return nil
}
