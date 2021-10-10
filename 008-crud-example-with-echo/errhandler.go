package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	DefaultServerErr  = &BaseError{http.StatusInternalServerError, "some error", "SOME_ERROR"}
	EntityNotFoundErr = &BaseError{http.StatusNotFound, "entity not found", "ENTITY_NOT_FOUND"}
	ValidationErr     = &BaseError{http.StatusBadRequest, "invalid request", "INVALID_REQUEST"}
)

type BaseError struct {
	statusCode         int
	message, errorCode string
}

type ErrorView struct {
	Time       time.Time `json:"time"`
	StatusCode int       `json:"-"`
	Message    string    `json:"message"`
	ErrorCode  string    `json:"errorCode"`
}

func (be *BaseError) Error() string {
	return be.message
}

func createView(origErr error) ErrorView {
	msg := origErr.Error()
	var be *BaseError
	if !errors.As(origErr, &be) {
		var he *echo.HTTPError
		if errors.As(origErr, &he) {
			msg, _ = he.Message.(string)
			be = &BaseError{he.Code, msg, "HTTP_ERROR"}
		} else {
			be = DefaultServerErr
		}
	}
	return ErrorView{
		Time:       time.Now(),
		StatusCode: be.statusCode,
		Message:    msg,
		ErrorCode:  be.errorCode,
	}
}
