package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/strwys/fms/internal/redis"
	"github.com/strwys/fms/internal/service"
)

type AuthHandler interface {
	Login(c echo.Context) error
}

type authHandler struct {
	authSvc service.AuthService
	cache   redis.RedisCache
}

func NewAuthHandler(e *echo.Echo, authSvc service.AuthService, cache redis.RedisCache) AuthHandler {
	handler := &authHandler{
		authSvc: authSvc,
		cache:   cache,
	}

	return handler
}

func (h *authHandler) Login(c echo.Context) error {
	if !h.cache.Allow() {
		c.JSON(400, "Request rate limit exceed...")
	}

	return c.JSON(200, "Request allowed...")
}
