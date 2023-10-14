package vax

import (
	"fmt"
	"reflect"
)

// ErrMultipleOfInvalid is the error that returns when a value is not multiple of a base.
var ErrMultipleOfInvalid = NewError("validate.multiple", "must be multiple of %v")

// MultipleOf returns a validation rule that checks if a value is a multiple of the "base" value.
// Note that "base" should be of integer type.
func MultipleOf(base interface{}) MultipleOfRule {
	return MultipleOfRule{
		base: base,
		err:  ErrMultipleOfInvalid,
	}
}

// MultipleOfRule is a validation rule that checks if a value is a multiple of the "base" value.
type MultipleOfRule struct {
	base interface{}
	err  Error
}

func (r MultipleOfRule) Code(code string) Rule {
	r.err = r.err.SetCode(code)
	return r
}

func (r MultipleOfRule) Msg(msg string) Rule {
	r.err = r.err.SetMessage(msg)
	return r
}

// Error sets the error message for the rule.
func (r MultipleOfRule) Error(message string) MultipleOfRule {
	r.err = r.err.SetMessage(message)
	return r
}

// ErrorObject sets the error struct for the rule.
func (r MultipleOfRule) ErrorObject(err Error) MultipleOfRule {
	r.err = err
	return r
}

// Validate checks if the value is a multiple of the "base" value.
func (r MultipleOfRule) Validate(lang string, value interface{}) error {
	r.err = r.err.SetLang(lang)
	rv := reflect.ValueOf(r.base)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := ToInt(value)
		if err != nil {
			return err
		}
		if v%rv.Int() == 0 {
			return nil
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v, err := ToUint(value)
		if err != nil {
			return err
		}

		if v%rv.Uint() == 0 {
			return nil
		}
	default:
		return fmt.Errorf("type not supported: %v", rv.Type())
	}

	return r.err.AddParams(r.base)
}
