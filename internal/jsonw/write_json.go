package jsonw

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Write is function for writing http response to provided ResponseWriter.
// Any invalid Response will cause panic.
func Write(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.HTTPStatus())

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(r)
	if err != nil {
		panic(err)
	}

	buf.WriteTo(w)
}
