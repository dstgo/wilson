// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package vax validation provides configurable and extensible rules for validating data of various types.
package vax

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
)

type (
	// Validatable is the interface indicating the type implementing it supports data validation.
	Validatable interface {
		// Validate validates the data and returns an error if validation fails.
		Validate(lang string) error
	}

	// ValidatableWithContext is the interface indicating the type implementing it supports context-aware data validation.
	ValidatableWithContext interface {
		// ValidateWithContext validates the data with the given context and returns an error if validation fails.
		ValidateWithContext(ctx context.Context, lang string) error
	}

	// Rule represents a validation rule.
	Rule interface {
		// Validate validates a value and returns a value if validation fails.
		Validate(lang string, value interface{}) error
		Code(code string) Rule
		Msg(msg string) Rule
	}

	// RuleWithContext represents a context-aware validation rule.
	RuleWithContext interface {
		// ValidateWithContext validates a value and returns a value if validation fails.
		ValidateWithContext(ctx context.Context, lang string, value interface{}) error
	}

	// RuleFunc represents a validator function.
	// You may wrap it as a Rule by calling By().
	RuleFunc func(lang string, value interface{}) error

	// RuleWithContextFunc represents a validator function that is context-aware.
	// You may wrap it as a Rule by calling WithContext().
	RuleWithContextFunc func(ctx context.Context, lang string, value interface{}) error

	Translator interface {
		// Default return the default language used of the translator.
		Default() string
		Get(lang string, key string, params ...any) string
	}
)

var (
	// LabelTag is the struct tag name used to customize the error field name for a struct field.
	LabelTag = "label"

	// translator is the default translator.
	translator Translator

	// Skip is a special validation rule that indicates all rules following it should be skipped.
	Skip = skipRule{skip: true}

	validatableType            = reflect.TypeOf((*Validatable)(nil)).Elem()
	validatableWithContextType = reflect.TypeOf((*ValidatableWithContext)(nil)).Elem()
)

func SetTranslator(t Translator) {
	translator = t
}

// Validate validates the given value and returns the validation error, if any.
//
// Validate performs validation using the following steps:
//  1. For each rule, call its `Validate()` to validate the value. Return if any error is found.
//  2. If the value being validated implements `Validatable`, call the value's `Validate()`.
//     Return with the validation result.
//  3. If the value being validated is a map/slice/array, and the element type implements `Validatable`,
//     for each element call the element value's `Validate()`. Return with the validation result.
func Validate(lang string, value interface{}, rules ...Rule) error {
	for _, rule := range rules {
		if s, ok := rule.(skipRule); ok && s.skip {
			return nil
		}
		if err := rule.Validate(lang, value); err != nil {
			return err
		}
	}

	rv := reflect.ValueOf(value)
	if (rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface) && rv.IsNil() {
		return nil
	}

	if v, ok := value.(Validatable); ok {
		return v.Validate(lang)
	}

	switch rv.Kind() {
	case reflect.Map:
		if rv.Type().Elem().Implements(validatableType) {
			return validateMap(lang, rv)
		}
	case reflect.Slice, reflect.Array:
		if rv.Type().Elem().Implements(validatableType) {
			return validateSlice(lang, rv)
		}
	case reflect.Ptr, reflect.Interface:
		return Validate(lang, rv.Elem().Interface())
	}

	return nil
}

