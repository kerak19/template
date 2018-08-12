package users

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kerak19/template/internal/controller/middleware/reqctx"
	"github.com/kerak19/template/internal/request"
)

// ByID returns info of user's by id provided in authorization token
func (h Handle) ByID(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params) {
	user := reqctx.User(r.Context())
	request.Success(w, user)
}
