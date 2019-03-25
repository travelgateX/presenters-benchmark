package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type SurchargeResolver struct {
	Surcharge *domainHotelCommon.Surcharge
}

func (r *SurchargeResolver) ChargeType() domainHotelCommon.ChargeType {
	return (*r.Surcharge).ChargeType
}

func (r *SurchargeResolver) Mandatory() bool {
	return (*r.Surcharge).Mandatory
}

func (r *SurchargeResolver) Price() *PriceResolver {
	priceResolver := PriceResolver{Price: &((*r.Surcharge).Price)}
	return &priceResolver
}

func (r *SurchargeResolver) Description() *string {
	return (*r.Surcharge).Description
}
