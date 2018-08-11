package users

import (
	"encoding/json"
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/julienschmidt/httprouter"
	"github.com/kerak19/template/internal/request"
	"github.com/kerak19/template/internal/validate"
	"github.com/sirupsen/logrus"
)

var loginValidators = map[string][]validate.Validator{
	"login": []validate.Validator{
		validate.MinLength{
			Min:     2,
			Message: "login is too short",
		},
	},
	"password": []validate.Validator{
		validate.MinLength{
			Min:     6,
			Message: "password is too short",
		},
	},
}

// Login handles user's login process
func (h Handle) Login(w http.ResponseWriter, r *http.Request,
	_ httprouter.Params) {
	data := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		logrus.WithError(err).Error("Error while decoding JSON body")
		request.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	errors := validate.Validate(data, loginValidators)
	if len(errors) != 0 {
		request.Fail(w, errors)
		return
	}

	user, err := h.Users.LoginUser(r.Context(), data.Login, data.Password)
	if err != nil {
		logrus.WithError(err).Error("Error while loging in user")
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	authToken, err := h.jwtToken(user.ID)
	if err != nil {
		logrus.WithError(err).Error("Error while generating JWT token")
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", authToken)

	request.Success(w, user)
}

var signingMethod = crypto.SigningMethodHS512

func (h Handle) jwtToken(userID int64) (string, error) {
	claims := jws.Claims{
		"id": userID,
	}
	jwt := jws.NewJWT(claims, signingMethod)
	token, err := jwt.Serialize([]byte(h.Config.JWTSecret))
	return string(token), err
}
