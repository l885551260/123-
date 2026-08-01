package service

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/QuantumNous/new-api/common"
	"github.com/smartwalle/alipay/v3"
)

// GetAlipayClient creates an Alipay SDK client from the configured credentials.
// Returns nil if the required configuration is missing.
func GetAlipayClient() *alipay.Client {
	if common.AlipayAppId == "" || common.AlipayPrivateKey == "" {
		return nil
	}
	client, err := alipay.New(common.AlipayAppId, common.AlipayPrivateKey, true)
	if err != nil {
		common.SysError("alipay client init error: " + err.Error())
		return nil
	}
	// Load alipay public key for notification verification
	if common.AlipayPublicKey != "" {
		_ = client.LoadAliPayPublicKey(common.AlipayPublicKey)
	}
	return client
}

// CreateAlipayTrade creates an Alipay payment order.
// Uses TradeWapPay (手机网站支付) for mobile browsers — auto-opens Alipay app.
// Uses TradePagePay (电脑网站支付) for desktop browsers.
func CreateAlipayTrade(tradeNo, subject, amount, clientIP string, isMobile bool, notifyUrl, returnUrl string) (payURL string, qrCode string, err error) {
	client := GetAlipayClient()
	if client == nil {
		return "", "", fmt.Errorf("alipay client not configured")
	}

	if isMobile {
		p := alipay.TradeWapPay{}
		p.NotifyURL = notifyUrl
		p.ReturnURL = returnUrl
		p.Subject = subject
		p.OutTradeNo = tradeNo
		p.TotalAmount = amount
		p.ProductCode = "QUICK_WAP_WAY"

		payUrl, err := client.TradeWapPay(p)
		if err != nil {
			return "", "", fmt.Errorf("alipay wap pay create error: %w", err)
		}
		return payUrl.String(), "", nil
	}

	// Desktop: 电脑网站支付
	p := alipay.TradePagePay{}
	p.NotifyURL = notifyUrl
	p.ReturnURL = returnUrl
	p.Subject = subject
	p.OutTradeNo = tradeNo
	p.TotalAmount = amount
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	payUrl, err := client.TradePagePay(p)
	if err != nil {
		return "", "", fmt.Errorf("alipay page pay create error: %w", err)
	}
	return payUrl.String(), "", nil
}

// VerifyAlipayNotify verifies an Alipay async notification.
// Returns the out_trade_no (our trade number) and trade_status if valid.
func VerifyAlipayNotify(req *http.Request) (outTradeNo string, tradeStatus string, err error) {
	client := GetAlipayClient()
	if client == nil {
		return "", "", fmt.Errorf("alipay client not configured")
	}

	// Parse form if not already parsed
	if err := req.ParseForm(); err != nil {
		return "", "", fmt.Errorf("alipay notify parse form error: %w", err)
	}

	notification, err := client.DecodeNotification(context.Background(), req.Form)
	if err != nil {
		return "", "", fmt.Errorf("alipay notify verify error: %w", err)
	}
	return notification.OutTradeNo, string(notification.TradeStatus), nil
}

// unused import guard
var _ = url.Values{}
