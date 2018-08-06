package jsonw

// Response is any valid JSON response with HTTPStatus method
type Response interface {
	HTTPStatus() int
}
