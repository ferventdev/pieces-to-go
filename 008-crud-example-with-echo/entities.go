package main

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `format:"uuid"`
	FirstName string
	LastName  string
	Email     string
}
