package service

import (
	"strings"

	"github.com/Project Contributors/new-api/common"
	"github.com/Project Contributors/new-api/setting/system_setting"
)

func PaymentReturnURL(suffix string) string {
	base := strings.TrimRight(system_setting.ServerAddress, "/")
	return base + common.ThemeAwarePath(suffix)
}
