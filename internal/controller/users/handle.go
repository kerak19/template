package users

import (
	"context"

	"github.com/kerak19/template/internal/config"
	"github.com/kerak19/template/internal/repo/usersdb"
)

// Users is handling communication between handlers and database
type Users interface {
	CreateUser(ctx context.Context, login, password string) (usersdb.User, error)
	LoginUser(ctx context.Context, login, password string) (usersdb.User, error)
}

// Handle is an type gathering user's handlers
type Handle struct {
	Users  Users
	Config config.Config
}
