package main

import (
	"fmt"
	"gold-dashboard-be/internal/apps/rest/middlewares"
	"gold-dashboard-be/internal/logger"
	"gold-dashboard-be/internal/pkg/validators"
	"os"
	"strings"

	echo_middlewares "github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load environment variables with error: " + err.Error())
	}

	e := echo.New()

	// Initialize validator
	e.Validator = validators.NewValidator()

	if err := logger.SetupLogger(); err != nil {
		panic("Failed to setup logger with error: " + err.Error())
	}

	e.Use(echo_middlewares.CORSWithConfig(echo_middlewares.CORSConfig{
		AllowOrigins: strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
	}))
	e.Use(middlewares.GenerateRequestID)
	e.Use(middlewares.Log)
	e.Use(middlewares.ContextTimeoutMiddleware)

	module := Module{}
	module.New(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}