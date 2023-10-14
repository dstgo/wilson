// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vax

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	// Error interface represents an validation error
	Error interface {
		Error() string
		SetLang(string) Error
		Lang() string
		SetCode(string) Error
		Code() string
		Message() string
		SetMessage(string) Error
		Params() []any
		AddParams(args ...any) Error
	}

	// ErrorObject is the default validation error
	// that implements the Error interface.
	ErrorObject struct {
		lang    string
		code    string
		message string
		params  []any
	}

	Err struct {
		Key string
		Err error
	}

	// Errors represents the validation errors that are indexed by struct field names, map or slice keys.
	// values are Error or Errors (for map, slice and array error value is Errors).
	Errors []Err

	// InternalError represents an error that should NOT be treated as a validation error.
	InternalError interface {
		error
		InternalError() error
	}

	internalError struct {
		error
	}
)

func (e ErrorObject) Lang() string {
	return e.lang
}

func (e ErrorObject) SetLang(s string) Error {
	e.lang = s
	return e
}

// SetCode set the error's translation code.
func (e ErrorObject) SetCode(code string) Error {
	e.code = code
	return e
}

// Code get the error's translation code.
func (e ErrorObject) Code() string {
	return e.code
}

// AddParams set the error's params.
func (e ErrorObject) AddParams(args ...any) Error {
	e.params = append(e.params, args...)
	return e
}

// Params returns the error's params.
func (e ErrorObject) Params() []any {
	return e.params
}

// SetMessage set the error's message.
func (e ErrorObject) SetMessage(message string) Error {
	e.message = message
	return e
}

// Message return the error's message.
func (e ErrorObject) Message() string {
	return e.message
}

// Error returns the error message.
func (e ErrorObject) Error() string {
	if e.code == "" || translator == nil {
		return fmt.Sprintf(e.message, e.params...)
	}
	return translator.Get(e.lang, e.code, e.params...)
}

// Error returns the error string of Errors.
func (es Errors) Error() string {
	if len(es) == 0 {
		return ""
	}
	var s strings.Builder
	for i, e := range es {
		if i > 0 {
			s.WriteString("; ")
		}
		fmt.Fprintf(&s, "%v: %v", e.Key, e.Err.Error())
	}
	return s.String()
}

// MarshalJSON converts the Errors into a valid JSON.
func (es Errors) MarshalJSON() ([]byte, error) {
	errs := map[string]interface{}{}
	for _, e := range es {
		if ms, ok := e.Err.(json.Marshaler); ok {
			errs[e.Key] = ms
		} else {
			errs[e.Key] = e.Err.Error()
		}
	}
	return json.Marshal(errs)
}

// NewError create new validation error.
func NewError(code, message string) Error {
	return ErrorObject{
		code:    code,
		message: message,
	}
}

// Assert that our ErrorObject implements the Error interface.
var _ Error = ErrorObject{}

// NewInternalError wraps a given error into an InternalError.
func NewInternalError(err error) InternalError {
	return internalError{error: err}
}

// InternalError returns the actual error that it wraps around.
func (e internalError) InternalError() error {
	return e.error
}
