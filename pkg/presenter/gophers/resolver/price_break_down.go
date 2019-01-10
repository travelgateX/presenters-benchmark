package graphResolver

import (
	"hub-aggregator/common/graphql"
	"rfc/presenters/pkg/domainHotelCommon"
)

type PriceBreakDownResolver struct {
	PriceBreakDown *domainHotelCommon.PriceBreakDown
}

func (r *PriceBreakDownResolver) EffectiveDate() graphql.Date {
	return graphql.Date((*r.PriceBreakDown).EffectiveDate)
}

func (r *PriceBreakDownResolver) ExpireDate() graphql.Date {
	return graphql.Date((*r.PriceBreakDown).ExpireDate)
}

func (r *PriceBreakDownResolver) Price() *PriceResolver {
	priceResolver := PriceResolver{Price: &((*r.PriceBreakDown).Price)}
	return &priceResolver
}
