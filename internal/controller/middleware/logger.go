package middleware

import (
	"net/http"

	"github.com/kerak19/template/internal/controller/middleware/reqctx"
	"github.com/sirupsen/logrus"
)

// Logger is an middleware extending logger scope
type Logger struct {
	Log  logrus.FieldLogger
	Next http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := reqctx.User(r.Context())

	log := l.Log.WithFields(logrus.Fields{
		"role": user.Role,
		"ip":   r.RemoteAddr,
	})

	if user.Role != "guest" {
		log = log.WithFields(logrus.Fields{
			"login": user.Login,
			"id":    user.ID,
		})
	}

	r = r.WithContext(reqctx.WithLogger(r.Context(), log))
	l.Next.ServeHTTP(w, r)
}
