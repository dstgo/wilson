// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: manager/dictionary/v1/manager_dictionary.proto

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

// Validate checks the field values on ListDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListDictionaryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListDictionaryRequestMultiError, or nil if none found.
func (m *ListDictionaryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListDictionaryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() < 1 {
		err := ListDictionaryRequestValidationError{
			field:  "Page",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val < 1 || val > 50 {
		err := ListDictionaryRequestValidationError{
			field:  "PageSize",
			reason: "value must be inside range [1, 50]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Keyword != nil {
		// no validation rules for Keyword
	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if len(errors) > 0 {
		return ListDictionaryRequestMultiError(errors)
	}

	return nil
}

// ListDictionaryRequestMultiError is an error wrapping multiple validation
// errors returned by ListDictionaryRequest.ValidateAll() if the designated
// constraints aren't met.
type ListDictionaryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListDictionaryRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListDictionaryRequestMultiError) AllErrors() []error { return m }

// ListDictionaryRequestValidationError is the validation error returned by
// ListDictionaryRequest.Validate if the designated constraints aren't met.
type ListDictionaryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDictionaryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDictionaryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDictionaryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDictionaryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDictionaryRequestValidationError) ErrorName() string {
	return "ListDictionaryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListDictionaryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDictionaryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDictionaryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDictionaryRequestValidationError{}

// Validate checks the field values on ListDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListDictionaryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListDictionaryReplyMultiError, or nil if none found.
func (m *ListDictionaryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListDictionaryReply) validate(all bool) error {
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
					errors = append(errors, ListDictionaryReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListDictionaryReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListDictionaryReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListDictionaryReplyMultiError(errors)
	}

	return nil
}

// ListDictionaryReplyMultiError is an error wrapping multiple validation
// errors returned by ListDictionaryReply.ValidateAll() if the designated
// constraints aren't met.
type ListDictionaryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListDictionaryReplyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListDictionaryReplyMultiError) AllErrors() []error { return m }

// ListDictionaryReplyValidationError is the validation error returned by
// ListDictionaryReply.Validate if the designated constraints aren't met.
type ListDictionaryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDictionaryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDictionaryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDictionaryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDictionaryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDictionaryReplyValidationError) ErrorName() string {
	return "ListDictionaryReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListDictionaryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDictionaryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDictionaryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDictionaryReplyValidationError{}

// Validate checks the field values on CreateDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateDictionaryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateDictionaryRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateDictionaryRequestMultiError, or nil if none found.
func (m *CreateDictionaryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateDictionaryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKeyword()) < 1 {
		err := CreateDictionaryRequestValidationError{
			field:  "Keyword",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := CreateDictionaryRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Description != nil {
		// no validation rules for Description
	}

	if len(errors) > 0 {
		return CreateDictionaryRequestMultiError(errors)
	}

	return nil
}

// CreateDictionaryRequestMultiError is an error wrapping multiple validation
// errors returned by CreateDictionaryRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateDictionaryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateDictionaryRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateDictionaryRequestMultiError) AllErrors() []error { return m }

// CreateDictionaryRequestValidationError is the validation error returned by
// CreateDictionaryRequest.Validate if the designated constraints aren't met.
type CreateDictionaryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDictionaryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDictionaryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDictionaryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDictionaryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDictionaryRequestValidationError) ErrorName() string {
	return "CreateDictionaryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateDictionaryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDictionaryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDictionaryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDictionaryRequestValidationError{}

// Validate checks the field values on CreateDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateDictionaryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateDictionaryReplyMultiError, or nil if none found.
func (m *CreateDictionaryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateDictionaryReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateDictionaryReplyMultiError(errors)
	}

	return nil
}

// CreateDictionaryReplyMultiError is an error wrapping multiple validation
// errors returned by CreateDictionaryReply.ValidateAll() if the designated
// constraints aren't met.
type CreateDictionaryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateDictionaryReplyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateDictionaryReplyMultiError) AllErrors() []error { return m }

// CreateDictionaryReplyValidationError is the validation error returned by
// CreateDictionaryReply.Validate if the designated constraints aren't met.
type CreateDictionaryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDictionaryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDictionaryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDictionaryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDictionaryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDictionaryReplyValidationError) ErrorName() string {
	return "CreateDictionaryReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateDictionaryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDictionaryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDictionaryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDictionaryReplyValidationError{}

// Validate checks the field values on UpdateDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateDictionaryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateDictionaryRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateDictionaryRequestMultiError, or nil if none found.
func (m *UpdateDictionaryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateDictionaryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() < 1 {
		err := UpdateDictionaryRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Keyword

	// no validation rules for Name

	if m.Description != nil {
		// no validation rules for Description
	}

	if len(errors) > 0 {
		return UpdateDictionaryRequestMultiError(errors)
	}

	return nil
}

// UpdateDictionaryRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateDictionaryRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateDictionaryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateDictionaryRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateDictionaryRequestMultiError) AllErrors() []error { return m }

// UpdateDictionaryRequestValidationError is the validation error returned by
// UpdateDictionaryRequest.Validate if the designated constraints aren't met.
type UpdateDictionaryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDictionaryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDictionaryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDictionaryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDictionaryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDictionaryRequestValidationError) ErrorName() string {
	return "UpdateDictionaryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDictionaryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDictionaryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDictionaryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDictionaryRequestValidationError{}

// Validate checks the field values on UpdateDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateDictionaryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateDictionaryReplyMultiError, or nil if none found.
func (m *UpdateDictionaryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateDictionaryReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateDictionaryReplyMultiError(errors)
	}

	return nil
}

// UpdateDictionaryReplyMultiError is an error wrapping multiple validation
// errors returned by UpdateDictionaryReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateDictionaryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateDictionaryReplyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateDictionaryReplyMultiError) AllErrors() []error { return m }

// UpdateDictionaryReplyValidationError is the validation error returned by
// UpdateDictionaryReply.Validate if the designated constraints aren't met.
type UpdateDictionaryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDictionaryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDictionaryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDictionaryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDictionaryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDictionaryReplyValidationError) ErrorName() string {
	return "UpdateDictionaryReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDictionaryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDictionaryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDictionaryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDictionaryReplyValidationError{}

// Validate checks the field values on DeleteDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteDictionaryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteDictionaryRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteDictionaryRequestMultiError, or nil if none found.
func (m *DeleteDictionaryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteDictionaryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() < 1 {
		err := DeleteDictionaryRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteDictionaryRequestMultiError(errors)
	}

	return nil
}

// DeleteDictionaryRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteDictionaryRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteDictionaryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteDictionaryRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteDictionaryRequestMultiError) AllErrors() []error { return m }

// DeleteDictionaryRequestValidationError is the validation error returned by
// DeleteDictionaryRequest.Validate if the designated constraints aren't met.
type DeleteDictionaryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteDictionaryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteDictionaryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteDictionaryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteDictionaryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteDictionaryRequestValidationError) ErrorName() string {
	return "DeleteDictionaryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteDictionaryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteDictionaryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteDictionaryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteDictionaryRequestValidationError{}

// Validate checks the field values on DeleteDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteDictionaryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteDictionaryReplyMultiError, or nil if none found.
func (m *DeleteDictionaryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteDictionaryReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteDictionaryReplyMultiError(errors)
	}

	return nil
}

