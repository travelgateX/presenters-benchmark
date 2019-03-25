package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type PriceBreakDownResolver struct {
	PriceBreakDown *domainHotelCommon.PriceBreakDown
}

func (r *PriceBreakDownResolver) EffectiveDate() Date {
	return Date((*r.PriceBreakDown).EffectiveDate)
}

func (r *PriceBreakDownResolver) ExpireDate() Date {
	return Date((*r.PriceBreakDown).ExpireDate)
}

func (r *PriceBreakDownResolver) Price() *PriceResolver {
	priceResolver := PriceResolver{Price: &((*r.PriceBreakDown).Price)}
	return &priceResolver
}
