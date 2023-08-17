// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vax

import (
	"fmt"
	"reflect"
	"time"
)

var (
	// ErrMinGreaterEqualThanRequired is the error that returns when a value is less than a specified threshold.
	ErrMinGreaterEqualThanRequired = NewError("validate.gte", "must be greater than or equal to %v")
	// ErrMaxLessEqualThanRequired is the error that returns when a value is greater than a specified threshold.
	ErrMaxLessEqualThanRequired = NewError("validate.lte", "must be less than or equal to %v")
	// ErrMinGreaterThanRequired is the error that returns when a value is less than or equal to a specified threshold.
	ErrMinGreaterThanRequired = NewError("validate.gt", "must be greater than %v")
	// ErrMaxLessThanRequired is the error that returns when a value is greater than or equal to a specified threshold.
	ErrMaxLessThanRequired = NewError("validate.lt", "must be less than %v")
	ErrEqualRequired       = NewError("validate.eq", "must be equal to %v")
)

// CompareRule is a validation rule that checks if a value satisfies the specified threshold requirement.
type CompareRule struct {
	threshold interface{}
	operator  int
	err       Error
}

func (r CompareRule) Code(code string) Rule {
	r.err = r.err.SetCode(code)
	return r
}

func (r CompareRule) Msg(msg string) Rule {
	r.err = r.err.SetMessage(msg)
	return r
}

const (
	greaterThan = iota
	greaterEqualThan
	lessThan
	lessEqualThan
	equal
)

func Gte(min interface{}) CompareRule {
	return CompareRule{
		threshold: min,
		operator:  greaterEqualThan,
		err:       ErrMinGreaterEqualThanRequired,
	}
}

func Gt(min interface{}) CompareRule {
	return CompareRule{
		threshold: min,
		operator:  greaterThan,
		err:       ErrMinGreaterThanRequired,
	}
}

func Lte(max interface{}) CompareRule {
	return CompareRule{
		threshold: max,
		operator:  lessEqualThan,
		err:       ErrMaxLessEqualThanRequired,
	}
}

func Lt(max interface{}) CompareRule {
	return CompareRule{
		threshold: max,
		operator:  lessThan,
		err:       ErrMaxLessThanRequired,
	}
}

func Eq(num any) CompareRule {
	return CompareRule{
		threshold: num,
		operator:  equal,
		err:       ErrEqualRequired,
	}
}

// Validate checks if the given value is valid or not.
func (r CompareRule) Validate(lang string, value interface{}) error {
	r.err = r.err.SetLang(lang)
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return nil
	}

	rv := reflect.ValueOf(r.threshold)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := ToInt(value)
		if err != nil {
			return err
		}
		if r.compareInt(rv.Int(), v) {
			return nil
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v, err := ToUint(value)
		if err != nil {
			return err
		}
		if r.compareUint(rv.Uint(), v) {
			return nil
		}

	case reflect.Float32, reflect.Float64:
		v, err := ToFloat(value)
		if err != nil {
			return err
		}
		if r.compareFloat(rv.Float(), v) {
			return nil
		}

	case reflect.Struct:
		t, ok := r.threshold.(time.Time)
		if !ok {
			return fmt.Errorf("type not supported: %v", rv.Type())
		}
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("cannot convert %v to time.Time", reflect.TypeOf(value))
		}
		if v.IsZero() || r.compareTime(t, v) {
			return nil
		}

	default:
		return fmt.Errorf("type not supported: %v", rv.Type())
	}

	return r.err.AddParams(r.threshold)
}

func (r CompareRule) compareInt(threshold, value int64) bool {
	switch r.operator {
	case greaterThan:
		return value > threshold
	case greaterEqualThan:
		return value >= threshold
	case lessThan:
		return value < threshold
	case equal:
		return value == threshold
	default:
		return value <= threshold
	}
}

func (r CompareRule) compareUint(threshold, value uint64) bool {
	switch r.operator {
	case greaterThan:
		return value > threshold
	case greaterEqualThan:
		return value >= threshold
	case lessThan:
		return value < threshold
	case equal:
		return value == threshold
	default:
		return value <= threshold
	}
}

func (r CompareRule) compareFloat(threshold, value float64) bool {
	switch r.operator {
	case greaterThan:
		return value > threshold
	case greaterEqualThan:
		return value >= threshold
	case lessThan:
		return value < threshold
	case equal:
		return value == threshold
	default:
		return value <= threshold
	}
}

func (r CompareRule) compareTime(threshold, value time.Time) bool {
	switch r.operator {
	case greaterThan:
		return value.After(threshold)
	case greaterEqualThan:
		return value.After(threshold) || value.Equal(threshold)
	case lessThan:
		return value.Before(threshold)
	case equal:
		return value.Equal(threshold)
	default:
		return value.Before(threshold) || value.Equal(threshold)
	}
}
