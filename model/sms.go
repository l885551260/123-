// Modified by aytdai on 2026-07-18 under AGPLv3
// Phone/SMS registration feature — phone availability check (mirror of
// EnsureEmailAvailable / IsEmailAvailable). These helpers are invoked by
// the register controller before inserting a new user so the same
// mainland-China mobile number cannot be bound to two accounts.

package model

import (
	"strings"

	"github.com/QuantumNous/new-api/common"
)

// NormalizePhone trims whitespace and returns the canonical phone string.
// Mainland China numbers don't need case/whitespace normalization beyond
// TrimSpace, but having a dedicated helper keeps the API symmetric with
// NormalizeEmail and leaves room for future country-code support.
func NormalizePhone(phone string) string {
	return strings.TrimSpace(phone)
}

// IsPhoneAvailable returns true when no non-deleted user currently owns
// the given phone. excludeUserID lets callers skip a specific row (e.g.
// when a user edits their own profile). An empty phone is treated as
// "no phone bound" and is always considered available.
func IsPhoneAvailable(phone string, excludeUserID int) (bool, error) {
	phone = NormalizePhone(phone)
	if phone == "" {
		return true, nil
	}
	query := DB.Model(&User{}).Where("phone = ?", phone)
	if excludeUserID > 0 {
		query = query.Where("id <> ?", excludeUserID)
	}
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}

// EnsurePhoneAvailable returns nil when the phone is free to use, or
// ErrPhoneAlreadyTaken when another account already claims it. Database
// errors are propagated verbatim.
func EnsurePhoneAvailable(phone string, excludeUserID int) error {
	available, err := IsPhoneAvailable(phone, excludeUserID)
	if err != nil {
		return err
	}
	if !available {
		return ErrPhoneAlreadyTaken
	}
	return nil
}

// IsPhoneAlreadyTaken is a convenience wrapper matching the style of
// IsEmailAlreadyTaken.
func IsPhoneAlreadyTaken(phone string) bool {
	ok, err := IsPhoneAvailable(phone, 0)
	if err != nil {
		// On DB errors we conservatively report "taken" so registration
		// fails closed rather than accidentally allowing duplicates.
		common.SysLog("IsPhoneAlreadyTaken db error: " + err.Error())
		return true
	}
	return !ok
}

// InitPhoneIndex creates the partial unique index on users.phone
// (excluding empty strings) using raw SQL. This works around SQLite's
// limitation where ALTER TABLE cannot add a UNIQUE column — instead of
// letting GORM try (and fail) to create the constraint, we create the
// index ourselves after AutoMigrate has added the plain text column.
//
// Idempotent: "IF NOT EXISTS" makes it safe to call on every startup.
// Call this function once right after the AutoMigrate block in
// model/main.go.
func InitPhoneIndex() error {
	if DB == nil {
		return nil
	}
	sql := `CREATE UNIQUE INDEX IF NOT EXISTS idx_users_phone_unique
		ON users(phone)
		WHERE phone != '' AND phone IS NOT NULL`
	if err := DB.Exec(sql).Error; err != nil {
		common.SysLog("InitPhoneIndex failed: " + err.Error())
		return err
	}
	common.SysLog("InitPhoneIndex: partial unique index ensured on users.phone")
	return nil
}
