package jsonw

import "net/http"

// Success is used when request was valid and there was no error on server side
type Success struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// HTTPStatus returns http.StatusOK
func (Success) HTTPStatus() int {
	return http.StatusOK
}

// Fail is used when there was no error, but request was invalid somehow
type Fail struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// HTTPStatus returns http.StatusOK
func (Fail) HTTPStatus() int {
	return http.StatusOK
}

// Error is user when there was an error during handling request. It contains
// data with extra context and code with error code
type Error struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
}

// HTTPStatus returns any status passed to error during creation
func (e Error) HTTPStatus() int {
	return e.Code
}