// ValidateWithContext validates the given value with the given context and returns the validation error, if any.
//
// ValidateWithContext performs validation using the following steps:
//  1. For each rule, call its `ValidateWithContext()` to validate the value if the rule implements `RuleWithContext`.
//     Otherwise call `Validate()` of the rule. Return if any error is found.
//  2. If the value being validated implements `ValidatableWithContext`, call the value's `ValidateWithContext()`
//     and return with the validation result.
//  3. If the value being validated implements `Validatable`, call the value's `Validate()`
//     and return with the validation result.
//  4. If the value being validated is a map/slice/array, and the element type implements `ValidatableWithContext`,
//     for each element call the element value's `ValidateWithContext()`. Return with the validation result.
//  5. If the value being validated is a map/slice/array, and the element type implements `Validatable`,
//     for each element call the element value's `Validate()`. Return with the validation result.
func ValidateWithContext(ctx context.Context, lang string, value interface{}, rules ...Rule) error {
	for _, rule := range rules {
		if s, ok := rule.(skipRule); ok && s.skip {
			return nil
		}
		if rc, ok := rule.(RuleWithContext); ok {
			if err := rc.ValidateWithContext(ctx, lang, value); err != nil {
				return err
			}
		} else if err := rule.Validate(lang, value); err != nil {
			return err
		}
	}

	rv := reflect.ValueOf(value)
	if (rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface) && rv.IsNil() {
		return nil
	}

	if v, ok := value.(ValidatableWithContext); ok {
		return v.ValidateWithContext(ctx, lang)
	}

	if v, ok := value.(Validatable); ok {
		return v.Validate(lang)
	}

	switch rv.Kind() {
	case reflect.Map:
		if rv.Type().Elem().Implements(validatableWithContextType) {
			return validateMapWithContext(ctx, lang, rv)
		}
		if rv.Type().Elem().Implements(validatableType) {
			return validateMap(lang, rv)
		}
	case reflect.Slice, reflect.Array:
		if rv.Type().Elem().Implements(validatableWithContextType) {
			return validateSliceWithContext(ctx, lang, rv)
		}
		if rv.Type().Elem().Implements(validatableType) {
			return validateSlice(lang, rv)
		}
	case reflect.Ptr, reflect.Interface:
		return ValidateWithContext(ctx, lang, rv.Elem().Interface())
	}

	return nil
}

// validateMap validates a map of validatable elements
func validateMap(lang string, rv reflect.Value) error {
	errs := make(Errors, 0)
	for _, key := range rv.MapKeys() {
		if mv := rv.MapIndex(key).Interface(); mv != nil {
			if err := mv.(Validatable).Validate(lang); err != nil {
				errs = append(errs, Err{fmt.Sprintf("%v", key.Interface()), err})
			}
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// validateMapWithContext validates a map of validatable elements with the given context.
func validateMapWithContext(ctx context.Context, lang string, rv reflect.Value) error {
	errs := make(Errors, 0)
	for _, key := range rv.MapKeys() {
		if mv := rv.MapIndex(key).Interface(); mv != nil {
			if err := mv.(ValidatableWithContext).ValidateWithContext(ctx, lang); err != nil {
				errs = append(errs, Err{fmt.Sprintf("%v", key.Interface()), err})
			}
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// validateSlice validates a slice/array of validatable elements
func validateSlice(lang string, rv reflect.Value) error {
	errs := make(Errors, 0)
	l := rv.Len()
	for i := 0; i < l; i++ {
		if ev := rv.Index(i).Interface(); ev != nil {
			if err := ev.(Validatable).Validate(lang); err != nil {
				errs = append(errs, Err{strconv.Itoa(i), err})
			}
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// validateSliceWithContext validates a slice/array of validatable elements with the given context.
func validateSliceWithContext(ctx context.Context, lang string, rv reflect.Value) error {
	errs := make(Errors, 0)
	l := rv.Len()
	for i := 0; i < l; i++ {
		if ev := rv.Index(i).Interface(); ev != nil {
			if err := ev.(ValidatableWithContext).ValidateWithContext(ctx, lang); err != nil {
				errs = append(errs, Err{strconv.Itoa(i), err})
			}
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

type skipRule struct {
	skip bool
}

func (r skipRule) Code(code string) Rule {
	return r
}

func (r skipRule) Msg(msg string) Rule {
	return r
}

func (r skipRule) Validate(string, interface{}) error {
	return nil
}

// When determines if all rules following it should be skipped.
func (r skipRule) When(condition bool) skipRule {
	r.skip = condition
	return r
}

type inlineRule struct {
	f  RuleFunc
	fc RuleWithContextFunc
}

func (r *inlineRule) Code(code string) Rule {
	return r
}

func (r *inlineRule) Msg(msg string) Rule {
	return r
}

func (r *inlineRule) Validate(lang string, value interface{}) error {
	if r.f == nil {
		return r.fc(context.Background(), lang, value)
	}
	return r.f(lang, value)
}

func (r *inlineRule) ValidateWithContext(ctx context.Context, lang string, value interface{}) error {
	if r.fc == nil {
		return r.f(lang, value)
	}
	return r.fc(ctx, lang, value)
}

// By wraps a RuleFunc into a Rule.
func By(f RuleFunc) Rule {
	return &inlineRule{f: f}
}

// WithContext wraps a RuleWithContextFunc into a context-aware Rule.
func WithContext(f RuleWithContextFunc) Rule {
	return &inlineRule{fc: f}
}
