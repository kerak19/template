package reqctx

import (
	"context"

	"github.com/kerak19/template/internal/repo/usersdb"
)

// WithUser extends context with request's issuer
func WithUser(ctx context.Context, user usersdb.User) context.Context {
	return context.WithValue(ctx, keyUser, user)
}

// User returns user contained in request's context
func User(ctx context.Context) usersdb.User {
	return ctx.Value(keyUser).(usersdb.User)
}
