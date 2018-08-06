package validate

import "errors"

// Required is validator which requires string to contain something
type Required struct {
	Message string
	_       struct{} // force named fields
}

// Valid checks whether provided v is empty
func (r Required) Valid(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		panic("invalid data type in Required validation")
	}
	if val == "" {
		return errors.New(r.Message)
	}
	return nil
}
