package middleware

import (
	"github-trending/model"
	"github-trending/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.KEY),
	}
	return middleware.JWTWithConfig(config)
}
