// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: address/v1/api.proto

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

// Validate checks the field values on NewBip44Request with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *NewBip44Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewBip44Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NewBip44RequestMultiError, or nil if none found.
func (m *NewBip44Request) ValidateAll() error {
	return m.validate(true)
}

func (m *NewBip44Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMnemonic()) < 2 {
		err := NewBip44RequestValidationError{
			field:  "Mnemonic",
			reason: "value length must be at least 2 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPass()) < 0 {
		err := NewBip44RequestValidationError{
			field:  "Pass",
			reason: "value length must be at least 0 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Index

	// no validation rules for Count

	// no validation rules for Ty

	// no validation rules for Flag

	if len(errors) > 0 {
		return NewBip44RequestMultiError(errors)
	}

	return nil
}

// NewBip44RequestMultiError is an error wrapping multiple validation errors
// returned by NewBip44Request.ValidateAll() if the designated constraints
// aren't met.
type NewBip44RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewBip44RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewBip44RequestMultiError) AllErrors() []error { return m }

// NewBip44RequestValidationError is the validation error returned by
// NewBip44Request.Validate if the designated constraints aren't met.
type NewBip44RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewBip44RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewBip44RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewBip44RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewBip44RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewBip44RequestValidationError) ErrorName() string { return "NewBip44RequestValidationError" }

// Error satisfies the builtin error interface
func (e NewBip44RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewBip44Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewBip44RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewBip44RequestValidationError{}

// Validate checks the field values on NewBip44Result with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NewBip44Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewBip44Result with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NewBip44ResultMultiError,
// or nil if none found.
func (m *NewBip44Result) ValidateAll() error {
	return m.validate(true)
}

func (m *NewBip44Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WalletAddress

	// no validation rules for PrivKey

	// no validation rules for Id

	// no validation rules for WalletAddressSH

	// no validation rules for PrivKeySH

	// no validation rules for WalletAddressWPKH

	// no validation rules for PrivKeyWPKH

	// no validation rules for WalletAddressTaproot

	// no validation rules for PrivKeyTaproot

	if len(errors) > 0 {
		return NewBip44ResultMultiError(errors)
	}

	return nil
}

// NewBip44ResultMultiError is an error wrapping multiple validation errors
// returned by NewBip44Result.ValidateAll() if the designated constraints
// aren't met.
type NewBip44ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewBip44ResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewBip44ResultMultiError) AllErrors() []error { return m }

// NewBip44ResultValidationError is the validation error returned by
// NewBip44Result.Validate if the designated constraints aren't met.
type NewBip44ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewBip44ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewBip44ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewBip44ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewBip44ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewBip44ResultValidationError) ErrorName() string { return "NewBip44ResultValidationError" }

// Error satisfies the builtin error interface
func (e NewBip44ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewBip44Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewBip44ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewBip44ResultValidationError{}

// Validate checks the field values on NewBip44Reply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NewBip44Reply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewBip44Reply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NewBip44ReplyMultiError, or
// nil if none found.
func (m *NewBip44Reply) ValidateAll() error {
	return m.validate(true)
}

func (m *NewBip44Reply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for Code

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NewBip44ReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NewBip44ReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NewBip44ReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NewBip44ReplyMultiError(errors)
	}

	return nil
}

// NewBip44ReplyMultiError is an error wrapping multiple validation errors
// returned by NewBip44Reply.ValidateAll() if the designated constraints
// aren't met.
type NewBip44ReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewBip44ReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewBip44ReplyMultiError) AllErrors() []error { return m }

// NewBip44ReplyValidationError is the validation error returned by
// NewBip44Reply.Validate if the designated constraints aren't met.
type NewBip44ReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewBip44ReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewBip44ReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewBip44ReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewBip44ReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewBip44ReplyValidationError) ErrorName() string { return "NewBip44ReplyValidationError" }

// Error satisfies the builtin error interface
func (e NewBip44ReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewBip44Reply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewBip44ReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewBip44ReplyValidationError{}

// Validate checks the field values on NewBip441Result with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *NewBip441Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewBip441Result with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NewBip441ResultMultiError, or nil if none found.
func (m *NewBip441Result) ValidateAll() error {
	return m.validate(true)
}

func (m *NewBip441Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WalletAddress

	// no validation rules for PrivKey

	// no validation rules for Id

	if len(errors) > 0 {
		return NewBip441ResultMultiError(errors)
	}

	return nil
}

// NewBip441ResultMultiError is an error wrapping multiple validation errors
// returned by NewBip441Result.ValidateAll() if the designated constraints
// aren't met.
type NewBip441ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewBip441ResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewBip441ResultMultiError) AllErrors() []error { return m }

// NewBip441ResultValidationError is the validation error returned by
// NewBip441Result.Validate if the designated constraints aren't met.
type NewBip441ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewBip441ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewBip441ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewBip441ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewBip441ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewBip441ResultValidationError) ErrorName() string { return "NewBip441ResultValidationError" }

// Error satisfies the builtin error interface
func (e NewBip441ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewBip441Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewBip441ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewBip441ResultValidationError{}

// Validate checks the field values on NewBip441Reply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NewBip441Reply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewBip441Reply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NewBip441ReplyMultiError,
// or nil if none found.
func (m *NewBip441Reply) ValidateAll() error {
	return m.validate(true)
}

func (m *NewBip441Reply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for Code

	// no validation rules for Message

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NewBip441ReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NewBip441ReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NewBip441ReplyValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return NewBip441ReplyMultiError(errors)
	}

	return nil
}

// NewBip441ReplyMultiError is an error wrapping multiple validation errors
// returned by NewBip441Reply.ValidateAll() if the designated constraints
// aren't met.
type NewBip441ReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewBip441ReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewBip441ReplyMultiError) AllErrors() []error { return m }

// NewBip441ReplyValidationError is the validation error returned by
// NewBip441Reply.Validate if the designated constraints aren't met.
type NewBip441ReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewBip441ReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewBip441ReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewBip441ReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewBip441ReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewBip441ReplyValidationError) ErrorName() string { return "NewBip441ReplyValidationError" }

// Error satisfies the builtin error interface
func (e NewBip441ReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewBip441Reply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewBip441ReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewBip441ReplyValidationError{}
