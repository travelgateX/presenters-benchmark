package graphResolver

import (
	"hub-aggregator/common/graphql"
	"rfc/presenters/pkg/domainHotelCommon"
)

type RatePlanResolver struct {
	RatePlan *domainHotelCommon.RatePlan
}

func (r *RatePlanResolver) Code() string {
	return *(*r.RatePlan).Code
}

func (r *RatePlanResolver) Name() *string {
	return (*r.RatePlan).Name
}

func (r *RatePlanResolver) EffectiveDate() *graphql.Date {
	if r.RatePlan.EffectiveDate == nil {
		return nil
	}
	tmp := graphql.Date(*r.RatePlan.EffectiveDate)
	return &tmp
}

func (r *RatePlanResolver) ExpireDate() *graphql.Date {
	if r.RatePlan.ExpireDate == nil {
		return nil
	}
	tmp := graphql.Date(*r.RatePlan.ExpireDate)
	return &tmp
}
