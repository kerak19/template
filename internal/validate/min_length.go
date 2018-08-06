package validate

import (
	"errors"
	"unicode/utf8"
)

// MinLength is validator which requires string to contain at least min characters
type MinLength struct {
	Min     int
	Message string
	_       struct{} // force named fields
}

// Valid checks whether provided v is having at least Min characters
func (m MinLength) Valid(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		panic("invalid data type in MinLength validation")
	}
	if utf8.RuneCountInString(val) < m.Min {
		return errors.New(m.Message)
	}
	return nil
}
