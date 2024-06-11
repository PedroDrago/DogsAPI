package validator

import "regexp"

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Check(ok bool, key string, message string) {
	if !ok {
		if _, exists := v.Errors[key]; !exists {
			v.Errors[key] = message
		}
	}
}

func (v *Validator) Match(target string, rx *regexp.Regexp) bool {
	return rx.MatchString(target)
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}
