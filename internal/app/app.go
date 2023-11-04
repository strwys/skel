package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/strwys/fms/config"
	"github.com/strwys/fms/internal/redis"
)

func RunServer() {
	config := config.NewConfig()

	db, err := config.NewDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	redisClient, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	redisCache, err := redis.NewRedisCache(redisClient, config)
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
	}))

	RegistryRoute(e, db, redisCache)

	// Starting server
	go func() {
		if config.App.HTTPPort == "" {
			config.App.HTTPPort = "8000"
		}

		err := e.Start(":" + config.App.HTTPPort)
		if err != nil {
			log.Fatal("error starting server: ", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Block until a signal is received.
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e.Shutdown(ctx)
}
