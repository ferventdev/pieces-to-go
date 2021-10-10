package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

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

	// Server startup
	s := http.Server{
		Addr:    ":8080",
		Handler: e,
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
