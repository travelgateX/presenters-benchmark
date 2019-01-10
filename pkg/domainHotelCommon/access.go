package domainHotelCommon

import "rfc/presenters/pkg/access"

type Access struct {
	AccessId      string
	Configuration *access.AccessConfiguration
	Settings      *SettingsBase
}
