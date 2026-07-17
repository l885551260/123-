// Modified by aytdai on 2026-07-18 under AGPLv3
// Copyright (C) 2023-2026 QuantumNous (original project)
// Phone/SMS registration feature added by aytdai.

package common

import (
	"errors"
	"fmt"
	"regexp"
)

// ErrSMSRateLimited is returned when the phone/IP has exceeded sending limits.
var ErrSMSRateLimited = errors.New("sms rate limited")

// ErrSMSInvalidPhone is returned when the phone format is invalid.
var ErrSMSInvalidPhone = errors.New("invalid phone number")

// SMSProvider is the interface implemented by every SMS backend (aliyun,
// console, etc.). SendCode must be idempotent for the same (phone, code)
// pair within the code TTL window; providers are not responsible for rate
// limiting — that is handled by SendCodeWithLimitCheck below.
type SMSProvider interface {
	SendCode(phone, code string) error
}

// GetSMSProvider returns the provider configured by the admin via the
// SMSServiceProvider option. Unknown values fall back to the console
// provider, which is safe for development and avoids silent failures.
func GetSMSProvider() SMSProvider {
	switch SMSServiceProvider {
	case "aliyun":
		return &AliyunSMSProvider{}
	case "console":
		fallthrough
	default:
		return &ConsoleSMSProvider{}
	}
}

// phoneRegex matches mainland China mobile numbers (11 digits, starting
// with 1 followed by 3-9). International numbers are intentionally not
// supported in this first iteration.
var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

// IsValidPhone returns true if phone matches the Chinese mobile format.
func IsValidPhone(phone string) bool {
	return phoneRegex.MatchString(phone)
}

// SendCodeWithLimitCheck validates the phone format, enforces per-phone
// and per-IP rate limits, generates a verification code, registers it in
// the in-memory verification map, and finally dispatches the SMS through
// the configured provider. On success, it returns nil and the caller can
// tell the user "SMS code sent". On failure, the verification code is
// NOT registered, so a retry can simply call this function again.
func SendCodeWithLimitCheck(phone, ip string) error {
	if !IsValidPhone(phone) {
		return ErrSMSInvalidPhone
	}

	// Rate limiting — per-phone interval, per-phone daily, per-IP hourly/daily.
	if err := SMSCheckRateLimit(phone, ip); err != nil {
		return err
	}

	// Generate a numeric code of configured length.
	code := GenerateNumericCode(SMSCodeLength)

	// Dispatch through the configured provider.
	provider := GetSMSProvider()
	if err := provider.SendCode(phone, code); err != nil {
		// Log the failure but do not record the send in the limiter,
		// so the user can retry without burning their daily quota.
		SysLog(fmt.Sprintf("[SMS] send to %s failed: %v", phone, err))
		return err
	}

	// Record successful send in the limiter (this increments counters).
	SMSRecordSend(phone, ip)

	// Register the code in the shared verification map. TTL is respected
	// by VerifyCodeWithKey via VerificationValidMinutes.
	RegisterVerificationCodeWithKey(phone, code, SMSVerificationPurpose)

	SysLog(fmt.Sprintf("[SMS] code sent to %s (provider=%s, ip=%s)", phone, SMSServiceProvider, ip))
	return nil
}

// GenerateNumericCode returns a numeric string of the given length, using
// a deterministic substring of a UUID. This mirrors the style of
// GenerateVerificationCode but forces digits only for SMS friendliness.
func GenerateNumericCode(length int) string {
	if length <= 0 {
		length = 6
	}
	u := GenerateVerificationCode(0) // 32 hex chars
	out := make([]byte, 0, length)
	for i := 0; i < len(u) && len(out) < length; i++ {
		c := u[i]
		if c >= '0' && c <= '9' {
			out = append(out, c)
		} else if c >= 'a' && c <= 'f' {
			// Map hex letter to a digit (a->1, b->2, ... f->6).
			out = append(out, '1'+(c-'a'))
		}
	}
	// Pad with '0' if somehow not enough digits (extremely unlikely).
	for len(out) < length {
		out = append(out, '0')
	}
	return string(out[:length])
}
