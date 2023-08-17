// Copyright 2018 Qiang Xue, Google LLC. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vax

// ErrNotInInvalid is the error that returns when a value is in a list.
var ErrNotInInvalid = NewError("validate.notoneof", "cannot be one of several values [%v]")

// NotIn returns a validation rule that checks if a value is absent from the given list of values.
// Note that the value being checked and the possible range of values must be of the same type.
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
func NotIn(values ...interface{}) NotInRule {
	return NotInRule{
		elements: values,
		err:      ErrNotInInvalid.AddParams(Combine(values...)),
	}
}

// NotInRule is a validation rule that checks if a value is absent from the given list of values.
type NotInRule struct {
	elements []interface{}
	err      Error
}

func (r NotInRule) Code(code string) Rule {
	r.err = r.err.SetMessage(code)
	return r
}

func (r NotInRule) Msg(msg string) Rule {
	r.err = r.err.SetMessage(msg)
	return r
}

// Validate checks if the given value is valid or not.
func (r NotInRule) Validate(lang string, value interface{}) error {
	r.err = r.err.SetLang(lang)
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return nil
	}

	for _, e := range r.elements {
		if e == value {
			return r.err
		}
	}
	return nil
}
