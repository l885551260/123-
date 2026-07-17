// Modified by aytdai on 2026-07-18 under AGPLv3
// Phone/SMS registration feature — in-memory rate limiter.

package common

import (
	"errors"
	"sync"
	"time"
)

// smsBucket tracks the last send time plus rolling counters for a single
// key (either a phone number or an IP). Counters are reset lazily on the
// next access once the relevant window has elapsed — no background
// goroutine is needed, which keeps the implementation simple and avoids
// goroutine leaks on long-running servers.
type smsBucket struct {
	lastSendAt     time.Time
	dailyCount     int
	dailyWindowAt  time.Time
	hourlyCount    int
	hourlyWindowAt time.Time
}

var (
	smsLimitMu    sync.Mutex
	smsPhoneLimit = map[string]*smsBucket{} // keyed by phone
	smsIPLimit    = map[string]*smsBucket{} // keyed by IP
)

// smsResetDaily zeroes the daily counter if its 24h window has elapsed.
func (b *smsBucket) resetDaily(now time.Time) {
	if now.Sub(b.dailyWindowAt) >= 24*time.Hour {
		b.dailyCount = 0
		b.dailyWindowAt = now
	}
}

// smsResetHourly zeroes the hourly counter if its 1h window has elapsed.
func (b *smsBucket) resetHourly(now time.Time) {
	if now.Sub(b.hourlyWindowAt) >= time.Hour {
		b.hourlyCount = 0
		b.hourlyWindowAt = now
	}
}

// SMSCheckRateLimit inspects the per-phone and per-IP buckets without
// mutating them. Returns ErrSMSRateLimited if any limit would be
// exceeded by the next send. The caller is expected to call
// SMSRecordSend after a successful dispatch.
func SMSCheckRateLimit(phone, ip string) error {
	smsLimitMu.Lock()
	defer smsLimitMu.Unlock()

	now := time.Now()

	if b, ok := smsPhoneLimit[phone]; ok {
		b.resetDaily(now)
		// Per-phone minimum interval between two sends.
		if now.Sub(b.lastSendAt) < time.Duration(SMSPerPhoneInterval)*time.Second {
			return ErrSMSRateLimited
		}
		// Per-phone daily cap.
		if b.dailyCount >= SMSPerPhoneDailyMax {
			return ErrSMSRateLimited
		}
	}
	if b, ok := smsIPLimit[ip]; ok {
		b.resetHourly(now)
		b.resetDaily(now)
		if b.hourlyCount >= SMSPerIPHourlyMax {
			return ErrSMSRateLimited
		}
		if b.dailyCount >= SMSPerIPDailyMax {
			return ErrSMSRateLimited
		}
	}
	return nil
}

// SMSRecordSend increments the per-phone and per-IP counters after a
// successful SMS dispatch. It is a no-op when both keys are empty.
func SMSRecordSend(phone, ip string) {
	smsLimitMu.Lock()
	defer smsLimitMu.Unlock()

	now := time.Now()
	if phone != "" {
		b := smsPhoneLimit[phone]
		if b == nil {
			b = &smsBucket{dailyWindowAt: now, hourlyWindowAt: now}
			smsPhoneLimit[phone] = b
		}
		b.resetDaily(now)
		b.lastSendAt = now
		b.dailyCount++
	}
	if ip != "" {
		b := smsIPLimit[ip]
		if b == nil {
			b = &smsBucket{dailyWindowAt: now, hourlyWindowAt: now}
			smsIPLimit[ip] = b
		}
		b.resetHourly(now)
		b.resetDaily(now)
		b.hourlyCount++
		b.dailyCount++
	}

	// Defensive pruning: once the phone map grows beyond a threshold we
	// drop stale entries (older than 25h) to avoid unbounded growth on
	// long-running servers. The IP map is pruned the same way.
	pruneStaleBuckets(smsPhoneLimit, now)
	pruneStaleBuckets(smsIPLimit, now)
}

const smsBucketStaleThreshold = 25 * time.Hour

func pruneStaleBuckets(m map[string]*smsBucket, now time.Time) {
	if len(m) < 10000 {
		return
	}
	for k, b := range m {
		if now.Sub(b.lastSendAt) > smsBucketStaleThreshold &&
			now.Sub(b.dailyWindowAt) > smsBucketStaleThreshold {
			delete(m, k)
		}
	}
}

// SMSResetLimitsForTest clears all counters. Only useful for unit tests
// and the admin console; not called by production code paths.
func SMSResetLimitsForTest() {
	smsLimitMu.Lock()
	defer smsLimitMu.Unlock()
	smsPhoneLimit = map[string]*smsBucket{}
	smsIPLimit = map[string]*smsBucket{}
}

// unused-import guard (errors is used by ErrSMSRateLimited in sms.go,
// but this file imports it to keep the bucket helpers self-contained if
// they are ever split out).
var _ = errors.New
