package middleware

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Logger is an middleware extending logger scope
type Logger struct {
	Log  logrus.FieldLogger
	Next http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	l.Next.ServeHTTP(w, r)
}

func withLogger(ctx context.Context, logger logrus.FieldLogger) context.Context {
	return context.WithValue(ctx, "logger", logger)
}
