package users

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kerak19/template/internal/controller/middleware/reqctx"
	"github.com/kerak19/template/internal/request"
	"github.com/kerak19/template/internal/validate"
	"github.com/sirupsen/logrus"
)

var registerValidators = map[string][]validate.Validator{
	"login": []validate.Validator{
		validate.MinLength{
			Min:     2,
			Message: "Login is too short",
		},
	},
	"password": []validate.Validator{
		validate.MinLength{
			Min:     6,
			Message: "Password is too short",
		},
	},
}

// Create creates new user
func (h Handle) Create(w http.ResponseWriter, r *http.Request,
	_ httprouter.Params) {
	log := reqctx.Logger(r.Context())

	var data struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.WithError(err).Error("Error while decoding request body")
		request.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log = log.WithFields(logrus.Fields{
		"login":           data.Login,
		"password_length": len(data.Password),
	})

	errors := validate.Validate(data, registerValidators)
	if len(errors) != 0 {
		log.WithField("errors", errors).Debug("Validation errors") // debug only
		log.Error("Validation error during user creation")
		request.Fail(w, errors)
		return
	}

	user, err := h.Users.CreateUser(r.Context(), data.Login, data.Password)
	if err != nil {
		log.WithError(err).Error("Error during creating user")
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	log.Info("User created")
	request.Success(w, user)
}
