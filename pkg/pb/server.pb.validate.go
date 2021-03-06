// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: server.proto

package pb

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _server_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Label with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Label) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Value

	return nil
}

// LabelValidationError is the validation error returned by Label.Validate if
// the designated constraints aren't met.
type LabelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LabelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LabelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LabelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LabelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LabelValidationError) ErrorName() string { return "LabelValidationError" }

// Error satisfies the builtin error interface
func (e LabelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLabel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LabelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LabelValidationError{}

// Validate checks the field values on LabelSet with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LabelSet) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetLabels()) < 1 {
		return LabelSetValidationError{
			field:  "Labels",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetLabels() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LabelSetValidationError{
					field:  fmt.Sprintf("Labels[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LabelSetValidationError is the validation error returned by
// LabelSet.Validate if the designated constraints aren't met.
type LabelSetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LabelSetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LabelSetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LabelSetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LabelSetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LabelSetValidationError) ErrorName() string { return "LabelSetValidationError" }

// Error satisfies the builtin error interface
func (e LabelSetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLabelSet.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LabelSetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LabelSetValidationError{}

// Validate checks the field values on RegisterGuestAccountRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegisterGuestAccountRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ServerId

	// no validation rules for AcceptTos

	return nil
}

// RegisterGuestAccountRequestValidationError is the validation error returned
// by RegisterGuestAccountRequest.Validate if the designated constraints
// aren't met.
type RegisterGuestAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterGuestAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterGuestAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterGuestAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterGuestAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterGuestAccountRequestValidationError) ErrorName() string {
	return "RegisterGuestAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterGuestAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterGuestAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterGuestAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterGuestAccountRequestValidationError{}

// Validate checks the field values on RegisterGuestAccountResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegisterGuestAccountResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	return nil
}

// RegisterGuestAccountResponseValidationError is the validation error returned
// by RegisterGuestAccountResponse.Validate if the designated constraints
// aren't met.
type RegisterGuestAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterGuestAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterGuestAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterGuestAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterGuestAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterGuestAccountResponseValidationError) ErrorName() string {
	return "RegisterGuestAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterGuestAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterGuestAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterGuestAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterGuestAccountResponseValidationError{}

// Validate checks the field values on RegisterHostnameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegisterHostnameRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLabels() == nil {
		return RegisterHostnameRequestValidationError{
			field:  "Labels",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetLabels()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterHostnameRequestValidationError{
				field:  "Labels",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Hostname.(type) {

	case *RegisterHostnameRequest_Generate:

		if v, ok := interface{}(m.GetGenerate()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterHostnameRequestValidationError{
					field:  "Generate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *RegisterHostnameRequest_Exact:

		if _, ok := _RegisterHostnameRequest_Exact_NotInLookup[m.GetExact()]; ok {
			return RegisterHostnameRequestValidationError{
				field:  "Exact",
				reason: "value must not be in list [admin api blog hzn horizon waypoint]",
			}
		}

		if utf8.RuneCountInString(m.GetExact()) < 3 {
			return RegisterHostnameRequestValidationError{
				field:  "Exact",
				reason: "value length must be at least 3 runes",
			}
		}

		if !_RegisterHostnameRequest_Exact_Pattern.MatchString(m.GetExact()) {
			return RegisterHostnameRequestValidationError{
				field:  "Exact",
				reason: "value does not match regex pattern \"\\\\w+[\\\\w\\\\d-]*\"",
			}
		}

	default:
		return RegisterHostnameRequestValidationError{
			field:  "Hostname",
			reason: "value is required",
		}

	}

	return nil
}

// RegisterHostnameRequestValidationError is the validation error returned by
// RegisterHostnameRequest.Validate if the designated constraints aren't met.
type RegisterHostnameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterHostnameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterHostnameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterHostnameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterHostnameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterHostnameRequestValidationError) ErrorName() string {
	return "RegisterHostnameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterHostnameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterHostnameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterHostnameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterHostnameRequestValidationError{}

var _RegisterHostnameRequest_Exact_NotInLookup = map[string]struct{}{
	"admin":    {},
	"api":      {},
	"blog":     {},
	"hzn":      {},
	"horizon":  {},
	"waypoint": {},
}

var _RegisterHostnameRequest_Exact_Pattern = regexp.MustCompile("\\w+[\\w\\d-]*")

// Validate checks the field values on RegisterHostnameResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegisterHostnameResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Hostname

	// no validation rules for Fqdn

	return nil
}

// RegisterHostnameResponseValidationError is the validation error returned by
// RegisterHostnameResponse.Validate if the designated constraints aren't met.
type RegisterHostnameResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterHostnameResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterHostnameResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterHostnameResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterHostnameResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterHostnameResponseValidationError) ErrorName() string {
	return "RegisterHostnameResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterHostnameResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterHostnameResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterHostnameResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterHostnameResponseValidationError{}

// Validate checks the field values on ListHostnamesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListHostnamesRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListHostnamesRequestValidationError is the validation error returned by
// ListHostnamesRequest.Validate if the designated constraints aren't met.
type ListHostnamesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListHostnamesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListHostnamesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListHostnamesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListHostnamesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListHostnamesRequestValidationError) ErrorName() string {
	return "ListHostnamesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListHostnamesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListHostnamesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListHostnamesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListHostnamesRequestValidationError{}

// Validate checks the field values on ListHostnamesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListHostnamesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHostnames() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListHostnamesResponseValidationError{
					field:  fmt.Sprintf("Hostnames[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListHostnamesResponseValidationError is the validation error returned by
// ListHostnamesResponse.Validate if the designated constraints aren't met.
type ListHostnamesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListHostnamesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListHostnamesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListHostnamesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListHostnamesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListHostnamesResponseValidationError) ErrorName() string {
	return "ListHostnamesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListHostnamesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListHostnamesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListHostnamesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListHostnamesResponseValidationError{}

// Validate checks the field values on DeleteHostnameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteHostnameRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Hostname

	return nil
}

// DeleteHostnameRequestValidationError is the validation error returned by
// DeleteHostnameRequest.Validate if the designated constraints aren't met.
type DeleteHostnameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteHostnameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteHostnameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteHostnameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteHostnameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteHostnameRequestValidationError) ErrorName() string {
	return "DeleteHostnameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteHostnameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteHostnameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteHostnameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteHostnameRequestValidationError{}

// Validate checks the field values on ListHostnamesResponse_Hostname with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListHostnamesResponse_Hostname) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Hostname

	// no validation rules for Fqdn

	if v, ok := interface{}(m.GetLabels()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListHostnamesResponse_HostnameValidationError{
				field:  "Labels",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListHostnamesResponse_HostnameValidationError is the validation error
// returned by ListHostnamesResponse_Hostname.Validate if the designated
// constraints aren't met.
type ListHostnamesResponse_HostnameValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListHostnamesResponse_HostnameValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListHostnamesResponse_HostnameValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListHostnamesResponse_HostnameValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListHostnamesResponse_HostnameValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListHostnamesResponse_HostnameValidationError) ErrorName() string {
	return "ListHostnamesResponse_HostnameValidationError"
}

// Error satisfies the builtin error interface
func (e ListHostnamesResponse_HostnameValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListHostnamesResponse_Hostname.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListHostnamesResponse_HostnameValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListHostnamesResponse_HostnameValidationError{}
