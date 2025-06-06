// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: configure/business/v1/configure_business_value.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ListBusinessValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListBusinessValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListBusinessValueRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListBusinessValueRequestMultiError, or nil if none found.
func (m *ListBusinessValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListBusinessValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBusinessId() <= 0 {
		err := ListBusinessValueRequestValidationError{
			field:  "BusinessId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListBusinessValueRequestMultiError(errors)
	}

	return nil
}

// ListBusinessValueRequestMultiError is an error wrapping multiple validation
// errors returned by ListBusinessValueRequest.ValidateAll() if the designated
// constraints aren't met.
type ListBusinessValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListBusinessValueRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListBusinessValueRequestMultiError) AllErrors() []error { return m }

// ListBusinessValueRequestValidationError is the validation error returned by
// ListBusinessValueRequest.Validate if the designated constraints aren't met.
type ListBusinessValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBusinessValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBusinessValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBusinessValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBusinessValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBusinessValueRequestValidationError) ErrorName() string {
	return "ListBusinessValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListBusinessValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBusinessValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBusinessValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBusinessValueRequestValidationError{}

// Validate checks the field values on ListBusinessValueReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListBusinessValueReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListBusinessValueReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListBusinessValueReplyMultiError, or nil if none found.
func (m *ListBusinessValueReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListBusinessValueReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListBusinessValueReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListBusinessValueReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListBusinessValueReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListBusinessValueReplyMultiError(errors)
	}

	return nil
}

// ListBusinessValueReplyMultiError is an error wrapping multiple validation
// errors returned by ListBusinessValueReply.ValidateAll() if the designated
// constraints aren't met.
type ListBusinessValueReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListBusinessValueReplyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListBusinessValueReplyMultiError) AllErrors() []error { return m }

// ListBusinessValueReplyValidationError is the validation error returned by
// ListBusinessValueReply.Validate if the designated constraints aren't met.
type ListBusinessValueReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBusinessValueReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBusinessValueReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBusinessValueReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBusinessValueReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBusinessValueReplyValidationError) ErrorName() string {
	return "ListBusinessValueReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListBusinessValueReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBusinessValueReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBusinessValueReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBusinessValueReplyValidationError{}

// Validate checks the field values on UpdateBusinessValueRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateBusinessValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateBusinessValueRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateBusinessValueRequestMultiError, or nil if none found.
func (m *UpdateBusinessValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateBusinessValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetList()) < 1 {
		err := UpdateBusinessValueRequestValidationError{
			field:  "List",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpdateBusinessValueRequestValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpdateBusinessValueRequestValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateBusinessValueRequestValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetBusinessId() <= 0 {
		err := UpdateBusinessValueRequestValidationError{
			field:  "BusinessId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateBusinessValueRequestMultiError(errors)
	}

	return nil
}

// UpdateBusinessValueRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateBusinessValueRequest.ValidateAll() if
// the designated constraints aren't met.
type UpdateBusinessValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateBusinessValueRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateBusinessValueRequestMultiError) AllErrors() []error { return m }

// UpdateBusinessValueRequestValidationError is the validation error returned
// by UpdateBusinessValueRequest.Validate if the designated constraints aren't met.
type UpdateBusinessValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateBusinessValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateBusinessValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateBusinessValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateBusinessValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateBusinessValueRequestValidationError) ErrorName() string {
	return "UpdateBusinessValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateBusinessValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateBusinessValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateBusinessValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateBusinessValueRequestValidationError{}

// Validate checks the field values on UpdateBusinessValueReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateBusinessValueReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateBusinessValueReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateBusinessValueReplyMultiError, or nil if none found.
func (m *UpdateBusinessValueReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateBusinessValueReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateBusinessValueReplyMultiError(errors)
	}

	return nil
}

// UpdateBusinessValueReplyMultiError is an error wrapping multiple validation
// errors returned by UpdateBusinessValueReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateBusinessValueReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateBusinessValueReplyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateBusinessValueReplyMultiError) AllErrors() []error { return m }

// UpdateBusinessValueReplyValidationError is the validation error returned by
// UpdateBusinessValueReply.Validate if the designated constraints aren't met.
type UpdateBusinessValueReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateBusinessValueReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateBusinessValueReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateBusinessValueReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateBusinessValueReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateBusinessValueReplyValidationError) ErrorName() string {
	return "UpdateBusinessValueReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateBusinessValueReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateBusinessValueReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateBusinessValueReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateBusinessValueReplyValidationError{}

// Validate checks the field values on ListBusinessValueReply_BusinessValue
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *ListBusinessValueReply_BusinessValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListBusinessValueReply_BusinessValue
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// ListBusinessValueReply_BusinessValueMultiError, or nil if none found.
func (m *ListBusinessValueReply_BusinessValue) ValidateAll() error {
	return m.validate(true)
}

func (m *ListBusinessValueReply_BusinessValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for EnvId

	// no validation rules for BusinessId

	// no validation rules for Value

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return ListBusinessValueReply_BusinessValueMultiError(errors)
	}

	return nil
}

