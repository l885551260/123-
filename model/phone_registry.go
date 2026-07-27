// Modified on 2026-07-18 by project contributor under AGPLv3
// Phone registration audit log — records every SMS send and registration
// event for legal compliance and security auditing.

package model

import "time"

// Action types for phone_registry audit trail.
const (
	PhoneActionSMSSend       = "sms_send"
	PhoneActionSMSFail       = "sms_fail"
	PhoneActionSMSRateLimit  = "sms_rate_limit"
	PhoneActionRegister      = "register"
	PhoneActionRegisterFail  = "register_fail"
)

// PhoneRegistry records phone-related operations for auditing purposes.
// This table is append-only — rows are never updated or deleted in
// normal operation. Administrators may query it directly via SQL for
// compliance reviews or law-enforcement cooperation requests.
type PhoneRegistry struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Phone     string    `gorm:"size:32;index:idx_phone_registry_phone" json:"phone"`
	UserID    int       `gorm:"index:idx_phone_registry_user_id" json:"user_id"`
	Username  string    `gorm:"size:64" json:"username"`
	Action    string    `gorm:"size:32;index:idx_phone_registry_action" json:"action"`
	IP        string    `gorm:"size:45" json:"ip"`
	UserAgent string    `gorm:"size:512" json:"user_agent"`
	Provider  string    `gorm:"size:32" json:"provider"`
	Result    string    `gorm:"size:32" json:"result"`
	Error     string    `gorm:"size:256" json:"error"`
	CreatedAt time.Time `gorm:"index:idx_phone_registry_created_at" json:"created_at"`
}

func (PhoneRegistry) TableName() string {
	return "phone_registry"
}

// RecordPhoneAction is a convenience helper that inserts an audit row.
// It never returns an error to the caller — logging failures are silently
// swallowed so they cannot break the main request flow.
func RecordPhoneAction(phone string, userID int, username, action, ip, ua, provider, result, errMsg string) {
	if DB == nil {
		return
	}
	entry := PhoneRegistry{
		Phone:     phone,
		UserID:    userID,
		Username:  username,
		Action:    action,
		IP:        ip,
		UserAgent: ua,
		Provider:  provider,
		Result:    result,
		Error:     errMsg,
	}
	if err := DB.Create(&entry).Error; err != nil {
		// Best-effort: don't let audit failures break the main flow.
		_ = err
	}
}
