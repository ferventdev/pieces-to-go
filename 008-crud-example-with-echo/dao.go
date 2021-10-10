package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type UserRepo struct {
	// this is a map from user's id (UUID) to User
	sync.Map
}

func (r *UserRepo) get(id uuid.UUID) (User, error) {
	u, _ := r.Load(id)
	if u == nil {
		return User{}, fmt.Errorf("user %w", EntityNotFoundErr)
	}
	return u.(User), nil
}

func (r *UserRepo) create() (User, error) {
	u := User{uuid.New(), "", "", ""}

	return u, nil
}
