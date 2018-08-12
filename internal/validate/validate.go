package validate

import (
	"reflect"
)

// Validate checks whether v is valid by checking it with every provided Validator.
// If provided v is not a struct, it'll return empty errors map. If v contains
// not exported fields, Validate will panic, so it's caller responsibility to make sure
// all v fields are exported
// TODO i18n support ?
func Validate(v interface{}, validators map[string][]Validator) map[string][]string {
	errors := make(map[string][]string)
	iterateFields(v, validators, errors)
	return errors
}

func iterateFields(v interface{}, validators map[string][]Validator, errors map[string][]string) {
	valueOf := reflect.ValueOf(v)
	if valueOf.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < valueOf.NumField(); i++ {
		tag := valueOf.Type().Field(i).Tag.Get("json")
		val := valueOf.Field(i)

		switch val.Kind() {
		case reflect.Struct:
			iterateFields(val.Interface(), validators, errors)
			continue
		case reflect.Ptr:
			val = val.Elem()
		}

		validateField(tag, val, validators, errors)
	}
}

func validateField(tag string, v reflect.Value, validators map[string][]Validator,
	errors map[string][]string) {
	for _, validator := range validators[tag] {
		if err := validator.Valid(v.Interface()); err != nil {
			errors[tag] = append(errors[tag], err.Error())
		}
	}
}
