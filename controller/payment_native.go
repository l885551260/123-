package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/i18n"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/setting/operation_setting"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// isMobileDevice checks User-Agent to determine if the request comes from a mobile device.
func isMobileDevice(userAgent string) bool {
	ua := strings.ToLower(userAgent)
	return strings.Contains(ua, "mobile") || strings.Contains(ua, "android") ||
		strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad") ||
		strings.Contains(ua, "windows phone")
}

// ============================================================================
// Alipay Direct Payment
// ============================================================================

// RequestAlipayPay handles POST /api/user/alipay/pay
func RequestAlipayPay(c *gin.Context) {
	if !common.AlipayEnabled {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgPaymentAlipayDisabled)})
		return
	}

	var req struct {
		Amount        int64  `json:"amount"`
		PaymentMethod string `json:"payment_method"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgInvalidParams)})
		return
	}
	if req.Amount < getMinTopup() {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": fmt.Sprintf(i18n.T(c, i18n.MsgPaymentMinTopup), getMinTopup())})
		return
	}

	id := c.GetInt("id")
	group, err := model.GetUserGroup(id, true)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgPaymentGetGroupFailed)})
		return
	}

	payMoney := getPayMoney(req.Amount, group)
	if payMoney < 0.01 {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgPaymentAmountTooLow)})
		return
	}

	// Generate trade number
	tradeNo := fmt.Sprintf("AL%d%s%d", id, common.GetRandomString(6), time.Now().Unix())

	callBackAddress := service.GetCallbackAddress()
	notifyUrl := callBackAddress + "/api/user/alipay/notify"
	returnUrl := callBackAddress + "/console/log"

	isMobile := isMobileDevice(c.Request.UserAgent())
	amountStr := strconv.FormatFloat(payMoney, 'f', 2, 64)
	subject := fmt.Sprintf("TUC%d", req.Amount)

	payURL, qrCode, err := service.CreateAlipayTrade(tradeNo, subject, amountStr, c.ClientIP(), isMobile, notifyUrl, returnUrl)
	if err != nil {
		logger.LogError(c.Request.Context(), fmt.Sprintf("支付宝 创建订单失败 user_id=%d trade_no=%s error=%q", id, tradeNo, err.Error()))
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgPaymentCreateFailed)})
		return
	}

	// Handle token display mode
	amount := req.Amount
	if operation_setting.GetQuotaDisplayType() == operation_setting.QuotaDisplayTypeTokens {
		dAmount := decimal.NewFromInt(int64(amount))
		dQuotaPerUnit := decimal.NewFromFloat(common.QuotaPerUnit)
		amount = dAmount.Div(dQuotaPerUnit).IntPart()
	}

	// Create TopUp record
	topUp := &model.TopUp{
		UserId:          id,
		Amount:          amount,
		Money:           payMoney,
		TradeNo:         tradeNo,
		PaymentMethod:   "alipay",
		PaymentProvider: model.PaymentProviderAlipay,
		CreateTime:      time.Now().Unix(),
		Status:          common.TopUpStatusPending,
	}
	if err := topUp.Insert(); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgPaymentCreateFailed)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"pay_url":  payURL,
			"qr_code":  qrCode,
			"trade_no": tradeNo,
		},
	})
}

// AlipayNotify handles async payment notification from Alipay.
func AlipayNotify(c *gin.Context) {
	outTradeNo, tradeStatus, err := service.VerifyAlipayNotify(c.Request)
	if err != nil {
		common.SysError(fmt.Sprintf("alipay notify verify error: %v", err))
		c.String(http.StatusOK, "fail")
		return
	}

	if tradeStatus == "TRADE_SUCCESS" || tradeStatus == "TRADE_FINISHED" {
		LockOrder(outTradeNo)
		defer UnlockOrder(outTradeNo)

		topUp := model.GetTopUpByTradeNo(outTradeNo)
		if topUp == nil {
			common.SysError(fmt.Sprintf("alipay notify: topup not found trade_no=%s", outTradeNo))
			c.String(http.StatusOK, "fail")
			return
		}

		if topUp.PaymentProvider != model.PaymentProviderAlipay {
			common.SysError(fmt.Sprintf("alipay notify: payment provider mismatch trade_no=%s provider=%s", outTradeNo, topUp.PaymentProvider))
			c.String(http.StatusOK, "fail")
			return
		}

		if topUp.Status != common.TopUpStatusPending {
			c.String(http.StatusOK, "success")
			return
		}

		topUp.CompleteTime = common.GetTimestamp()
		topUp.Status = common.TopUpStatusSuccess
		if err := topUp.Update(); err != nil {
			common.SysError(fmt.Sprintf("alipay notify: update topup error: %v", err))
			c.String(http.StatusOK, "fail")
			return
		}

		if err := model.IncreaseUserQuota(topUp.UserId, int(float64(topUp.Amount)*common.QuotaPerUnit), true); err != nil {
			common.SysError(fmt.Sprintf("alipay notify: increase quota error: %v", err))
			c.String(http.StatusOK, "fail")
			return
		}

		model.RecordLog(topUp.UserId, model.LogTypeTopup,
			fmt.Sprintf("Alipay top-up success: %s, amount: %.2f", outTradeNo, topUp.Money))
	}

	c.String(http.StatusOK, "success")
}

// ============================================================================
// WeChat Pay Direct Payment
// ============================================================================

// RequestWxPayPay handles POST /api/user/wxpay/pay
func RequestWxPayPay(c *gin.Context) {
	if !common.WxPayEnabled {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgPaymentWxPayDisabled)})
		return
	}

	var req struct {
		Amount        int64  `json:"amount"`
		PaymentMethod string `json:"payment_method"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgInvalidParams)})
		return
	}
	if req.Amount < getMinTopup() {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": fmt.Sprintf(i18n.T(c, i18n.MsgPaymentMinTopup), getMinTopup())})
		return
	}

	id := c.GetInt("id")
	group, err := model.GetUserGroup(id, true)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgPaymentGetGroupFailed)})
		return
	}

	payMoney := getPayMoney(req.Amount, group)
	if payMoney < 0.01 {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgPaymentAmountTooLow)})
		return
	}

	// Generate trade number
	tradeNo := fmt.Sprintf("WX%d%s%d", id, common.GetRandomString(6), time.Now().Unix())

	callBackAddress := service.GetCallbackAddress()
	notifyUrl := callBackAddress + "/api/user/wxpay/notify"

	isMobile := isMobileDevice(c.Request.UserAgent())
	amountFen := int64(payMoney * 100) // WeChat Pay uses fen (cents)
	description := fmt.Sprintf("TUC%d", req.Amount)

	codeUrl, h5Url, err := service.CreateWxPayOrder(tradeNo, description, amountFen, c.ClientIP(), isMobile, notifyUrl)
	if err != nil {
		logger.LogError(c.Request.Context(), fmt.Sprintf("微信支付 创建订单失败 user_id=%d trade_no=%s error=%q", id, tradeNo, err.Error()))
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgPaymentCreateFailed)})
		return
	}

	// Handle token display mode
	amount := req.Amount
	if operation_setting.GetQuotaDisplayType() == operation_setting.QuotaDisplayTypeTokens {
		dAmount := decimal.NewFromInt(int64(amount))
		dQuotaPerUnit := decimal.NewFromFloat(common.QuotaPerUnit)
		amount = dAmount.Div(dQuotaPerUnit).IntPart()
	}

	// Create TopUp record
	topUp := &model.TopUp{
		UserId:          id,
		Amount:          amount,
		Money:           payMoney,
		TradeNo:         tradeNo,
		PaymentMethod:   "wxpay",
		PaymentProvider: model.PaymentProviderWxPay,
		CreateTime:      time.Now().Unix(),
		Status:          common.TopUpStatusPending,
	}
	if err := topUp.Insert(); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": i18n.T(c, i18n.MsgPaymentCreateFailed)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"qr_code":  codeUrl,
			"pay_url":  h5Url,
			"trade_no": tradeNo,
		},
	})
}

