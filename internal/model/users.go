package model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kerak19/template/internal/repo/usersdb"
	"github.com/lhecker/argon2"
)

// Users is an upper level interface to database
type Users struct {
	DB     *sql.DB
	Hasher *argon2.Config
}

// CreateUser created new user with provided email and password.
func (u Users) CreateUser(ctx context.Context, login, password string) (usersdb.User, error) {
	bytePassword, err := u.Hasher.Hash([]byte(password), nil)
	if err != nil {
		return usersdb.User{}, err
	}
	return usersdb.CreateUser(ctx, u.DB, login, bytePassword.Encode())
}

// LoginUser is checking whether provided credentials are valid
func (u Users) LoginUser(ctx context.Context, login, password string) (usersdb.User, error) {
	user, err := usersdb.FetchUser(ctx, u.DB, login)
	if err != nil {
		return user, err
	}
	rPass, err := argon2.Decode(user.Pass)
	if err != nil {
		return user, err
	}
	match, err := rPass.Verify([]byte(password))
	if err != nil {
		return user, err
	}
	if !match {
		return user, errors.New("invalid login or password")
	}
	return user, nil
}
