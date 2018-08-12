package middleware

import (
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/kerak19/template/internal/controller/middleware/reqctx"
	"github.com/sirupsen/logrus"
)

// JWT is an middleware which parses, validates and adds jwt Claims to the
// context of request
type JWT struct {
	JWTSecret string
	Next      http.Handler
}

var emptyClaims = jwt.Claims{}

func (j JWT) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")

	log := logrus.WithField("auth_token", authToken)

	if authToken == "" {
		log.Debug("No authorization header, guest request")
		r = r.WithContext(reqctx.WithClaims(r.Context(), emptyClaims))
		j.Next.ServeHTTP(w, r)
		return
	}

	jwt, err := jws.ParseJWT([]byte(authToken))
	if err != nil {
		log.WithError(err).Error("Error while decoding JWT token")
		r = r.WithContext(reqctx.WithClaims(r.Context(), emptyClaims))
		j.Next.ServeHTTP(w, r)
		return
	}

	err = jwt.Validate([]byte(j.JWTSecret), crypto.SigningMethodHS512)
	if err != nil {
		log.WithError(err).Error("Error while validating JWT token")
		r = r.WithContext(reqctx.WithClaims(r.Context(), emptyClaims))
		j.Next.ServeHTTP(w, r)
		return
	}

	withClaims := reqctx.WithClaims(r.Context(), jwt.Claims())
	r = r.WithContext(withClaims)
	j.Next.ServeHTTP(w, r)
}
