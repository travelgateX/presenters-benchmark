package domainHotelCommon

import (
	"rfc/presenters/pkg/access"
	"hub-aggregator/common/resource"
)

type AccessData struct {
	Client              string
	SettingsBase        *SettingsBase
	AccessConfiguration *access.AccessConfiguration
	Supplier            *resource.Supplier
	HCHERedirect        bool
}
