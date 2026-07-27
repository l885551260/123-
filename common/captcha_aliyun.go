// Added on 2026-07-23 by project contributor under AGPLv3
// Aliyun Captcha 2.0 server-side verification (VerifyIntelligentCaptcha).
//
// The frontend pops an Aliyun captcha widget when the user requests an SMS
// code and forwards the resulting captchaVerifyParam verbatim. This file
// verifies that param against Aliyun's captcha service before the SMS is
// actually sent, blocking bots / SMS-bombing. It reuses the existing
// alibaba-cloud-sdk-go dependency (same SDK used for Aliyun SMS) via a
// CommonRequest, so no new module is introduced.

package common

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

// aliyunCaptchaResponse models the VerifyIntelligentCaptcha reply. Only the
// fields we rely on are declared; the remainder is ignored.
type aliyunCaptchaResponse struct {
	Code    string `json:"Code"`
	Success bool   `json:"Success"`
	Result  struct {
		VerifyCode   string `json:"VerifyCode"`
		VerifyResult bool   `json:"VerifyResult"`
	} `json:"Result"`
}

// VerifyAliyunCaptcha verifies the captchaVerifyParam produced by the frontend
// Aliyun Captcha 2.0 widget. It returns true only when Aliyun confirms the
// verification passed (Result.VerifyResult == true).
//
// The captchaVerifyParam must be forwarded verbatim from the client — Aliyun
// rejects any tampering. Region/endpoint is cn-shanghai (中国内地), which must
// match the frontend region ("cn").
func VerifyAliyunCaptcha(captchaVerifyParam string) (bool, error) {
	if AliyunCaptchaAccessKeyID == "" || AliyunCaptchaAccessKeySecret == "" {
		return false, fmt.Errorf("aliyun captcha credentials not configured")
	}
	if captchaVerifyParam == "" {
		return false, fmt.Errorf("empty captcha verify param")
	}

	client, err := sdk.NewClientWithAccessKey("cn-shanghai", AliyunCaptchaAccessKeyID, AliyunCaptchaAccessKeySecret)
	if err != nil {
		return false, fmt.Errorf("aliyun captcha client init: %w", err)
	}

	req := requests.NewCommonRequest()
	req.Method = "POST"
	req.Scheme = "https"
	req.Domain = "captcha.cn-shanghai.aliyuncs.com"
	req.Version = "2023-03-05"
	req.ApiName = "VerifyIntelligentCaptcha"
	req.QueryParams["CaptchaVerifyParam"] = captchaVerifyParam
	// SceneId is optional but recommended: it prevents the frontend from being
	// tampered to reuse a param from a different scene.
	if AliyunCaptchaSceneId != "" {
		req.QueryParams["SceneId"] = AliyunCaptchaSceneId
	}

	resp, err := client.ProcessCommonRequest(req)
	if err != nil {
		return false, fmt.Errorf("aliyun captcha verify request: %w", err)
	}

	var parsed aliyunCaptchaResponse
	if err := json.Unmarshal(resp.GetHttpContentBytes(), &parsed); err != nil {
		return false, fmt.Errorf("aliyun captcha parse response: %w", err)
	}
	if !parsed.Success {
		return false, fmt.Errorf("aliyun captcha business error: code=%s verify_code=%s", parsed.Code, parsed.Result.VerifyCode)
	}
	return parsed.Result.VerifyResult, nil
}
