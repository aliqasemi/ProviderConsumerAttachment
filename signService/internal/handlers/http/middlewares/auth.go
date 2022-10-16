package middleware

import (
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func AuthApi(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authContext := &services.AuthContext{Context: c}
		return next(authContext)
	}
}

func ValidateToken() echo.MiddlewareFunc {
	jwtConfig := middleware.JWTConfig{
		SigningKey: services.JwtKey,
		Claims:     &services.JWTClaim{},
	}
	return middleware.JWTWithConfig(jwtConfig)
}

func Authorize(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authUser := services.AuthBuilder(c)
			if authUser.HasRole(role) {
				return next(c)
			} else {
				return &echo.HTTPError{
					Message: "Unauthorized",
					Code:    http.StatusForbidden,
				}
			}
		}
	}
}
