package graphResolver

import (
	"hub-aggregator/common/graphql"
	"rfc/presenters/pkg/domainHotelCommon"
)

type CancelPenaltyResolver struct {
	CancelPenalty *domainHotelCommon.CancelPenalty
}

func (r *CancelPenaltyResolver) HoursBefore() int32 {
	return int32(r.CancelPenalty.HoursBefore)
}

func (r *CancelPenaltyResolver) PenaltyType() domainHotelCommon.CancelPenaltyType {
	return (*r.CancelPenalty).Type
}

func (r *CancelPenaltyResolver) Currency() graphql.Currency {
	return graphql.Currency((*r.CancelPenalty).Currency)
}

func (r *CancelPenaltyResolver) Value() float64 {
	return r.CancelPenalty.Value
}
