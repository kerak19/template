package reqctx

import (
	"context"

	"github.com/sirupsen/logrus"
)

// WithLogger extends request's context with fieldLogger
func WithLogger(ctx context.Context, logger logrus.FieldLogger) context.Context {
	return context.WithValue(ctx, keyLogger, logger)
}

// Logger returns logrus.FieldLogger contained in context
func Logger(ctx context.Context) logrus.FieldLogger {
	return ctx.Value(keyLogger).(logrus.FieldLogger)
}
