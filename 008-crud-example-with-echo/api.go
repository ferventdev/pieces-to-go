package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserApi struct {
	repo *UserRepo
}

func SetupUserApi(e *echo.Echo, r *UserRepo) *UserApi {
	api := &UserApi{r}
	e.GET("/users/:id", api.get)
	return api
}

func (a *UserApi) get(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return fmt.Errorf("%w: id parameter missing or invalid", ValidationErr)
	}
	u, err := a.repo.get(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
