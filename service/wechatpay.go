package service

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"

	"github.com/QuantumNous/new-api/common"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// getWxPayClient creates or returns a WeChat Pay client using public key auth.
func getWxPayClient() (*core.Client, *rsa.PublicKey, error) {
	if common.WxPayMchId == "" || common.WxPayPrivateKey == "" || common.WxPayCertSerialNo == "" {
		return nil, nil, fmt.Errorf("wechat pay not configured")
	}
	if common.WxPayPublicKeyId == "" || common.WxPayPublicKey == "" {
		return nil, nil, fmt.Errorf("wechat pay public key not configured")
	}

	mchPrivateKey, err := utils.LoadPrivateKey(common.WxPayPrivateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("load merchant private key error: %w", err)
	}

	wechatPayPublicKey, err := utils.LoadPublicKey(common.WxPayPublicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("load wechat pay public key error: %w", err)
	}

	ctx := context.Background()
	client, err := core.NewClient(
		ctx,
		option.WithWechatPayPublicKeyAuthCipher(
			common.WxPayMchId,
			common.WxPayCertSerialNo,
			mchPrivateKey,
			common.WxPayPublicKeyId,
			wechatPayPublicKey,
		),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("wechat pay client init error: %w", err)
	}

	return client, wechatPayPublicKey, nil
}

// CreateWxPayOrder creates a WeChat Pay order.
// For PC (isMobile=false): Native payment, returns code_url for QR code generation.
// For Mobile (isMobile=true): H5 payment, returns h5_url for redirect.
func CreateWxPayOrder(tradeNo, description string, amountFen int64, clientIP string, isMobile bool, notifyUrl string) (codeUrl string, h5Url string, err error) {
	client, _, err := getWxPayClient()
	if err != nil {
		return "", "", err
	}

	ctx := context.Background()

	if isMobile {
		svc := h5.H5ApiService{Client: client}
		resp, _, err := svc.Prepay(ctx, h5.PrepayRequest{
			Appid:       core.String(common.WxPayAppId),
			Mchid:       core.String(common.WxPayMchId),
			Description: core.String(description),
			OutTradeNo:  core.String(tradeNo),
			NotifyUrl:   core.String(notifyUrl),
			Amount: &h5.Amount{
				Total: core.Int64(amountFen),
			},
			SceneInfo: &h5.SceneInfo{
				PayerClientIp: core.String(clientIP),
				H5Info: &h5.H5Info{
					Type: core.String("Wap"),
				},
			},
		})
		if err != nil {
			return "", "", fmt.Errorf("wechat pay h5 prepay error: %w", err)
		}
		if resp.H5Url == nil {
			return "", "", fmt.Errorf("wechat pay h5 prepay returned empty h5_url")
		}
		return "", *resp.H5Url, nil
	}

	// Native Payment (PC QR code)
	svc := native.NativeApiService{Client: client}
	resp, _, err := svc.Prepay(ctx, native.PrepayRequest{
		Appid:       core.String(common.WxPayAppId),
		Mchid:       core.String(common.WxPayMchId),
		Description: core.String(description),
		OutTradeNo:  core.String(tradeNo),
		NotifyUrl:   core.String(notifyUrl),
		Amount: &native.Amount{
			Total: core.Int64(amountFen),
		},
	})
	if err != nil {
		return "", "", fmt.Errorf("wechat pay native prepay error: %w", err)
	}
	if resp.CodeUrl == nil {
		return "", "", fmt.Errorf("wechat pay native prepay returned empty code_url")
	}
	return *resp.CodeUrl, "", nil
}

// WxPayNotifyResult holds the parsed notification data.
type WxPayNotifyResult struct {
	OutTradeNo    string
	TradeState    string
	TransactionId string
}

// VerifyWxPayNotify verifies and decrypts a WeChat Pay notification.
func VerifyWxPayNotify(req *http.Request) (*WxPayNotifyResult, error) {
	if common.WxPayPublicKeyId == "" || common.WxPayPublicKey == "" {
		return nil, fmt.Errorf("wechat pay public key not configured")
	}

	wechatPayPublicKey, err := utils.LoadPublicKey(common.WxPayPublicKey)
	if err != nil {
		return nil, fmt.Errorf("load wechat pay public key error: %w", err)
	}

	verifier := verifiers.NewSHA256WithRSAPubkeyVerifier(common.WxPayPublicKeyId, *wechatPayPublicKey)
	handler := notify.NewNotifyHandler(common.WxPayAPIv3Key, verifier)
	transaction := new(payments.Transaction)
	_, err = handler.ParseNotifyRequest(context.Background(), req, transaction)
	if err != nil {
		return nil, fmt.Errorf("wechat pay notify parse error: %w", err)
	}

	result := &WxPayNotifyResult{}
	if transaction.OutTradeNo != nil {
		result.OutTradeNo = *transaction.OutTradeNo
	}
	if transaction.TradeState != nil {
		result.TradeState = *transaction.TradeState
	}
	if transaction.TransactionId != nil {
		result.TransactionId = *transaction.TransactionId
	}
	return result, nil
}