// ListBusinessValueReply_BusinessValueMultiError is an error wrapping multiple
// validation errors returned by
// ListBusinessValueReply_BusinessValue.ValidateAll() if the designated
// constraints aren't met.
type ListBusinessValueReply_BusinessValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListBusinessValueReply_BusinessValueMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListBusinessValueReply_BusinessValueMultiError) AllErrors() []error { return m }

// ListBusinessValueReply_BusinessValueValidationError is the validation error
// returned by ListBusinessValueReply_BusinessValue.Validate if the designated
// constraints aren't met.
type ListBusinessValueReply_BusinessValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBusinessValueReply_BusinessValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBusinessValueReply_BusinessValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBusinessValueReply_BusinessValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBusinessValueReply_BusinessValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBusinessValueReply_BusinessValueValidationError) ErrorName() string {
	return "ListBusinessValueReply_BusinessValueValidationError"
}

// Error satisfies the builtin error interface
func (e ListBusinessValueReply_BusinessValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBusinessValueReply_BusinessValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBusinessValueReply_BusinessValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBusinessValueReply_BusinessValueValidationError{}

// Validate checks the field values on UpdateBusinessValueRequest_Value with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *UpdateBusinessValueRequest_Value) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateBusinessValueRequest_Value with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// UpdateBusinessValueRequest_ValueMultiError, or nil if none found.
func (m *UpdateBusinessValueRequest_Value) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateBusinessValueRequest_Value) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetEnvId() <= 0 {
		err := UpdateBusinessValueRequest_ValueValidationError{
			field:  "EnvId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetValue()) < 1 {
		err := UpdateBusinessValueRequest_ValueValidationError{
			field:  "Value",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateBusinessValueRequest_ValueMultiError(errors)
	}

	return nil
}

// UpdateBusinessValueRequest_ValueMultiError is an error wrapping multiple
// validation errors returned by
// UpdateBusinessValueRequest_Value.ValidateAll() if the designated
// constraints aren't met.
type UpdateBusinessValueRequest_ValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateBusinessValueRequest_ValueMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateBusinessValueRequest_ValueMultiError) AllErrors() []error { return m }

// UpdateBusinessValueRequest_ValueValidationError is the validation error
// returned by UpdateBusinessValueRequest_Value.Validate if the designated
// constraints aren't met.
type UpdateBusinessValueRequest_ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateBusinessValueRequest_ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateBusinessValueRequest_ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateBusinessValueRequest_ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateBusinessValueRequest_ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateBusinessValueRequest_ValueValidationError) ErrorName() string {
	return "UpdateBusinessValueRequest_ValueValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateBusinessValueRequest_ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateBusinessValueRequest_Value.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateBusinessValueRequest_ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateBusinessValueRequest_ValueValidationError{}
