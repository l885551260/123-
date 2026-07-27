package main

import (
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

// One-shot probe: verify alibaba-cloud-sdk-go CommonRequest can call the
// Aliyun Captcha 2.0 VerifyIntelligentCaptcha API with the old SDK signer.
// Usage: go run ./tmp_captcha_check <AccessKeyID> <AccessKeySecret> <SceneId>
func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage: go run ./tmp_captcha_check <ak> <sk> <sceneId>")
		return
	}
	ak := os.Args[1]
	sk := os.Args[2]
	sceneId := os.Args[3]

	client, err := sdk.NewClientWithAccessKey("cn-shanghai", ak, sk)
	if err != nil {
		fmt.Println("CLIENT_INIT_ERROR:", err)
		return
	}
	req := requests.NewCommonRequest()
	req.Method = "POST"
	req.Scheme = "https"
	req.Domain = "captcha.cn-shanghai.aliyuncs.com"
	req.Version = "2023-03-05"
	req.ApiName = "VerifyIntelligentCaptcha"
	req.QueryParams["CaptchaVerifyParam"] = "probe_invalid_param"
	req.QueryParams["SceneId"] = sceneId

	resp, err := client.ProcessCommonRequest(req)
	if err != nil {
		fmt.Println("API_ERROR:", err)
		return
	}
	fmt.Println("HTTP_STATUS:", resp.GetHttpStatus())
	fmt.Println("RESPONSE:", resp.GetHttpContentString())
}
