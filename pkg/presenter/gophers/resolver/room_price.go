package graphResolver

import "github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"

type RoomPriceResolver struct {
	RoomPrice *domainHotelCommon.RoomPrice
}

func (r *RoomPriceResolver) Price() *PriceResolver {
	priceResolver := PriceResolver{Price: &(*r.RoomPrice).Price}
	return &priceResolver
}

func (r *RoomPriceResolver) Breakdown() *[]*PriceBreakDownResolver {
	if r.RoomPrice == nil || len(r.RoomPrice.Breakdown) == 0 {
		return nil
	}
	pbds := make([]*PriceBreakDownResolver, 0, len(r.RoomPrice.Breakdown))
	for _, pbd := range r.RoomPrice.Breakdown {
		pbd_aux := pbd
		pbds = append(pbds, &PriceBreakDownResolver{PriceBreakDown: &pbd_aux})
	}
	return &pbds
}
