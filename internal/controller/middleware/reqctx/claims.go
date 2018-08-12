package reqctx

import (
	"context"

	"github.com/SermoDigital/jose/jwt"
)

// WithClaims extends context with jwt Claims
func WithClaims(ctx context.Context, claims jwt.Claims) context.Context {
	return context.WithValue(ctx, keyClaims, claims)
}

// Claims returns jwt Claims contained in context
func Claims(ctx context.Context) jwt.Claims {
	return ctx.Value(keyClaims).(jwt.Claims)
}
