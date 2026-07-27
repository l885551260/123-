package common

import (
	"crypto/tls"
	//"os"
	//"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

var StartTime = time.Now().Unix() // unit: second
var Version = "v0.0.0"            // this hard coding will be replaced automatically when building, no need to manually change
var SystemName = "New API"
var Footer = ""
var Logo = ""
var TopUpLink = ""

var themeValue atomic.Value // stores string; safe for concurrent read/write

func init() {
	themeValue.Store("classic")
}

func GetTheme() string {
	return themeValue.Load().(string)
}

// SetTheme updates the frontend theme atomically.
// Only "default" and "classic" are accepted; other values are silently ignored.
func SetTheme(t string) {
	if t == "default" || t == "classic" {
		themeValue.Store(t)
	}
}

// ThemeAwarePath rewrites legacy /console/* paths to the default-theme
// equivalents when the active theme is "default".  For "classic" (or any
// other theme) the path is returned unchanged.  The function only touches
// known prefixes so it is safe to call with arbitrary suffixes and query
// strings.
func ThemeAwarePath(suffix string) string {
	if GetTheme() != "default" {
		return suffix
	}
	switch {
	case strings.HasPrefix(suffix, "/console/topup"):
		return strings.Replace(suffix, "/console/topup", "/wallet", 1)
	case strings.HasPrefix(suffix, "/console/log"):
		return strings.Replace(suffix, "/console/log", "/usage-logs", 1)
	case strings.HasPrefix(suffix, "/console/personal"):
		return strings.Replace(suffix, "/console/personal", "/profile", 1)
	}
	return suffix
}

// var ChatLink = ""
// var ChatLink2 = ""
var QuotaPerUnit = 500 * 1000.0 // $0.002 / 1K tokens
// 保留旧变量以兼容历史逻辑，实际展示由 general_setting.quota_display_type 控制
var DisplayInCurrencyEnabled = true
var DisplayTokenStatEnabled = true
var DrawingEnabled = true
var TaskEnabled = true
var DataExportEnabled = true
var DataExportInterval = 5         // unit: minute
var DataExportDefaultTime = "hour" // unit: minute
var DefaultCollapseSidebar = false // default value of collapse sidebar

// Any options with "Secret", "Token" in its key won't be return by GetOptions

var SessionSecret = uuid.New().String()
var CryptoSecret = uuid.New().String()
var SessionCookieSecure = false
var SessionCookieTrustedURLs []string

var OptionMap map[string]string
var OptionMapRWMutex sync.RWMutex

var ItemsPerPage = 10
var MaxRecentItems = 1000

var PasswordLoginEnabled = true
var PasswordRegisterEnabled = true
var EmailVerificationEnabled = false
var GitHubOAuthEnabled = false
var LinuxDOOAuthEnabled = false
var WeChatAuthEnabled = false
var TelegramOAuthEnabled = false
var TurnstileCheckEnabled = false
var RegisterEnabled = true

var EmailDomainRestrictionEnabled = false // 是否启用邮箱域名限制
var EmailAliasRestrictionEnabled = false  // 是否启用邮箱别名限制
var EmailDomainWhitelist = []string{
	"gmail.com",
	"163.com",
	"126.com",
	"qq.com",
	"outlook.com",
	"hotmail.com",
	"icloud.com",
	"yahoo.com",
	"foxmail.com",
}
var EmailLoginAuthServerList = []string{
	"smtp.sendcloud.net",
	"smtp.azurecomm.net",
}

var DebugEnabled bool
var MemoryCacheEnabled bool

var LogConsumeEnabled = true

var TLSInsecureSkipVerify bool
var InsecureTLSConfig = &tls.Config{InsecureSkipVerify: true}

var SMTPServer = ""
var SMTPPort = 587
var SMTPSSLEnabled = false
var SMTPStartTLSEnabled = false
var SMTPInsecureSkipVerify = false
var SMTPForceAuthLogin = false
var SMTPAccount = ""
var SMTPFrom = ""
var SMTPToken = ""

var GitHubClientId = ""
var GitHubClientSecret = ""
var LinuxDOClientId = ""
var LinuxDOClientSecret = ""
var LinuxDOMinimumTrustLevel = 0

var WeChatServerAddress = ""
var WeChatServerToken = ""
var WeChatAccountQRCodeImageURL = ""

var TurnstileSiteKey = ""
var TurnstileSecretKey = ""

// ============================================================================
// Aliyun Captcha 2.0 human verification (added on 2026-07-23 by project contributor, AGPLv3)
// Protects the "send SMS code" action against bots / SMS-bombing. The frontend
// pops an Aliyun captcha widget; the backend verifies the returned param via
// the VerifyIntelligentCaptcha API (see common/captcha_aliyun.go). All values
// are managed through the option system / admin panel, never hard-coded here.
// ============================================================================

// AliyunCaptchaEnabled gates the human-verification check on SMS sending.
var AliyunCaptchaEnabled = false

// AliyunCaptchaPrefix is the identity prefix (身份标) from the Aliyun captcha
// console; the frontend needs it to initialize the captcha widget.
var AliyunCaptchaPrefix = ""

// AliyunCaptchaSceneId is the Web/H5 scene id (场景ID) created in the console.
var AliyunCaptchaSceneId = ""

// Aliyun captcha server-side verification credentials. Must belong to a RAM
// sub-account holding the AliyunYundunAFSFullAccess policy. Never commit real
// values to source; set them via the admin panel / database.
var AliyunCaptchaAccessKeyID = ""
var AliyunCaptchaAccessKeySecret = ""

// ============================================================================
// SMS / Phone verification (added on 2026-07-18 by project contributor, AGPLv3)
// Mainland-China mobile numbers only; see common/sms.go for the dispatcher
// and common/sms_limit.go for the rate-limiter.
// ============================================================================

// PhoneVerificationEnabled gates the whole phone-verification feature.
// When false, the register UI hides the phone block and the backend
// skips all phone-related checks. Defaults to false so operators must
// explicitly opt in via the admin panel.
var PhoneVerificationEnabled = false

// SMSServiceProvider selects the backend dispatcher. Accepted values:
//   - "console" (default): log the code to journalctl, do NOT send SMS.
//     Safe for development and servers without Aliyun credentials.
//   - "aliyun": use Aliyun dysmsapi via common/sms_aliyun.go.
var SMSServiceProvider = "console"

// Aliyun credentials — MUST be set via the admin panel / database,
// never committed to source. The provider falls back to console mode
// when any of these is empty.
var SMSAccessKeyID = ""
var SMSAccessKeySecret = ""

// Aliyun SMS signature and template. Must be approved in the Aliyun
// console before sending real SMS. Example: SignName="我的网站",
// TemplateCode="SMS_123456" with template body {"code":"${code}"}.
var SMSSignName = ""
var SMSTemplateCode = ""

// SMSCodeLength is the number of digits in the generated verification
// code. 6 is industry-standard; shorter codes are easier to brute-force.
var SMSCodeLength = 6

// SMSCodeTTLSeconds controls how long a generated code stays valid.
// Defaults to 5 minutes — aligned with common.VerificationValidMinutes
// which is used by the underlying VerifyCodeWithKey helper.
var SMSCodeTTLSeconds = 300

// Rate-limit knobs — these are intentionally conservative to stop
// phone-number enumeration and SMS-fraud ("短信轰炸") out of the box.
var SMSPerPhoneInterval = 60   // seconds between two sends to the same phone
var SMSPerPhoneDailyMax = 10   // max sends per phone per 24h window
var SMSPerIPHourlyMax = 5      // max sends per client IP per hour
var SMSPerIPDailyMax = 20      // max sends per client IP per 24h

var TelegramBotToken = ""
var TelegramBotName = ""

var QuotaForNewUser = 0
var QuotaForInviter = 0
var QuotaForInvitee = 0
var ChannelDisableThreshold = 5.0
var AutomaticDisableChannelEnabled = false
var AutomaticEnableChannelEnabled = false
var QuotaRemindThreshold = 1000
var PreConsumedQuota = 500

var RetryTimes = 0

//var RootUserEmail = ""

var IsMasterNode bool

const (
	NodeNameSourceManual   = "manual"
	NodeNameSourceHostname = "hostname"
)

// NodeName 节点名称，优先从 NODE_NAME 环境变量读取，未配置时回退主机名。
// 用于审计日志和后台任务中标识节点身份；多实例部署时建议显式配置稳定 NODE_NAME。
var NodeName = ""

// NodeNameSource records how NodeName was chosen so future instance-management
// reporting can distinguish operator-configured names from automatic fallback.
var NodeNameSource = NodeNameSourceHostname

var NodeNameManuallyConfigured bool

var requestInterval int
var RequestInterval time.Duration

var SyncFrequency int // unit is second

var BatchUpdateEnabled = false
var BatchUpdateInterval int

var RelayTimeout int // unit is second

var RelayIdleConnTimeout int // unit is second
var RelayMaxIdleConns int
var RelayMaxIdleConnsPerHost int

var GeminiSafetySetting string

// https://docs.cohere.com/docs/safety-modes Type; NONE/CONTEXTUAL/STRICT
var CohereSafetySetting string

const (
	RequestIdKey         = "X-Oneapi-Request-Id"
	UpstreamRequestIdKey = "X-Upstream-Request-Id"
)

const (
	RoleGuestUser  = 0
	RoleCommonUser = 1
	RoleAdminUser  = 10
	RoleRootUser   = 100
)

func IsValidateRole(role int) bool {
	return role == RoleGuestUser || role == RoleCommonUser || role == RoleAdminUser || role == RoleRootUser
}

var (
	FileUploadPermission    = RoleGuestUser
	FileDownloadPermission  = RoleGuestUser
	ImageUploadPermission   = RoleGuestUser
	ImageDownloadPermission = RoleGuestUser
)

// All duration's unit is seconds
// Shouldn't larger then RateLimitKeyExpirationDuration
var (
	GlobalApiRateLimitEnable   bool
	GlobalApiRateLimitNum      int
	GlobalApiRateLimitDuration int64

	GlobalWebRateLimitEnable   bool
	GlobalWebRateLimitNum      int
	GlobalWebRateLimitDuration int64

	CriticalRateLimitEnable   bool
	CriticalRateLimitNum            = 20
	CriticalRateLimitDuration int64 = 20 * 60

	UploadRateLimitNum            = 10
	UploadRateLimitDuration int64 = 60

	DownloadRateLimitNum            = 10
	DownloadRateLimitDuration int64 = 60

	// Per-user search rate limit (applies after authentication, keyed by user ID)
	SearchRateLimitEnable         = true
	SearchRateLimitNum            = 10
	SearchRateLimitDuration int64 = 60
)

var RateLimitKeyExpirationDuration = 20 * time.Minute

const (
	UserStatusEnabled  = 1 // don't use 0, 0 is the default value!
	UserStatusDisabled = 2 // also don't use 0
)

const (
	TokenStatusEnabled   = 1 // don't use 0, 0 is the default value!
	TokenStatusDisabled  = 2 // also don't use 0
	TokenStatusExpired   = 3
	TokenStatusExhausted = 4
)

const (
	RedemptionCodeStatusEnabled  = 1 // don't use 0, 0 is the default value!
	RedemptionCodeStatusDisabled = 2 // also don't use 0
	RedemptionCodeStatusUsed     = 3 // also don't use 0
)

const (
	ChannelStatusUnknown          = 0
	ChannelStatusEnabled          = 1 // don't use 0, 0 is the default value!
	ChannelStatusManuallyDisabled = 2 // also don't use 0
	ChannelStatusAutoDisabled     = 3
)

const (
	TopUpStatusPending = "pending"
	TopUpStatusSuccess = "success"
	TopUpStatusFailed  = "failed"
	TopUpStatusExpired = "expired"
)

// ============================================================================
// Alipay Direct Payment Configuration
// ============================================================================

var AlipayEnabled = false
var AlipayAppId = ""
var AlipayPrivateKey = ""  // Application private key (RSA2)
var AlipayPublicKey = ""   // Alipay public key (for notification verification)

// ============================================================================
// WeChat Pay Direct Payment Configuration
// ============================================================================

var WxPayEnabled = false
var WxPayAppId = ""          // WeChat public account/app ID
var WxPayMchId = ""          // Merchant ID
var WxPayAPIv3Key = ""       // APIv3 key (32 chars)
var WxPayCertSerialNo = ""   // Certificate serial number
var WxPayPrivateKey = ""     // Merchant private key PEM content
var WxPayPublicKeyId = ""  // WeChat Pay public key ID (PUB_KEY_ID_...)
var WxPayPublicKey   = ""  // WeChat Pay public key PEM content
