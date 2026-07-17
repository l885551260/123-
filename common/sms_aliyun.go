// Modified by aytdai on 2026-07-18 under AGPLv3
// Phone/SMS registration feature — Aliyun dysmsapi provider.

package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

// AliyunSMSProvider implements SMSProvider via the official Aliyun SDK.
// All parameters (AccessKey, SignName, TemplateCode) are read from the
// in-memory common.* variables that the option system keeps in sync with
// the database. Never hard-code secrets in source.
type AliyunSMSProvider struct{}

// SendCode sends a single SMS via Aliyun dysmsapi. If the configuration
// is incomplete (missing AK/SK/SignName/TemplateCode), we fall back to
// console mode and log a clear warning so the operator can fix it.
func (p *AliyunSMSProvider) SendCode(phone, code string) error {
	if SMSAccessKeyID == "" || SMSAccessKeySecret == "" || SMSSignName == "" || SMSTemplateCode == "" {
		SysLog("[SMS-ALIYUN] incomplete configuration, falling back to console mode")
		return (&ConsoleSMSProvider{}).SendCode(phone, code)
	}

	// TemplateParam must be a JSON object string, e.g. {"code":"123456"}.
	templateParam, _ := json.Marshal(map[string]string{"code": code})

	client, err := dysmsapi.NewClientWithAccessKey(
		"cn-hangzhou",
		SMSAccessKeyID,
		SMSAccessKeySecret,
	)
	if err != nil {
		return fmt.Errorf("aliyun sms client init: %w", err)
	}
	// The SDK's underlying sdk.Client already sets a sane HTTP timeout
	// (~5s read). We don't override it here to stay compatible across
	// alibaba-cloud-sdk-go versions.

	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = phone
	req.SignName = SMSSignName
	req.TemplateCode = SMSTemplateCode
	req.TemplateParam = string(templateParam)

	resp, err := client.SendSms(req)
	if err != nil {
		return fmt.Errorf("aliyun sms send: %w", err)
	}
	// Aliyun returns "OK" on success; anything else is a business error.
	if resp == nil || !strings.EqualFold(resp.Code, "OK") {
		code := ""
		msg := ""
		if resp != nil {
			code = resp.Code
			msg = resp.Message
		}
		return fmt.Errorf("aliyun sms business error: code=%s message=%s", code, msg)
	}
	return nil
}

// ConsoleSMSProvider logs the SMS instead of sending it. Used during
// development and on servers that have not yet configured Aliyun.
// The [SMS-CONSOLE] prefix is intentionally stable so the operator can
// grep journalctl to find verification codes while testing.
type ConsoleSMSProvider struct{}

func (p *ConsoleSMSProvider) SendCode(phone, code string) error {
	SysLog(fmt.Sprintf("[SMS-CONSOLE] phone=%s code=%s (configure SMSServiceProvider=aliyun to send real SMS)", phone, code))
	return nil
}

// --- Self-test / dry-run helpers (unused in production, kept for future) ---

var _ SMSProvider = (*AliyunSMSProvider)(nil)
var _ SMSProvider = (*ConsoleSMSProvider)(nil)

// unused-import guards — errors is used in this file; the rest are
// placeholders kept for forward-compat if we split this file later.
var (
	_ = errors.New
)
