package graphResolver

import "github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"

type PaxResolver struct {
	Pax *domainHotelCommon.Pax
}

func (r *PaxResolver) Age() int32 {
	return int32(r.Pax.Age)
}
