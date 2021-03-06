// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: test/v1/test.proto

package testv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// define the regex for a UUID once up-front
var _test_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Test with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Test) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return TestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	// no validation rules for Name

	// no validation rules for Test

	// no validation rules for Update

	// no validation rules for Broken

	// no validation rules for Hello

	// no validation rules for IsCool

	return nil
}

func (m *Test) _validateUuid(uuid string) error {
	if matched := _test_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// TestValidationError is the validation error returned by Test.Validate if the
// designated constraints aren't met.
type TestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestValidationError) ErrorName() string { return "TestValidationError" }

// Error satisfies the builtin error interface
func (e TestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestValidationError{}

// Validate checks the field values on GetTestRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetTestRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetTestRequestValidationError is the validation error returned by
// GetTestRequest.Validate if the designated constraints aren't met.
type GetTestRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTestRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTestRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTestRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTestRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTestRequestValidationError) ErrorName() string { return "GetTestRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTestRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTestRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTestRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTestRequestValidationError{}

// Validate checks the field values on GetTestResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetTestResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTestResponseValidationError{
				field:  "Test",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetTestResponseValidationError is the validation error returned by
// GetTestResponse.Validate if the designated constraints aren't met.
type GetTestResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTestResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTestResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTestResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTestResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTestResponseValidationError) ErrorName() string { return "GetTestResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetTestResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTestResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTestResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTestResponseValidationError{}
