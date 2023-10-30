// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vax

var (
	// ErrLengthTooLong is the error that returns in case of too long length.
	ErrLengthTooLong = NewError("validate.length.max", "the length must be no more than %v")
	// ErrLengthTooShort is the error that returns in case of too short length.
	ErrLengthTooShort = NewError("validate.length.min", "the length must be no less than %v")
	// ErrLengthInvalid is the error that returns in case of an invalid length.
	ErrLengthInvalid = NewError("validate.length.eq", "the length must be exactly %v")
	// ErrLengthOutOfRange is the error that returns in case of out of range length.
	ErrLengthOutOfRange = NewError("validate.length.range", "the length must be between %v and %v")
)

func RangeLenRune(min, max int) RangeLengthRule {
	return RangeLength(min, max, true)
}

func MaxLenRune(max int) MaxLengthRule {
	return MaxLength(max, true)
}

func MinLenRune(min int) MinLengthRule {
	return MinLength(min, true)
}

func RangeLength(min, max int, rune bool) RangeLengthRule {
	if min < 0 || max < 0 {
		panic("min and max length must be non-negative")
	}
	if min > max {
		panic("min length must be greater than max length")
	}
	return RangeLengthRule{min: min, max: max, err: ErrLengthOutOfRange.AddParams(min, max), rune: rune}
}

// RangeLengthRule is a validation rule that checks if a value's length is within the specified range.
type RangeLengthRule struct {
	err Error

	min, max int
	rune     bool
}

func (r RangeLengthRule) Code(code string) Rule {
	r.err = r.err.SetCode(code)
	return r
}

func (r RangeLengthRule) Msg(msg string) Rule {
	r.err = r.err.SetMessage(msg)
	return r
}

// Validate checks if the given value is valid or not.
func (r RangeLengthRule) Validate(lang string, value interface{}) error {
	r.err = r.err.SetLang(lang)
	l, err := Len(value, r.rune)
	if err != nil {
		return err
	}
	if !(r.min <= l && l <= r.max) {
		return r.err
	}
	return nil
}

func MaxLength(max int, rune bool) MaxLengthRule {
	if max < 0 {
		panic("max length must be non-negative")
	}
	return MaxLengthRule{max: max, rune: rune, err: ErrLengthTooLong.AddParams(max)}
}

func MinLength(min int, rune bool) MinLengthRule {
	if min < 0 {
		panic("min length must be non-negative")
	}
	return MinLengthRule{min: min, rune: rune, err: ErrLengthTooShort.AddParams(min)}
}

func EqLength(l int, rune bool) EqLengthRule {
	if l < 0 {
		panic("eq length must be non-negative")
	}
	return EqLengthRule{eq: l, rune: rune, err: ErrLengthInvalid.AddParams(l)}
}

type MaxLengthRule struct {
	max  int
	err  Error
	rune bool
}

func (r MaxLengthRule) Validate(lang string, value interface{}) error {
	r.err = r.err.SetLang(lang)
	l, err := Len(value, false)
	if err != nil {
		return err
	}
	if l > r.max {
		return r.err
	}
	return nil
}

func (r MaxLengthRule) Code(code string) Rule {
	r.err.SetCode(code)
	return r
}

func (r MaxLengthRule) Msg(msg string) Rule {
	r.err.SetMessage(msg)
	return r
}

type MinLengthRule struct {
	min  int
	err  Error
	rune bool
}

func (r MinLengthRule) Validate(lang string, value interface{}) error {
	r.err = r.err.SetLang(lang)
	l, err := Len(value, false)
	if err != nil {
		return err
	}
	if l < r.min {
		return r.err
	}
	return nil
}

func (r MinLengthRule) Code(code string) Rule {
	r.err.SetCode(code)
	return r
}

func (r MinLengthRule) Msg(msg string) Rule {
	r.err.SetMessage(msg)
	return r
}

type EqLengthRule struct {
	eq   int
	err  Error
	rune bool
}

func (r EqLengthRule) Validate(lang string, value interface{}) error {
	r.err = r.err.SetLang(lang)
	l, err := Len(value, false)
	if err != nil {
		return err
	}
	if l != r.eq {
		return r.err
	}
	return nil
}

func (r EqLengthRule) Code(code string) Rule {
	r.err.SetCode(code)
	return r
}

func (r EqLengthRule) Msg(msg string) Rule {
	r.err.SetMessage(msg)
	return r
}
