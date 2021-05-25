// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pb/vcs_connection.proto

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

	"google.golang.org/protobuf/types/known/anypb"

	types "alfred/protos/types"
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

	_ = types.VCSProviders(0)

	_ = types.VCSProviders(0)

	_ = types.VCSProviders(0)

	_ = types.VCSProviders(0)

	_ = types.VCSProviders(0)

	_ = types.VCSProviders(0)

	_ = types.VCSProviders(0)
)

// Validate checks the field values on ListAllSupportedVCSProvidersResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ListAllSupportedVCSProvidersResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListAllSupportedVCSProvidersResponseValidationError is the validation error
// returned by ListAllSupportedVCSProvidersResponse.Validate if the designated
// constraints aren't met.
type ListAllSupportedVCSProvidersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAllSupportedVCSProvidersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAllSupportedVCSProvidersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAllSupportedVCSProvidersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAllSupportedVCSProvidersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAllSupportedVCSProvidersResponseValidationError) ErrorName() string {
	return "ListAllSupportedVCSProvidersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAllSupportedVCSProvidersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAllSupportedVCSProvidersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAllSupportedVCSProvidersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAllSupportedVCSProvidersResponseValidationError{}

// Validate checks the field values on AuthorizeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AuthorizeRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Provider

	// no validation rules for Label

	return nil
}

// AuthorizeRequestValidationError is the validation error returned by
// AuthorizeRequest.Validate if the designated constraints aren't met.
type AuthorizeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorizeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorizeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorizeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorizeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorizeRequestValidationError) ErrorName() string { return "AuthorizeRequestValidationError" }

// Error satisfies the builtin error interface
func (e AuthorizeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthorizeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorizeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorizeRequestValidationError{}

// Validate checks the field values on AuthorizeResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AuthorizeResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RedirectUrl

	return nil
}

// AuthorizeResponseValidationError is the validation error returned by
// AuthorizeResponse.Validate if the designated constraints aren't met.
type AuthorizeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorizeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorizeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorizeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorizeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorizeResponseValidationError) ErrorName() string {
	return "AuthorizeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthorizeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthorizeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorizeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorizeResponseValidationError{}

// Validate checks the field values on CallbackRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CallbackRequest) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := _CallbackRequest_Provider_NotInLookup[m.GetProvider()]; ok {
		return CallbackRequestValidationError{
			field:  "Provider",
			reason: "value must not be in list [0]",
		}
	}

	// no validation rules for State

	if utf8.RuneCountInString(m.GetCode()) < 3 {
		return CallbackRequestValidationError{
			field:  "Code",
			reason: "value length must be at least 3 runes",
		}
	}

	if utf8.RuneCountInString(m.GetAccountId()) < 3 {
		return CallbackRequestValidationError{
			field:  "AccountId",
			reason: "value length must be at least 3 runes",
		}
	}

	return nil
}

// CallbackRequestValidationError is the validation error returned by
// CallbackRequest.Validate if the designated constraints aren't met.
type CallbackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallbackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallbackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallbackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallbackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallbackRequestValidationError) ErrorName() string { return "CallbackRequestValidationError" }

// Error satisfies the builtin error interface
func (e CallbackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallbackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallbackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallbackRequestValidationError{}

var _CallbackRequest_Provider_NotInLookup = map[types.VCSProviders]struct{}{
	0: {},
}

// Validate checks the field values on AccountVCSConnection with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AccountVCSConnection) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 3 {
		return AccountVCSConnectionValidationError{
			field:  "Id",
			reason: "value length must be at least 3 runes",
		}
	}

	// no validation rules for Provider

	if utf8.RuneCountInString(m.GetAccountId()) < 3 {
		return AccountVCSConnectionValidationError{
			field:  "AccountId",
			reason: "value length must be at least 3 runes",
		}
	}

	// no validation rules for Label

	// no validation rules for UserName

	return nil
}

// AccountVCSConnectionValidationError is the validation error returned by
// AccountVCSConnection.Validate if the designated constraints aren't met.
type AccountVCSConnectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountVCSConnectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountVCSConnectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountVCSConnectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountVCSConnectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountVCSConnectionValidationError) ErrorName() string {
	return "AccountVCSConnectionValidationError"
}

// Error satisfies the builtin error interface
func (e AccountVCSConnectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountVCSConnection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountVCSConnectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountVCSConnectionValidationError{}

// Validate checks the field values on ListVCSConnectionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListVCSConnectionRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccountId

	// no validation rules for Provider

	return nil
}

// ListVCSConnectionRequestValidationError is the validation error returned by
// ListVCSConnectionRequest.Validate if the designated constraints aren't met.
type ListVCSConnectionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListVCSConnectionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListVCSConnectionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListVCSConnectionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListVCSConnectionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListVCSConnectionRequestValidationError) ErrorName() string {
	return "ListVCSConnectionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListVCSConnectionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListVCSConnectionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListVCSConnectionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListVCSConnectionRequestValidationError{}

// Validate checks the field values on ListVCSConnectionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListVCSConnectionResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetVcsConnection() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListVCSConnectionResponseValidationError{
					field:  fmt.Sprintf("VcsConnection[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListVCSConnectionResponseValidationError is the validation error returned by
// ListVCSConnectionResponse.Validate if the designated constraints aren't met.
type ListVCSConnectionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListVCSConnectionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListVCSConnectionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListVCSConnectionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListVCSConnectionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListVCSConnectionResponseValidationError) ErrorName() string {
	return "ListVCSConnectionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListVCSConnectionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListVCSConnectionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListVCSConnectionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListVCSConnectionResponseValidationError{}

// Validate checks the field values on GetVCSConnectionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetVCSConnectionRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetAccountId()) < 3 {
		return GetVCSConnectionRequestValidationError{
			field:  "AccountId",
			reason: "value length must be at least 3 runes",
		}
	}

	// no validation rules for Provider

	return nil
}

// GetVCSConnectionRequestValidationError is the validation error returned by
// GetVCSConnectionRequest.Validate if the designated constraints aren't met.
type GetVCSConnectionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetVCSConnectionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetVCSConnectionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetVCSConnectionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetVCSConnectionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetVCSConnectionRequestValidationError) ErrorName() string {
	return "GetVCSConnectionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetVCSConnectionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetVCSConnectionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetVCSConnectionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetVCSConnectionRequestValidationError{}

// Validate checks the field values on RevokeVCSTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RevokeVCSTokenRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Provider

	// no validation rules for VcsId

	return nil
}

// RevokeVCSTokenRequestValidationError is the validation error returned by
// RevokeVCSTokenRequest.Validate if the designated constraints aren't met.
type RevokeVCSTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RevokeVCSTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RevokeVCSTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RevokeVCSTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RevokeVCSTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RevokeVCSTokenRequestValidationError) ErrorName() string {
	return "RevokeVCSTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RevokeVCSTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRevokeVCSTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RevokeVCSTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RevokeVCSTokenRequestValidationError{}

// Validate checks the field values on RenewVCSTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RenewVCSTokenRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Provider

	return nil
}

// RenewVCSTokenRequestValidationError is the validation error returned by
// RenewVCSTokenRequest.Validate if the designated constraints aren't met.
type RenewVCSTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RenewVCSTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RenewVCSTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RenewVCSTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RenewVCSTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RenewVCSTokenRequestValidationError) ErrorName() string {
	return "RenewVCSTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RenewVCSTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRenewVCSTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RenewVCSTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RenewVCSTokenRequestValidationError{}
