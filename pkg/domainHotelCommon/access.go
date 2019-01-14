package domainHotelCommon

import "presenters-benchmark/pkg/access"

type Access struct {
	AccessId      string
	Configuration *access.AccessConfiguration
	Settings      *SettingsBase
}