// DeleteDictionaryReplyMultiError is an error wrapping multiple validation
// errors returned by DeleteDictionaryReply.ValidateAll() if the designated
// constraints aren't met.
type DeleteDictionaryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteDictionaryReplyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteDictionaryReplyMultiError) AllErrors() []error { return m }

// DeleteDictionaryReplyValidationError is the validation error returned by
// DeleteDictionaryReply.Validate if the designated constraints aren't met.
type DeleteDictionaryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteDictionaryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteDictionaryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteDictionaryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteDictionaryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteDictionaryReplyValidationError) ErrorName() string {
	return "DeleteDictionaryReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteDictionaryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteDictionaryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteDictionaryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteDictionaryReplyValidationError{}

// Validate checks the field values on GetDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDictionaryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDictionaryRequestMultiError, or nil if none found.
func (m *GetDictionaryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDictionaryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Id != nil {

		if m.GetId() < 1 {
			err := GetDictionaryRequestValidationError{
				field:  "Id",
				reason: "value must be greater than or equal to 1",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Keyword != nil {

		if utf8.RuneCountInString(m.GetKeyword()) < 1 {
			err := GetDictionaryRequestValidationError{
				field:  "Keyword",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return GetDictionaryRequestMultiError(errors)
	}

	return nil
}

// GetDictionaryRequestMultiError is an error wrapping multiple validation
// errors returned by GetDictionaryRequest.ValidateAll() if the designated
// constraints aren't met.
type GetDictionaryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDictionaryRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDictionaryRequestMultiError) AllErrors() []error { return m }

// GetDictionaryRequestValidationError is the validation error returned by
// GetDictionaryRequest.Validate if the designated constraints aren't met.
type GetDictionaryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDictionaryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDictionaryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDictionaryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDictionaryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDictionaryRequestValidationError) ErrorName() string {
	return "GetDictionaryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDictionaryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDictionaryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDictionaryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDictionaryRequestValidationError{}

// Validate checks the field values on GetDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDictionaryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDictionaryReplyMultiError, or nil if none found.
func (m *GetDictionaryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDictionaryReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Keyword

	// no validation rules for Name

	// no validation rules for Type

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if m.Description != nil {
		// no validation rules for Description
	}

	if len(errors) > 0 {
		return GetDictionaryReplyMultiError(errors)
	}

	return nil
}

// GetDictionaryReplyMultiError is an error wrapping multiple validation errors
// returned by GetDictionaryReply.ValidateAll() if the designated constraints
// aren't met.
type GetDictionaryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDictionaryReplyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDictionaryReplyMultiError) AllErrors() []error { return m }

// GetDictionaryReplyValidationError is the validation error returned by
// GetDictionaryReply.Validate if the designated constraints aren't met.
type GetDictionaryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDictionaryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDictionaryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDictionaryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDictionaryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDictionaryReplyValidationError) ErrorName() string {
	return "GetDictionaryReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetDictionaryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDictionaryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDictionaryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDictionaryReplyValidationError{}

// Validate checks the field values on ListDictionaryReply_Dictionary with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListDictionaryReply_Dictionary) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListDictionaryReply_Dictionary with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ListDictionaryReply_DictionaryMultiError, or nil if none found.
func (m *ListDictionaryReply_Dictionary) ValidateAll() error {
	return m.validate(true)
}

func (m *ListDictionaryReply_Dictionary) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Keyword

	// no validation rules for Type

	// no validation rules for Name

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if m.Description != nil {
		// no validation rules for Description
	}

	if len(errors) > 0 {
		return ListDictionaryReply_DictionaryMultiError(errors)
	}

	return nil
}

// ListDictionaryReply_DictionaryMultiError is an error wrapping multiple
// validation errors returned by ListDictionaryReply_Dictionary.ValidateAll()
// if the designated constraints aren't met.
type ListDictionaryReply_DictionaryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListDictionaryReply_DictionaryMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListDictionaryReply_DictionaryMultiError) AllErrors() []error { return m }

// ListDictionaryReply_DictionaryValidationError is the validation error
// returned by ListDictionaryReply_Dictionary.Validate if the designated
// constraints aren't met.
type ListDictionaryReply_DictionaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDictionaryReply_DictionaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDictionaryReply_DictionaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDictionaryReply_DictionaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDictionaryReply_DictionaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDictionaryReply_DictionaryValidationError) ErrorName() string {
	return "ListDictionaryReply_DictionaryValidationError"
}

// Error satisfies the builtin error interface
func (e ListDictionaryReply_DictionaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDictionaryReply_Dictionary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDictionaryReply_DictionaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDictionaryReply_DictionaryValidationError{}
