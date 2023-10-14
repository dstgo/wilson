// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vax

type stringValidator func(string) bool

// StringRule is a rule that checks a string variable using a specified stringValidator.
type StringRule struct {
	validate stringValidator
	err      Error
}

func (r StringRule) Code(code string) Rule {
	r.err = r.err.SetCode(code)
	return r
}

func (r StringRule) Msg(msg string) Rule {
	r.err = r.err.SetMessage(msg)
	return r
}

func StringCode(validator stringValidator, code string) StringRule {
	return StringRule{
		validate: validator,
		err:      NewError(code, ""),
	}
}

func StringMsg(validator stringValidator, msg string) StringRule {
	return StringRule{
		validate: validator,
		err:      NewError("", msg),
	}
}

func String(validator stringValidator, err Error) StringRule {
	return StringRule{
		validate: validator,
		err:      err,
	}
}

// Error sets the error message for the rule.
func (r StringRule) Error(message string) StringRule {
	r.err = r.err.SetMessage(message)
	return r
}

// ErrorObject sets the error struct for the rule.
func (r StringRule) ErrorObject(err Error) StringRule {
	r.err = err
	return r
}

// Validate checks if the given value is valid or not.
func (r StringRule) Validate(lang string, value interface{}) error {
	r.err = r.err.SetLang(lang)
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return nil
	}

	str, err := EnsureString(value)
	if err != nil {
		return err
	}

	if r.validate(str) {
		return nil
	}

	return r.err
}
