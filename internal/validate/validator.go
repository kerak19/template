package validate

// Validator is checking whether v is passing all restrictions provided by Validator
type Validator interface {
	Valid(v interface{}) error
}
