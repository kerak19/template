package controller

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kerak19/template/internal/config"
	"github.com/kerak19/template/internal/controller/users"
	"github.com/kerak19/template/internal/model"
	"github.com/lhecker/argon2"
	"github.com/sirupsen/logrus"
)

// Controller is an main router of application.
func Controller(ctx context.Context, db *sql.DB, config config.Config, log *logrus.Logger) http.Handler {
	router := httprouter.New()

	hasher := argon2.DefaultConfig()

	usersModel := model.Users{
		DB:     db,
		Hasher: &hasher,
	}

	users := users.Handle{
		Users:  usersModel,
		Config: config,
	}

	router.POST("/api/users/", users.Create)
	router.POST("/api/session/", users.Login)

	return router
}
