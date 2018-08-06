package request

import (
	"net/http"

	"github.com/kerak19/template/internal/jsonw"
)

// Success writes Success response to provided ResponseWriter along with body
func Success(w http.ResponseWriter, body interface{}) {
	payload := jsonw.Success{
		Status: "success",
		Data:   body,
	}

	jsonw.Write(w, payload)
}

// Fail writes Fail response to provided ResponseWriter along with body. Fail
// should be used for non critical errors, like validation
func Fail(w http.ResponseWriter, body interface{}) {
	payload := jsonw.Fail{
		Status: "fail",
		Data:   body,
	}

	jsonw.Write(w, payload)
}

// Error writes Error response to provided ResponseWriter along with body and code
func Error(w http.ResponseWriter, body interface{}, code int) {
	payload := jsonw.Error{
		Status: "error",
		Data:   body,
		Code:   code,
	}

	jsonw.Write(w, payload)
}
