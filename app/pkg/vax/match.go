// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vax

import (
	"regexp"
)

// ErrMatchInvalid is the error that returns in case of invalid format.
var ErrMatchInvalid = NewError("validate.match", "must match regular expression %v")

// Match returns a validation rule that checks if a value matches the specified regular expression.
// This rule should only be used for validating strings and byte slices, or a validation error will be reported.
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
func Match(re *regexp.Regexp) MatchRule {
	return MatchRule{
		re:  re,
		err: ErrMatchInvalid,
	}
}

// MatchRule is a validation rule that checks if a value matches the specified regular expression.
type MatchRule struct {
	re  *regexp.Regexp
	err Error
}

func (r MatchRule) Code(code string) Rule {
	r.err = r.err.SetCode(code)
	return r
}

func (r MatchRule) Msg(msg string) Rule {
	r.err = r.err.SetMessage(msg)
	return r
}

// Validate checks if the given value is valid or not.
func (r MatchRule) Validate(lang string, value interface{}) error {
	r.err = r.err.SetLang(lang)
	value, isNil := Indirect(value)
	if isNil {
		return nil
	}

	isString, str, isBytes, bs := StringOrBytes(value)
	if isString && (str == "" || r.re.MatchString(str)) {
		return nil
	} else if isBytes && (len(bs) == 0 || r.re.Match(bs)) {
		return nil
	}
	return r.err
}
