package middleware

import (
	"context"
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/sirupsen/logrus"
)

// JWT is an middleware which parses, validates and passes jwt claims in the
// context of request
type JWT struct {
	JWTSecret string
	Next      http.Handler
}

func (j JWT) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		logrus.Info("No authorization header, guest request")
		j.Next.ServeHTTP(w, r)
		return
	}

	jwt, err := jws.ParseJWT([]byte(authToken))
	if err != nil {
		logrus.WithError(err).Error("Error while decoding JWT token")
		j.Next.ServeHTTP(w, r)
		return
	}

	err = jwt.Validate([]byte(j.JWTSecret), crypto.SigningMethodHS512)
	if err != nil {
		logrus.WithError(err).Error("Error while validating JWT token")
		j.Next.ServeHTTP(w, r)
		return
	}

	r = r.WithContext(withClaims(r.Context(), jwt.Claims()))
	j.Next.ServeHTTP(w, r)
}

func withClaims(ctx context.Context, claims jwt.Claims) context.Context {
	return context.WithValue(ctx, "claims", claims)
}

func Claims(ctx context.Context) jwt.Claims {
	return ctx.Value("claims").(jwt.Claims)
}
