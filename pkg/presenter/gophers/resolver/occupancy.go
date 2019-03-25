package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type OccupancyResolver struct {
	Occupancy *domainHotelCommon.Occupancy
}

func (r *OccupancyResolver) Id() int32 {
	return int32(r.Occupancy.Id)
}

func (r *OccupancyResolver) Paxes() (paxes []*PaxResolver) {
	paxes = make([]*PaxResolver, 0, len(r.Occupancy.Paxes))

	if r.Occupancy != nil && r.Occupancy.Paxes != nil && len(r.Occupancy.Paxes) > 0 {
		for _, pax := range r.Occupancy.Paxes {
			pax_aux := pax
			paxes = append(paxes, &PaxResolver{Pax: &pax_aux})
		}
	}
	return
}
