// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto

package envoy_config_filter_http_ip_tagging_v2

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

// Validate checks the field values on IPTagging with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IPTagging) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IPTagging with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IPTaggingMultiError, or nil
// if none found.
func (m *IPTagging) ValidateAll() error {
	return m.validate(true)
}

func (m *IPTagging) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := IPTagging_RequestType_name[int32(m.GetRequestType())]; !ok {
		err := IPTaggingValidationError{
			field:  "RequestType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetIpTags()) < 1 {
		err := IPTaggingValidationError{
			field:  "IpTags",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetIpTags() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IPTaggingValidationError{
						field:  fmt.Sprintf("IpTags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IPTaggingValidationError{
						field:  fmt.Sprintf("IpTags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IPTaggingValidationError{
					field:  fmt.Sprintf("IpTags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return IPTaggingMultiError(errors)
	}
	return nil
}

// IPTaggingMultiError is an error wrapping multiple validation errors returned
// by IPTagging.ValidateAll() if the designated constraints aren't met.
type IPTaggingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IPTaggingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IPTaggingMultiError) AllErrors() []error { return m }

// IPTaggingValidationError is the validation error returned by
// IPTagging.Validate if the designated constraints aren't met.
type IPTaggingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IPTaggingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IPTaggingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IPTaggingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IPTaggingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IPTaggingValidationError) ErrorName() string { return "IPTaggingValidationError" }

// Error satisfies the builtin error interface
func (e IPTaggingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPTagging.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IPTaggingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IPTaggingValidationError{}

// Validate checks the field values on IPTagging_IPTag with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IPTagging_IPTag) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IPTagging_IPTag with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IPTagging_IPTagMultiError, or nil if none found.
func (m *IPTagging_IPTag) ValidateAll() error {
	return m.validate(true)
}

func (m *IPTagging_IPTag) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IpTagName

	for idx, item := range m.GetIpList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IPTagging_IPTagValidationError{
						field:  fmt.Sprintf("IpList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IPTagging_IPTagValidationError{
						field:  fmt.Sprintf("IpList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IPTagging_IPTagValidationError{
					field:  fmt.Sprintf("IpList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return IPTagging_IPTagMultiError(errors)
	}
	return nil
}

// IPTagging_IPTagMultiError is an error wrapping multiple validation errors
// returned by IPTagging_IPTag.ValidateAll() if the designated constraints
// aren't met.
type IPTagging_IPTagMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IPTagging_IPTagMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IPTagging_IPTagMultiError) AllErrors() []error { return m }

// IPTagging_IPTagValidationError is the validation error returned by
// IPTagging_IPTag.Validate if the designated constraints aren't met.
type IPTagging_IPTagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IPTagging_IPTagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IPTagging_IPTagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IPTagging_IPTagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IPTagging_IPTagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IPTagging_IPTagValidationError) ErrorName() string { return "IPTagging_IPTagValidationError" }

// Error satisfies the builtin error interface
func (e IPTagging_IPTagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPTagging_IPTag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IPTagging_IPTagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IPTagging_IPTagValidationError{}