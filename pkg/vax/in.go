// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vax

import (
	"reflect"
)

// ErrInInvalid is the error that returns in case of an invalid value for "in" rule.
var ErrInInvalid = NewError("validate.oneof", "must be oneof value in [%v]")

// In returns a validation rule that checks if a value can be found in the given list of values.
// reflect.DeepEqual() will be used to determine if two values are equal.
// For more details please refer to https://golang.org/pkg/reflect/#DeepEqual
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
func In(values ...interface{}) InRule {
	return InRule{
		elements: values,
		err:      ErrInInvalid.AddParams(Combine(values...)),
	}
}

// InRule is a validation rule that validates if a value can be found in the given list of values.
type InRule struct {
	elements []interface{}
	err      Error
}

func (r InRule) Code(code string) Rule {
	r.err.SetCode(code)
	return r
}
func (r InRule) Msg(msg string) Rule {
	r.err.SetMessage(msg)
	return r
}

// Validate checks if the given value is valid or not.
func (r InRule) Validate(lang string, value interface{}) error {
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return nil
	}

	for _, e := range r.elements {
		if value == e || reflect.DeepEqual(e, value) {
			return nil
		}
	}
	return r.err.SetLang(lang)
}
