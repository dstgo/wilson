// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vax

// ErrNotNilRequired is the error that returns when a value is Nil.
var ErrNotNilRequired = NewError("validate.notnil", "can not be nil")

// NotNil is a validation rule that checks if a value is not nil.
// NotNil only handles types including interface, pointer, slice, and map.
// All other types are considered valid.
var NotNil = notNilRule{}

type notNilRule struct {
	err Error
}

func (r notNilRule) Code(code string) Rule {
	r.err = r.err.SetCode(code)
	return r
}

func (r notNilRule) Msg(msg string) Rule {
	if r.err == nil {
		r.err = ErrNotNilRequired
	}
	r.err = r.err.SetMessage(msg)
	return r
}

// Validate checks if the given value is valid or not.
func (r notNilRule) Validate(lang string, value interface{}) error {
	_, isNil := Indirect(value)
	if isNil {
		if r.err != nil {
			return r.err
		}
		return ErrNotNilRequired.SetLang(lang)
	}
	return nil
}

// Error sets the error message for the rule.
func (r notNilRule) Error(message string) notNilRule {
	if r.err == nil {
		r.err = ErrNotNilRequired
	}
	r.err = r.err.SetMessage(message)
	return r
}
