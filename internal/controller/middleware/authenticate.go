package middleware

import (
	"context"
	"net/http"

	"github.com/kerak19/template/internal/controller/middleware/reqctx"
	"github.com/kerak19/template/internal/repo/usersdb"
	"github.com/sirupsen/logrus"
)

// Users is an interface between Authenticate middleware and database
type Users interface {
	FetchUserByID(ctx context.Context, id int64) (usersdb.User, error)
}

// Authenticate is an middleware which authenticates user based on provided
// jwt Claims
type Authenticate struct {
	Users Users
	Next  http.Handler
}

var guest = usersdb.User{
	ID:   -1,
	Role: "guest",
}

func (a Authenticate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	claims := reqctx.Claims(r.Context())

	id, ok := claims.Get("id").(float64) // for whatever reason jwt.Claims are keeping int64 as a float64
	if !ok {
		logrus.Debug("Empty claims, guest request")
		r = r.WithContext(reqctx.WithUser(r.Context(), guest))
		a.Next.ServeHTTP(w, r)
		return
	}

	user, err := a.Users.FetchUserByID(r.Context(), int64(id))
	if err != nil {
		logrus.WithField("id", id).WithError(err).Error("User not found")
		r = r.WithContext(reqctx.WithUser(r.Context(), guest))
		a.Next.ServeHTTP(w, r)
		return
	}

	r = r.WithContext(reqctx.WithUser(r.Context(), user))
	a.Next.ServeHTTP(w, r)
}