// WxPayNotify handles async payment notification from WeChat Pay.
func WxPayNotify(c *gin.Context) {
	result, err := service.VerifyWxPayNotify(c.Request)
	if err != nil {
		common.SysError(fmt.Sprintf("wxpay notify verify error: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"code": "FAIL", "message": err.Error()})
		return
	}

	if result.TradeState == "SUCCESS" {
		LockOrder(result.OutTradeNo)
		defer UnlockOrder(result.OutTradeNo)

		topUp := model.GetTopUpByTradeNo(result.OutTradeNo)
		if topUp == nil {
			common.SysError(fmt.Sprintf("wxpay notify: topup not found trade_no=%s", result.OutTradeNo))
			c.JSON(http.StatusInternalServerError, gin.H{"code": "FAIL", "message": "order not found"})
			return
		}

		if topUp.PaymentProvider != model.PaymentProviderWxPay {
			common.SysError(fmt.Sprintf("wxpay notify: payment provider mismatch trade_no=%s provider=%s", result.OutTradeNo, topUp.PaymentProvider))
			c.JSON(http.StatusInternalServerError, gin.H{"code": "FAIL", "message": "provider mismatch"})
			return
		}

		if topUp.Status != common.TopUpStatusPending {
			c.JSON(http.StatusOK, gin.H{"code": "SUCCESS", "message": "OK"})
			return
		}

		topUp.CompleteTime = common.GetTimestamp()
		topUp.Status = common.TopUpStatusSuccess
		if err := topUp.Update(); err != nil {
			common.SysError(fmt.Sprintf("wxpay notify: update topup error: %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{"code": "FAIL", "message": "update error"})
			return
		}

		if err := model.IncreaseUserQuota(topUp.UserId, int(float64(topUp.Amount)*common.QuotaPerUnit), true); err != nil {
			common.SysError(fmt.Sprintf("wxpay notify: increase quota error: %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{"code": "FAIL", "message": "quota error"})
			return
		}

		model.RecordLog(topUp.UserId, model.LogTypeTopup,
			fmt.Sprintf("WeChat Pay top-up success: %s, amount: %.2f, txn: %s", result.OutTradeNo, topUp.Money, result.TransactionId))
	}

	c.JSON(http.StatusOK, gin.H{"code": "SUCCESS", "message": "OK"})
}
