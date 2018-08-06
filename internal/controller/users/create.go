package users

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kerak19/template/internal/request"
	"github.com/kerak19/template/internal/validate"
)

var registerValidators = map[string][]validate.Validator{
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

// Create creates new user
func (h Handle) Create(w http.ResponseWriter, r *http.Request,
	_ httprouter.Params) {
	data := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		request.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	errors := validate.Validate(data, loginValidators)
	if len(errors) != 0 {
		request.Fail(w, errors)
		return
	}

	user, err := h.Users.CreateUser(r.Context(), data.Login, data.Password)
	if err != nil {
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	request.Success(w, user)
}
