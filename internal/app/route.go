package app

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/strwys/fms/internal/handler"
	"github.com/strwys/fms/internal/redis"
	"github.com/strwys/fms/internal/repository"
	"github.com/strwys/fms/internal/service"
)

func RegistryRoute(e *echo.Echo, db *sql.DB, redis redis.RedisCache) {
	authRepo := repository.NewAuthRepository(db)
	authSvc := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(e, authSvc, redis)

	e.GET("/login", authHandler.Login)
}
