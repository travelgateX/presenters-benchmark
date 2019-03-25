package domainHotelCommon

import "github.com/travelgateX/presenters-benchmark/pkg/access"

type Access struct {
	AccessId      string
	Configuration *access.AccessConfiguration
	Settings      *SettingsBase
}
