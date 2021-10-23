package main

import (
	"net/http"

	_ "github.com/ferventdev/pieces-to-go/008-crud-example-with-echo/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Simple CRUD example
// @version 0.0.1
// @description This is a simple server to demonstrate the CRUD operations with the Echo framework.
// @host localhost:8080
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Error handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		ev := createView(err)
		c.JSON(ev.StatusCode, ev)
	}

	// Data Access Objects
	userRepo := &UserRepo{}

	// Routes
	SetupUserApi(e, userRepo)

	// Swagger support
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Server startup
	s := http.Server{
		Addr:    ":8080",
		Handler: e,
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
