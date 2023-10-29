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
)

func RunServer() {
	cfg := config.NewConfig()

	_, err := cfg.PostgreSQLConnect()
	if err != nil {
		log.Fatal("error connecting to database: ", err.Error())
	}

	e := echo.New()

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

	// Starting server
	go func() {
		if cfg.App.HTTPPort == "" {
			cfg.App.HTTPPort = "8000"
		}

		err := e.Start(":" + cfg.App.HTTPPort)
		if err != nil {
			log.Fatal("error starting server: ", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// // Block until a signal is received.
	<-quit

	// log.Println("server shutdown of 5 second")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e.Shutdown(ctx)
}
